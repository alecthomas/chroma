package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// SmartGameFormat lexer. Lexer for Smart Game Format (sgf) file format.
//
// The format is used to store game records of board games for two players
// (mainly Go game). For more information about the definition of the format,
// see: https://www.red-bean.com/sgf/
var SmartGameFormat = internal.Register(MustNewLexer(
	&Config{
		Name:      "SmartGameFormat",
		Aliases:   []string{"sgf"},
		Filenames: []string{"*.sgf"},
	},
	Rules{
		"root": {},
	},
))
