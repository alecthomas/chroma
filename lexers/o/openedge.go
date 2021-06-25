package o

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// OpenEdge ABL lexer.
var OpenEdgeABL = internal.Register(MustNewLexer(
	&Config{
		Name:      "OpenEdge ABL",
		Aliases:   []string{"openedge", "abl", "progress"},
		Filenames: []string{"*.p", "*.cls"},
		MimeTypes: []string{"text/x-openedge", "application/x-openedge"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	// try to identify OpenEdge ABL based on a few common constructs.
	var result float32

	if strings.Contains(text, "END.") {
		result += 0.05
	}

	if strings.Contains(text, "END PROCEDURE.") {
		result += 0.05
	}

	if strings.Contains(text, "ELSE DO:") {
		result += 0.05
	}

	return result
}))
