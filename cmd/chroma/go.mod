module github.com/alecthomas/chroma/v2/cmd/chroma

go 1.17

replace github.com/alecthomas/chroma/v2 => ../../

require (
	github.com/alecthomas/chroma/v2 v2.0.0-00010101000000-000000000000
	github.com/alecthomas/kong v0.2.17
	github.com/mattn/go-colorable v0.1.12
	github.com/mattn/go-isatty v0.0.14
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
)
