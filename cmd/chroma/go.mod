module github.com/alecthomas/chroma/v2/cmd/chroma

go 1.25

replace github.com/alecthomas/chroma/v2 => ../../

require (
	github.com/alecthomas/chroma/v2 v2.26.1
	github.com/alecthomas/kong v1.15.0
	github.com/mattn/go-colorable v0.1.15
	github.com/mattn/go-isatty v0.0.22
)

require (
	github.com/dlclark/regexp2/v2 v2.2.1 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
