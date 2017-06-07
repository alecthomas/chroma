package lexers

import (
	"github.com/alecthomas/chroma"
)

// Fallback lexer if no other is found.
var Fallback chroma.Lexer = chroma.MustNewLexer(&chroma.Config{
	Name:      "fallback",
	Filenames: []string{"*"},
}, chroma.Rules{
	"root": []chroma.Rule{
		{`.+`, chroma.Text, nil},
		{`\n`, chroma.Text, nil},
	},
})
