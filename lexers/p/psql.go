package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// PostgresConsole lexer.
var PostgresConsole = internal.Register(MustNewLexer(
	&Config{
		Name:      "PostgreSQL console (psql)",
		Aliases:   []string{"psql", "postgresql-console", "postgres-console"},
		MimeTypes: []string{"text/x-postgresql-psql"},
	},
	Rules{
		"root": {},
	},
))
