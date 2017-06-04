package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Fallback lexer if no other is found.
var Fallback = Register(NewLexer(&Config{
	Name:      "fallback",
	Filenames: []string{"*"},
	Priority:  99,
}, Rules{
	"root": []Rule{
		{`.+`, Text, nil},
		{`\n`, Text, nil},
	},
}))
