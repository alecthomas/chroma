package styles

import (
	. "github.com/alecthomas/chroma/v2"
)

// MPL style.
var MPL = Register(MustNewStyle("mpl", StyleEntries{
	Background:       "#ffffff",
	Text:             "#000000",
	Error:            "#ff0000",
	Comment:          "#008000",
	CommentPreproc:   "#0000ff",
	Keyword:          "#0000ff",
	KeywordType:      "#0000ff",
	Operator:         "#000000",
	Punctuation:      "#000000",
	Name:             "#000000",
	NameBuiltin:      "#0000ff",
	NameClass:        "#0000ff",
	NameFunction:     "#0000ff",
	NameNamespace:    "#0000ff",
	NameVariable:     "#000000",
	LiteralString:    "#a31515",
	LiteralNumber:    "#098658",
	LiteralNumberHex: "#098658",
	TextWhitespace:   "#000000",
}))
