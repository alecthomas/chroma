module github.com/alecthomas/chroma/v2/cmd/chroma

go 1.19

replace github.com/alecthomas/chroma/v2 => ../../

require (
	github.com/alecthomas/chroma/v2 v2.12.0
	github.com/alecthomas/kong v0.9.0
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.20
)

require (
	github.com/dlclark/regexp2 v1.11.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)
