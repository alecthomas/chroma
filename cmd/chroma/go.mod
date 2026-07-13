module github.com/alecthomas/chroma/v3/cmd/chroma

go 1.25

replace github.com/alecthomas/chroma/v3 => ../../

require (
	github.com/alecthomas/chroma/v3 v3.0.0-alpha.5
	github.com/alecthomas/kong v1.15.0
	github.com/mattn/go-colorable v0.1.15
	github.com/mattn/go-isatty v0.0.22
)

require (
	github.com/dlclark/regexp2/v2 v2.2.2 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
