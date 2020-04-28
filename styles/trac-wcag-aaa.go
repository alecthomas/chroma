package styles

import (
	"github.com/alecthomas/chroma"
)

// Trac-hc style.
var Trac-hc = Register(chroma.MustNewStyle("trac-hc", chroma.StyleEntries{
	chroma.TextWhitespace:     "#555555",
	chroma.Comment:            "italic #595946",
	chroma.CommentPreproc:     "bold noitalic #585858",
	chroma.CommentSpecial:     "bold #585858",
	chroma.Operator:           "bold",
	chroma.LiteralString:      "#724B00",
	chroma.LiteralStringRegex: "#545400",
	chroma.LiteralNumber:      "#005A5A",
	chroma.Keyword:            "bold",
	chroma.KeywordType:        "#445588",
	chroma.NameBuiltin:        "#585858",
	chroma.NameFunction:       "bold #990000",
	chroma.NameClass:          "bold #445588",
	chroma.NameException:      "bold #990000",
	chroma.NameNamespace:      "#555555",
	chroma.NameVariable:       "#085C5C",
	chroma.NameConstant:       "#085C5C",
	chroma.NameTag:            "#000080",
	chroma.NameAttribute:      "#085C5C",
	chroma.NameEntity:         "#800080",
	chroma.GenericHeading:     "#585858",
	chroma.GenericSubheading:  "#575757",
	chroma.GenericDeleted:     "bg:#ffdddd #000000",
	chroma.GenericInserted:    "bg:#ddffdd #000000",
	chroma.GenericError:       "#aa0000",
	chroma.GenericEmph:        "italic",
	chroma.GenericStrong:      "bold",
	chroma.GenericPrompt:      "#555555",
	chroma.GenericOutput:      "#585858",
	chroma.GenericTrac-hceback:   "#aa0000",
	chroma.GenericUnderline:   "underline",
	chroma.Error:              "bg:#e3d2d2 #820101",
	chroma.Background:         " bg:#ffffff",
}))
