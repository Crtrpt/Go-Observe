// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func usage() {
	xprintf(`usage: go tool dist [command]
Commands are:

banner         print installation banner
bootstrap      rebuild everything
clean          deletes all built files
env [-p]       print environment (-p: include $PATH)
install [dir]  install individual directory
list [-json]   list all supported platforms
test [-h]      run Go test(s)
version        print Go version

All commands take -v flags to emit extra information.
`)
	xexit(2)
}

// commands records the available commands.
var commands = map[string]func(){
	"banner":    cmdbanner,
	"bootstrap": cmdbootstrap,
	"clean":     cmdclean,
	"env":       cmdenv,
	"install":   cmdinstall,
	"list":      cmdlist,
	"test":      cmdtest,
	"version":   cmdversion,
}

// main takes care of OS-specific startup and dispatches to xmain.
func main() {
	os.Setenv("TERM", "dumb") // disable escape codes in clang errors


	gohostos = runtime.GOOS


	sysinit()


	gohostos = "darwin"
	gohostarch = "amd64"


	// For deterministic make.bash debugging and for smallest-possible footprint,
	// pay attention to GOMAXPROCS=1.  This was a bad idea for 1.4 bootstrap, but
	// the bootstrap version is now 1.17+ and thus this is fine.
	if runtime.GOMAXPROCS(0) == 1 {
		maxbg = 1
	}
	bginit()
	xinit()
	xmain()
	xexit(0)
}

// The OS-specific main calls into the portable code here.
func xmain() {
	if len(os.Args) < 2 {
		usage()
	}
	cmd := os.Args[1]
	os.Args = os.Args[1:] // for flag parsing during cmd
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: go tool dist %s [options]\n", cmd)
		flag.PrintDefaults()
		os.Exit(2)
	}
	if f, ok := commands[cmd]; ok {
		f()
	} else {
		xprintf("unknown command %s\n", cmd)
		usage()
	}
}
