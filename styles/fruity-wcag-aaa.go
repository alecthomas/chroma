package styles

import (
	"github.com/alecthomas/chroma"
)

// Fruity-hc style.
var Fruity-hc = Register(chroma.MustNewStyle("fruity-hc", chroma.StyleEntries{
	chroma.TextWhitespace:    "#A6A6A6",
	chroma.Background:        "#ffffff bg:#111111",
	chroma.GenericOutput:     "#B1B1B1 bg:#222222",
	chroma.Keyword:           "#FF7642 bold",
	chroma.KeywordPseudo:     "nobold",
	chroma.LiteralNumber:     "#70A5FF bold",
	chroma.NameTag:           "#FF7642 bold",
	chroma.NameVariable:      "#FF7642",
	chroma.Comment:           "#51B651 bg:#0f140f italic",
	chroma.NameAttribute:     "#FF75A5 bold",
	chroma.LiteralString:     "#51A7F3",
	chroma.NameFunction:      "#FF75A5 bold",
	chroma.GenericHeading:    "#ffffff bold",
	chroma.KeywordType:       "#cdcaa9 bold",
	chroma.GenericSubheading: "#ffffff bold",
	chroma.NameConstant:      "#51A7F3",
	chroma.CommentPreproc:    "#FF7575 bold",
}))
