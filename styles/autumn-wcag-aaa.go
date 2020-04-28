package styles

import (
	"github.com/alecthomas/chroma"
)

// Autumn-hc style.
var Autumn-hc = Register(chroma.MustNewStyle("autumn-hc", chroma.StyleEntries{
	chroma.TextWhitespace:      "#555555",
	chroma.Comment:             "italic #575757",
	chroma.CommentPreproc:      "noitalic #345F02",
	chroma.CommentSpecial:      "italic #0000aa",
	chroma.Keyword:             "#0000aa",
	chroma.KeywordType:         "#056262",
	chroma.OperatorWord:        "#0000aa",
	chroma.NameBuiltin:         "#056262",
	chroma.NameFunction:        "#026302",
	chroma.NameClass:           "underline #026302",
	chroma.NameNamespace:       "underline #056262",
	chroma.NameVariable:        "#aa0000",
	chroma.NameConstant:        "#aa0000",
	chroma.NameEntity:          "bold #880000",
	chroma.NameAttribute:       "#00539A",
	chroma.NameTag:             "bold #00539A",
	chroma.NameDecorator:       "#585858",
	chroma.LiteralString:       "#8A4402",
	chroma.LiteralStringSymbol: "#0000aa",
	chroma.LiteralStringRegex:  "#005A5A",
	chroma.LiteralNumber:       "#005A5A",
	chroma.GenericHeading:      "bold #000080",
	chroma.GenericSubheading:   "bold #800080",
	chroma.GenericDeleted:      "#aa0000",
	chroma.GenericInserted:     "#026302",
	chroma.GenericError:        "#aa0000",
	chroma.GenericEmph:         "italic",
	chroma.GenericStrong:       "bold",
	chroma.GenericPrompt:       "#555555",
	chroma.GenericOutput:       "#585858",
	chroma.GenericTraceback:    "#aa0000",
	chroma.GenericUnderline:    "underline",
	chroma.Error:               "#6B0000 bg:#FAA",
	chroma.Background:          " bg:#ffffff",
}))
