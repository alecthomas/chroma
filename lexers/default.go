package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

// Default lexer if no other is found.
var Default = Register(NewLexer(&Config{
	Name:      "default",
	Filenames: []string{"*"},
	Priority:  99,
}, Rules{
	"root": []Rule{
		{`.+`, Text, nil},
		{`\n`, Text, nil},
	},
}))
