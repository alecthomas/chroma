package styles

import (
	"github.com/alecthomas/chroma"
)

// Borland-hc style.
var Borland-hc = Register(chroma.MustNewStyle("borland-hc", chroma.StyleEntries{
	chroma.TextWhitespace:    "#555555",
	chroma.Comment:           "italic #026302",
	chroma.CommentPreproc:    "noitalic #085C5C",
	chroma.CommentSpecial:    "noitalic bold",
	chroma.LiteralString:     "#0000FF",
	chroma.LiteralStringChar: "#800080",
	chroma.LiteralNumber:     "#0000FF",
	chroma.Keyword:           "bold #000080",
	chroma.OperatorWord:      "bold",
	chroma.NameTag:           "bold #000080",
	chroma.NameAttribute:     "#A80303",
	chroma.GenericHeading:    "#585858",
	chroma.GenericSubheading: "#575757",
	chroma.GenericDeleted:    "bg:#ffdddd #000000",
	chroma.GenericInserted:   "bg:#ddffdd #000000",
	chroma.GenericError:      "#aa0000",
	chroma.GenericEmph:       "italic",
	chroma.GenericStrong:     "bold",
	chroma.GenericPrompt:     "#555555",
	chroma.GenericOutput:     "#585858",
	chroma.GenericTraceback:  "#aa0000",
	chroma.GenericUnderline:  "underline",
	chroma.Error:             "bg:#e3d2d2 #820101",
	chroma.Background:        " bg:#ffffff",
}))
