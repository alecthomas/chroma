package lexers

import (
	"github.com/alecthomas/chroma/v2"
)

func init() {
	Register(XPhage)
}

// XPhage lexer definition.
var XPhage = chroma.MustNewLexer(
	&chroma.Config{
		Name:      "X-Phage",
		Aliases:   []string{"xphage"},
		Filenames: []string{"*.xp0", "*.xh", "*.xui"},
		MimeTypes: []string{"text/xphage"},
	},
	func() chroma.Rules {
		return chroma.Rules{
			"root": {
				// Whitespace – emit as Text
				{`\s+`, chroma.Text, nil},

				// Single-line comments
				{`//.*?$`, chroma.CommentSingle, nil},

				// Multi-line comments
				{`/\*`, chroma.CommentMultiline, chroma.Push("comment")},

				// Double‑quoted strings
				{`"(?:\\.|[^"\\])*"`, chroma.String, nil},

				// Keywords
				{`\b(pulse|fusion|matrix|global|quantum|bypass|synapse|chronos|ether|void|scan|vortex|~link|math)\b`, chroma.Keyword, nil},

				// Type keywords
				{`\b(atom|shadow)\b`, chroma.KeywordType, nil},

				// Built‑in functions
				{`\b(beam)\b`, chroma.NameFunction, nil},

				// UI component classes
				{`\b(Signal|Vision|Orbit|Trigger|Input|Z_Plane)\b`, chroma.NameClass, nil},

				// Attributes (e.g., @UIElement)
				{`@[a-zA-Z_]+`, chroma.NameAttribute, nil},

				// Numbers (decimal and hex)
				{`\b[0-9]+(\.[0-9]+)?\b`, chroma.Number, nil},
				{`\b0x[0-9a-fA-F]+\b`, chroma.Number, nil},

				// Operators
				{`->`, chroma.Operator, nil},
				{`[+\-*/=~]`, chroma.Operator, nil},

				// Punctuation
				{`[{}()\[\];,]`, chroma.Punctuation, nil},

				// Identifiers (everything else)
				{`\b[A-Za-z_][A-Za-z0-9_]*\b`, chroma.Name, nil},
			},
			"comment": {
				{`[^*/]+`, chroma.CommentMultiline, nil},
				{`\*/`, chroma.CommentMultiline, chroma.Pop(1)},
				{`[*/]`, chroma.CommentMultiline, nil},
			},
		}
	},
)