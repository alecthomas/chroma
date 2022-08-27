package styles

import (
	"github.com/alecthomas/chroma/v2"
)

// Modus Vivendi (dark) style.
var ModusVivendi = Register(chroma.MustNewStyle("modus-vivendi", chroma.StyleEntries{
	chroma.Keyword:         "#b6a0ff",
	chroma.KeywordConstant: "#00bcff",
	chroma.KeywordType:     "#6ae4b9",
	chroma.Comment:         "#a8a8a8",
	chroma.NameVariable:    "#00d3d0",
	chroma.Operator:        "#00d3d0",
	chroma.NameFunction:    "#feacd0",
	chroma.NameBuiltin:     "#f78fe7",
	chroma.Literal:         "#00bcff",
	chroma.String:          "#79a8ff",
	chroma.Background:      "#ffffff bg:#000000",
}))
