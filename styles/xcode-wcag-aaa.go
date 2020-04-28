package styles

import (
	"github.com/alecthomas/chroma"
)

// XcodeHighContrast style.
var XcodeHighContrast = Register(chroma.MustNewStyle("xcode-hc", chroma.StyleEntries{
	chroma.Comment:           "#105F00",
	chroma.CommentPreproc:    "#633820",
	chroma.LiteralString:     "#B50701",
	chroma.LiteralStringChar: "#2300CE",
	chroma.Operator:          "#000000",
	chroma.Keyword:           "#9A0384",
	chroma.Name:              "#000000",
	chroma.NameAttribute:     "#614E02",
	chroma.NameClass:         "#255B62",
	chroma.NameFunction:      "#000000",
	chroma.NameBuiltin:       "#9A0384",
	chroma.NameBuiltinPseudo: "#5B269A",
	chroma.NameVariable:      "#000000",
	chroma.NameTag:           "#000000",
	chroma.NameDecorator:     "#000000",
	chroma.NameLabel:         "#000000",
	chroma.Literal:           "#1C01CE",
	chroma.LiteralNumber:     "#1C01CE",
	chroma.Error:             "#000000",
	chroma.Background:        " bg:#ffffff",
}))
