package styles

import (
	"github.com/alecthomas/chroma"
)

// AbapHighContrast style.
var AbapHighContrast = Register(chroma.MustNewStyle("abap-hc", chroma.StyleEntries{
	chroma.Comment:        "italic #585858",
	chroma.CommentSpecial: "#585858",
	chroma.Keyword:        "#0000ff",
	chroma.OperatorWord:   "#0000ff",
	chroma.Name:           "#000000",
	chroma.LiteralNumber:  "#01588A",
	chroma.LiteralString:  "#2A6300",
	chroma.Error:          "#A80303",
	chroma.Background:     " bg:#ffffff",
}))
