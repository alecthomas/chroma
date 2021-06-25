package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Stata lexer. For Stata <http://www.stata.com/> do files.
//
// Syntax based on
// - http://fmwww.bc.edu/RePEc/bocode/s/synlightlist.ado
// - https://github.com/isagalaev/highlight.js/blob/master/src/languages/stata.js
// - https://github.com/jpitblado/vim-stata/blob/master/syntax/stata.vim
var Stata = internal.Register(MustNewLexer(
	&Config{
		Name:      "Stata",
		Aliases:   []string{"stata", "do"},
		Filenames: []string{"*.do", "*.ado"},
		MimeTypes: []string{"text/x-stata", "text/stata", "application/x-stata"},
	},
	Rules{
		"root": {},
	},
))
