module github.com/alecthomas/chroma/cmd/chromad

go 1.13

require (
	github.com/GeertJohan/go.rice v1.0.1-0.20191102153406-d954009f7238
	github.com/alecthomas/chroma v0.7.0
	github.com/alecthomas/kong v0.2.1
	github.com/alecthomas/kong-hcl v0.2.0
	github.com/gorilla/csrf v1.6.2
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
)

replace github.com/alecthomas/chroma => ../../
