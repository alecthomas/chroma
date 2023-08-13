module github.com/alecthomas/chroma/v2/cmd/chromad

go 1.19

require (
	github.com/alecthomas/chroma/v2 v2.0.0-00010101000000-000000000000
	github.com/alecthomas/kong v0.8.0
	github.com/alecthomas/kong-hcl v1.0.1
	github.com/gorilla/csrf v1.7.1
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

replace github.com/alecthomas/chroma/v2 => ../../
