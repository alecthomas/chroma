package lexers

import (
	"github.com/alecthomas/chroma/v2"
)

// Flint lexer.
var Flint = Register(chroma.MustNewLexer(
	&chroma.Config{
		Name:      "Flint",
		Aliases:   []string{"flint"},
		Filenames: []string{"*.fl"},
		MimeTypes: []string{"text/x-flint"},
	},
	func() chroma.Rules {
		return chroma.Rules{
			"root": {
				// Comments
				{Pattern: `--.*`, Type: chroma.CommentSingle},
				{Pattern: `\{-[\s\S]*?-\}`, Type: chroma.CommentMultiline},

				// Strings and Interpolation
				{Pattern: `"(\\\\|\\"|[^"])*"`, Type: chroma.StringDouble},
				{Pattern: "`[^`]*`", Type: chroma.StringBacktick},
				{Pattern: `\$"(\\\\|\\"|[^"])*"`, Type: chroma.StringInterpol},
				{Pattern: `\$` + "`[^`]*`", Type: chroma.StringInterpol},

				// Keywords
				{Pattern: `\b(import|as|if|else|for|while|stream|return|break|continue|fn|extern|var|const|struct)\b`, Type: chroma.Keyword},
				{Pattern: `\b(true|false|null)\b`, Type: chroma.KeywordConstant},

				// Types
				{Pattern: `\b(int|float|string|bool|void|val|arr|dict)\b`, Type: chroma.KeywordType},

				// Built-ins
				{Pattern: `\b(print|printerr|len|push|range|if_fail|fallback|ensure|to_int|to_str|to_float|clone|lines|chars|grep|embed_file|type_of)\b`, Type: chroma.NameBuiltin},

				// Numbers
				{Pattern: `\b\d+(\.\d+)?([eE][+-]?\d+)?\b`, Type: chroma.Number},

				// Operators
				{Pattern: `~>`, Type: chroma.Operator},
				{Pattern: `[+/*%=<>&|!\-\.]+`, Type: chroma.Operator},

				// Identifiers and Punctuation
				{Pattern: `[a-zA-Z_]\w*`, Type: chroma.Name},
				{Pattern: `[{}()\[\],;:]`, Type: chroma.Punctuation},

				// Whitespace
				{Pattern: `\s+`, Type: chroma.Text},
			},
		}
	},
))
