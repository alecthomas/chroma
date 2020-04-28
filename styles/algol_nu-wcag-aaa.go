package styles

import (
	"github.com/alecthomas/chroma"
)

// AlgolNuHighContrast style.
var AlgolNuHighContrast = Register(chroma.MustNewStyle("algol_nu-hc", chroma.StyleEntries{
	chroma.Comment:            "italic #585858",
	chroma.CommentPreproc:     "bold noitalic #585858",
	chroma.CommentSpecial:     "bold noitalic #585858",
	chroma.Keyword:            "bold",
	chroma.KeywordDeclaration: "italic",
	chroma.NameBuiltin:        "bold italic",
	chroma.NameBuiltinPseudo:  "bold italic",
	chroma.NameNamespace:      "bold italic #535353",
	chroma.NameClass:          "bold italic #535353",
	chroma.NameFunction:       "bold italic #535353",
	chroma.NameVariable:       "bold italic #535353",
	chroma.NameConstant:       "bold italic #535353",
	chroma.OperatorWord:       "bold",
	chroma.LiteralString:      "italic #535353",
	chroma.Error:              "border:#FF0000",
	chroma.Background:         " bg:#ffffff",
}))
