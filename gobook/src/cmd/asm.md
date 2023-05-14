# asm


```
go tool asm [flags] file
```
```
Flags:

	-D name[=value]
		Predefine symbol name with an optional simple value.
		Can be repeated to define multiple symbols.
	-I dir1 -I dir2
		Search for #include files in dir1, dir2, etc,
		after consulting $GOROOT/pkg/$GOOS_$GOARCH.
	-S
		Print assembly and machine code.
	-V
		Print assembler version and exit.
	-debug
		Dump instructions as they are parsed.
	-dynlink
		Support references to Go symbols defined in other shared libraries.
	-gensymabis
		Write symbol ABI information to output file. Don't assemble.
	-o file
		Write output to file. The default is foo.o for /a/b/c/foo.s.
	-p pkgpath
		Set expected package import to pkgpath.
	-shared
		Generate code that can be linked into a shared library.
	-spectre list
		Enable spectre mitigations in list (all, ret).
	-trimpath prefix
		Remove prefix from recorded source file paths.
```

```
Flags:

	-D name[=value]
		Predefine symbol name with an optional simple value.
		Can be repeated to define multiple symbols.
	-I dir1 -I dir2
		Search for #include files in dir1, dir2, etc,
		after consulting $GOROOT/pkg/$GOOS_$GOARCH.
	-S
		打印汇编和机器码
	-V
		打印版本·
	-debug
		Dump instructions as they are parsed.
	-dynlink
		Support references to Go symbols defined in other shared libraries.
	-gensymabis
		Write symbol ABI information to output file. Don't assemble.
	-o file
		Write output to file. The default is foo.o for /a/b/c/foo.s.
	-p pkgpath
		Set expected package import to pkgpath.
	-shared
		Generate code that can be linked into a shared library.
	-spectre list
		Enable spectre mitigations in list (all, ret).
	-trimpath prefix
		Remove prefix from recorded source file paths.
```