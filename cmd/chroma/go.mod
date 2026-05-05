module github.com/alecthomas/chroma/v2/cmd/chroma

go 1.23

toolchain go1.26.2

replace github.com/alecthomas/chroma/v2 => ../../

require (
	github.com/alecthomas/chroma/v2 v2.23.1
	github.com/alecthomas/kong v1.15.0
	github.com/mattn/go-colorable v0.1.14
	github.com/mattn/go-isatty v0.0.22
)

require (
	github.com/dlclark/regexp2 v1.12.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
