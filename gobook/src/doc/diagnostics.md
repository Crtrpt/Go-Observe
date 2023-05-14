# diagnostics 问题诊断


参考文档
- [官方文档:https://go.dev/doc/diagnostics#introduction](https://go.dev/doc/diagnostics#introduction)

诊断分类主要一下几个
- (剖析)Profiling: Profiling tools analyze the complexity and costs of a Go program such as its memory usage and frequently called functions to identify the expensive sections of a Go program.
- （跟踪）Tracing: Tracing is a way to instrument code to analyze latency throughout the lifecycle of a call or user request. Traces provide an overview of how much latency each component contributes to the overall latency in a system. Traces can span multiple Go processes.
-  (调试)Debugging: Debugging allows us to pause a Go program and examine its execution. Program state and flow can be verified with debugging.
- (运行时分析和事件)Runtime statistics and events: Collection and analysis of runtime stats and events provides a high-level overview of the health of Go programs. Spikes/dips of metrics helps us to identify changes in throughput, utilization, and performance.


## Profiling 剖析
主要的`Profiling`包

- cpu: CPU profile determines where a program spends its time while actively consuming CPU cycles (as opposed to while sleeping or waiting for I/O).
- (堆分析)heap: Heap profile reports memory allocation samples; used to monitor current and historical memory usage, and to check for memory leaks.
(线程的创建)threadcreate: Thread creation profile reports the sections of the program that lead the creation of new OS threads.
- (协程)goroutine: Goroutine profile reports the stack traces of all current goroutines.
- (阻塞)block: Block profile shows where goroutines block waiting on synchronization primitives (including timer channels). Block profile is not enabled by default; use runtime.SetBlockProfileRate to enable it.
- (锁)mutex: Mutex profile reports the lock contentions. When you think your CPU is not fully utilized due to a mutex contention, use this profile. Mutex profile is not enabled by default, see runtime.SetMutexProfileFraction to enable it.

- 还不够?那就上操作系统级别的工具
  [参考文档](https://perf.wiki.kernel.org/index.php/Tutorial)
  [ebpf]

## Tracing 跟踪
go 目前暂不提供自动的函数调用跟踪办法 需要手动添加


## Debugging 调试
- Delve: go 实现的程序调试器
- GDB: 通用调试器 gccgo 提供对 go程序的支持

更好的调试建议
关闭编译器优化
```
 go build -gcflags=all="-N -l"
```
1.10  之后新的编译参数
```
$ go build -gcflags="-dwarflocationlists=true"
```

使用 core dump  文件 进行线上调试


## Runtime statistics and events 运行时统计和事件

- runtime.ReadMemStats 
- debug.ReadGCStats
- debug.Stack
- debug.WriteHeapDump 
- runtime.NumGoroutine

### Execution tracer
Go comes with a runtime execution tracer to capture a wide range of runtime events. Scheduling, syscall, garbage collections, heap size, and other events are collected by runtime and available for visualization by the go tool trace. Execution tracer is a tool to detect latency and utilization problems. You can examine how well the CPU is utilized, and when networking or syscalls are a cause of preemption for the goroutines.

### GODEBUG 设置debug参数

GODEBUG=gctrace=1 prints garbage collector events at each collection, summarizing the amount of memory collected and the length of the pause.
GODEBUG=inittrace=1 prints a summary of execution time and memory allocation information for completed package initialization work.
GODEBUG=schedtrace=X prints scheduling events every X milliseconds.

GODEBUG=cpu.all=off disables the use of all optional instruction set extensions.
GODEBUG=cpu.extension=off disables use of instructions from the specified instruction set extension.
extension is the lower case name for the instruction set extension such as sse41 or avx.    
