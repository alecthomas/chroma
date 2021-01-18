package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Tea lexer. Lexer for Tea Templates <http://teatrove.org/>.
var Tea = internal.Register(MustNewLexer(
	&Config{
		Name:      "Tea",
		Aliases:   []string{"tea"},
		Filenames: []string{"*.tea"},
		MimeTypes: []string{"text/x-tea"},
	},
	Rules{
		"root": {},
	},
))
