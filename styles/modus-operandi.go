package styles

import (
	"github.com/alecthomas/chroma/v2"
)

// Modus Operandi (light) style.
var ModusOperandi = Register(chroma.MustNewStyle("modus-operandi", chroma.StyleEntries{
	chroma.Keyword:         "#5317ac",
	chroma.KeywordConstant: "#0000c0",
	chroma.KeywordType:     "#005a5f",
	chroma.Comment:         "#505050",
	chroma.NameVariable:    "#00538b",
	chroma.Operator:        "#00538b",
	chroma.NameFunction:    "#721045",
	chroma.NameBuiltin:     "#8f0075",
	chroma.Literal:         "#0000c0",
	chroma.String:          "#2544bb",
	chroma.Background:      "#000000 bg:#ffffff",
}))
