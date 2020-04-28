package styles

import (
	"github.com/alecthomas/chroma"
)

// SolarizedLight-hc style.
var SolarizedLight-hc = Register(chroma.MustNewStyle("solarized-light-hc", chroma.StyleEntries{
	chroma.Text:             "bg:#eee8d5 #304951",
	chroma.Keyword:          "#445001",
	chroma.KeywordConstant:  "bold",
	chroma.KeywordNamespace: "#9C0701 bold",
	chroma.KeywordType:      "bold",
	chroma.Name:             "#034974",
	chroma.NameBuiltin:      "#7D2800",
	chroma.NameClass:        "#7D2800",
	chroma.NameTag:          "bold",
	chroma.Literal:          "#02554F",
	chroma.LiteralNumber:    "bold",
	chroma.OperatorWord:     "#445001",
	chroma.Comment:          "#3C4D4D italic",
	chroma.Generic:          "#8A004F",
	chroma.Background:       " bg:#eee8d5",
}))
