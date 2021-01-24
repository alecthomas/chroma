package q

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// QVTO lexer. For the QVT Operational Mapping language <http://www.omg.org/spec/QVT/1.1/>.
var QVTO = internal.Register(MustNewLexer(
	&Config{
		Name:      "QVTO",
		Aliases:   []string{"qvto", "qvt"},
		Filenames: []string{"*.qvto"},
	},
	Rules{
		"root": {},
	},
))
