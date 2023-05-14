# Data Race Detector(竞态检测)


## 什么是竞态
两个协程至少一个写入变量的时候发生竞争

Here is an example of a data race that can lead to crashes and memory corruption:
```golang
func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

如何使用
使用 `-race` 编译参数


```
go run -race  main.go 
```

问题

```
首先，使用竞争检测器 (go test -race) 运行测试。 竞争检测器只能发现在运行时发生的竞争，因此它无法在未执行的代码路径中发现竞争。 如果您的测试覆盖不完整，您可能会通过在实际工作负载下运行使用 -race 构建的二进制文件来发现更多竞争。
```


常见案例

## 循环计数器
```golang
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}
```
通常输出 `55555`

修复

```golang
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

## 共享变量

``` golang
// ParallelWrite writes data to file1 and file2, returns the errors.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
```
修复

```golang
            ...
			_, err := f1.Write(data)
			...
			_, err := f2.Write(data)
			...
```


## 没有同步访问的全局变量


```golang
var service map[string]net.Addr

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}
```

修复
```golang
var (
	service   map[string]net.Addr
	serviceMu sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) net.Addr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}
```


## 未保护的原始变量

```golang
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
```

无锁修复
```golang
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}

```


### 不同步的发送和关闭操作


```golang
c := make(chan struct{}) // or buffered channel

// The race detector cannot derive the happens before relation
// for the following send and close operations. These two operations
// are unsynchronized and happen concurrently.
go func() { c <- struct{}{} }()
close(c)
```
修复
```golang
c := make(chan struct{}) // or buffered channel

go func() { c <- struct{}{} }()
<-c
close(c)
```

开启竞态分析的开销

内存 5-10倍 执行事件 2-20 倍



竞争检测器目前为每个延迟和恢复语句额外分配 8 个字节。 在 goroutine 退出之前，这些额外的分配不会被回收。 这意味着，如果您有一个长时间运行的 goroutine，它会定期发出 defer 和 recover 调用，程序内存使用量可能会无限增长。 这些内存分配不会显示在 runtime.ReadMemStats 或 runtime/pprof 的输出中。