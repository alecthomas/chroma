module github.com/alecthomas/chroma/v2/cmd/chroma

go 1.19

replace github.com/alecthomas/chroma/v2 => ../../

require (
	github.com/alecthomas/chroma/v2 v2.0.0-00010101000000-000000000000
	github.com/alecthomas/kong v0.8.0
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.19
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)
