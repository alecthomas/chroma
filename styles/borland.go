package styles

import (
	. "github.com/alecthomas/chroma" // nolint: golint
)

// Borland style.
var Borland = Register(New("borland", Entries{
	Whitespace: "#bbbbbb",

	Comment:        "italic #008800",
	CommentPreproc: "noitalic #008080",
	CommentSpecial: "noitalic bold",

	String:        "#0000FF",
	StringChar:    "#800080",
	Number:        "#0000FF",
	Keyword:       "bold #000080",
	OperatorWord:  "bold",
	NameTag:       "bold #000080",
	NameAttribute: "#FF0000",

	GenericHeading:    "#999999",
	GenericSubheading: "#aaaaaa",
	GenericDeleted:    "bg:#ffdddd #000000",
	GenericInserted:   "bg:#ddffdd #000000",
	GenericError:      "#aa0000",
	GenericEmph:       "italic",
	GenericStrong:     "bold",
	GenericPrompt:     "#555555",
	GenericOutput:     "#888888",
	GenericTraceback:  "#aa0000",

	Error: "bg:#e3d2d2 #a61717",
}))
