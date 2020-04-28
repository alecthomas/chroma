package styles

import (
	"github.com/alecthomas/chroma"
)

// ArduinoHighContrast style.
var ArduinoHighContrast = Register(chroma.MustNewStyle("arduino-hc", chroma.StyleEntries{
	chroma.Error:           "#a61717",
	chroma.Comment:         "#445657",
	chroma.CommentPreproc:  "#495C00",
	chroma.Keyword:         "#495C00",
	chroma.KeywordConstant: "#096264",
	chroma.KeywordPseudo:   "#096264",
	chroma.KeywordReserved: "#096264",
	chroma.KeywordType:     "#096264",
	chroma.Operator:        "#495C00",
	chroma.Name:            "#434f54",
	chroma.NameBuiltin:     "#495C00",
	chroma.NameFunction:    "#8C3500",
	chroma.NameOther:       "#495C00",
	chroma.LiteralNumber:   "#5E4F1D",
	chroma.LiteralString:   "#4D5B5C",
	chroma.Background:      " bg:#ffffff",
}))
