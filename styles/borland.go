package styles

import (
	"github.com/alecthomas/chroma"
)

// Borland style.
var Borland = Register(chroma.NewStyle("borland", chroma.StyleEntries{
	chroma.Whitespace: "#bbbbbb",

	chroma.Comment:        "italic #008800",
	chroma.CommentPreproc: "noitalic #008080",
	chroma.CommentSpecial: "noitalic bold",

	chroma.String:        "#0000FF",
	chroma.StringChar:    "#800080",
	chroma.Number:        "#0000FF",
	chroma.Keyword:       "bold #000080",
	chroma.OperatorWord:  "bold",
	chroma.NameTag:       "bold #000080",
	chroma.NameAttribute: "#FF0000",

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

	chroma.Error: "bg:#e3d2d2 #a61717",
}))
