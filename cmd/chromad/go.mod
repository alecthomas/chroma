module github.com/alecthomas/chroma/v2/cmd/chromad

go 1.22

toolchain go1.24.2

require (
	github.com/alecthomas/chroma/v2 v2.16.0
	github.com/alecthomas/kong v1.10.0
	github.com/alecthomas/kong-hcl v1.0.1
	github.com/gorilla/csrf v1.7.3
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
)

require (
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

replace github.com/alecthomas/chroma/v2 => ../../
