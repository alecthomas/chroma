package styles

import (
	"github.com/alecthomas/chroma"
)

// RrtHighContrast style.
var RrtHighContrast = Register(chroma.MustNewStyle("rrt-hc", chroma.StyleEntries{
	chroma.CommentPreproc:      "#e5e5e5",
	chroma.Comment:             "#00ff00",
	chroma.KeywordType:         "#ee82ee",
	chroma.Keyword:             "#FF6060",
	chroma.LiteralNumber:       "#ff6600",
	chroma.LiteralStringSymbol: "#ff6600",
	chroma.LiteralString:       "#87ceeb",
	chroma.NameFunction:        "#ffff00",
	chroma.NameConstant:        "#7fffd4",
	chroma.NameVariable:        "#eedd82",
	chroma.Background:          "#f8f8f2 bg:#000000",
}))
