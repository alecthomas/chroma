package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// ClojureScript lexer.
var ClojureScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "ClojureScript",
		Aliases:   []string{"clojurescript", "cljs"},
		Filenames: []string{"*.cljs"},
		MimeTypes: []string{"text/x-clojurescript", "application/x-clojurescript"},
	},
	Rules{
		"root": {},
	},
))
