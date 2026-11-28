module github.com/alecthomas/chroma/v3/cmd/chroma

go 1.25

// Always built from this checkout, so the required chroma version is a
// placeholder that never needs bumping. This module cannot be go installed.
replace github.com/alecthomas/chroma/v3 => ../../

require (
	github.com/alecthomas/chroma/v3 v3.0.0-alpha.5
	github.com/alecthomas/kong v1.16.1
	github.com/mattn/go-colorable v0.1.16
	github.com/mattn/go-isatty v0.0.24
)

require (
	github.com/dlclark/regexp2/v2 v2.10.1 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
