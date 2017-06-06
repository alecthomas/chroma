package styles

import (
	"github.com/alecthomas/chroma"
)

// Trac colour scheme.
var Trac = Register(chroma.NewStyle("trac", map[chroma.TokenType]string{
	chroma.Whitespace:        "#bbbbbb",
	chroma.Comment:           "italic #999988",
	chroma.CommentPreproc:    "bold noitalic #999999",
	chroma.CommentSpecial:    "bold #999999",
	chroma.Operator:          "bold",
	chroma.String:            "#bb8844",
	chroma.StringRegex:       "#808000",
	chroma.Number:            "#009999",
	chroma.Keyword:           "bold",
	chroma.KeywordType:       "#445588",
	chroma.NameBuiltin:       "#999999",
	chroma.NameFunction:      "bold #990000",
	chroma.NameClass:         "bold #445588",
	chroma.NameException:     "bold #990000",
	chroma.NameNamespace:     "#555555",
	chroma.NameVariable:      "#008080",
	chroma.NameConstant:      "#008080",
	chroma.NameTag:           "#000080",
	chroma.NameAttribute:     "#008080",
	chroma.NameEntity:        "#800080",
	chroma.GenericHeading:    "#999999",
	chroma.GenericSubheading: "#aaaaaa",
	chroma.GenericDeleted:    "bg:#ffdddd #000000",
	chroma.GenericInserted:   "bg:#ddffdd #000000",
	chroma.GenericError:      "#aa0000",
	chroma.GenericEmph:       "italic",
	chroma.GenericStrong:     "bold",
	chroma.GenericPrompt:     "#555555",
	chroma.GenericOutput:     "#888888",
	chroma.GenericTraceback:  "#aa0000",
	chroma.Error:             "bg:#e3d2d2 #a61717",
}))
