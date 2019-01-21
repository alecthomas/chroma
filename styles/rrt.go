package styles

import (
	"github.com/alecthomas/chroma"
)

// Rrt style.
var Rrt = Register(chroma.MustNewStyle("rrt", chroma.StyleEntries{
	chroma.CommentMultiline:    "#00ff00",
	chroma.CommentPreproc:      "#e5e5e5",
	chroma.CommentSingle:       "#00ff00",
	chroma.CommentSpecial:      "#00ff00",
	chroma.Comment:             "#00ff00",
	chroma.KeywordType:         "#ee82ee",
	chroma.Keyword:             "#ff0000",
	chroma.LiteralNumber:       "#ff6600",
	chroma.LiteralStringRegex:  "#87ceeb",
	chroma.LiteralStringSymbol: "#ff6600",
	chroma.LiteralString:       "#87ceeb",
	chroma.NameFunction:        "#ffff00",
	chroma.NameConstant:        "#7fffd4",
	chroma.NameAttribute:       "#f8f8f2",
	chroma.NameBuiltinPseudo:   "#f8f8f2",
	chroma.NameBuiltin:         "#f8f8f2",
	chroma.NameClass:           "#f8f8f2",
	chroma.NameNamespace:       "#f8f8f2",
	chroma.NameTag:             "#f8f8f2",
	chroma.NameVariable:        "#eedd82",
	chroma.Operator:            "#f8f8f2",
	chroma.Text:                "#f8f8f2",
	chroma.Background:          " bg:#000000",
}))
