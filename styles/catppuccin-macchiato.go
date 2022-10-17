package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	ctpMacchiatoFlamingo = "#f0c6c6"
	ctpMacchiatoPink     = "#f5bde6"
	ctpMacchiatoMauve    = "#c6a0f6"
	ctpMacchiatoRed      = "#ed8796"
	ctpMacchiatoMaroon   = "#ee99a0"
	ctpMacchiatoPeach    = "#f5a97f"
	ctpMacchiatoYellow   = "#eed49f"
	ctpMacchiatoGreen    = "#a6da95"
	ctpMacchiatoTeal     = "#8bd5ca"
	ctpMacchiatoSky      = "#91d7e3"
	ctpMacchiatoBlue     = "#8aadf4"
	ctpMacchiatoLavender = "#b7bdf8"
	ctpMacchiatoText     = "#cad3f5"
	ctpMacchiatoOverlay0 = "#6e738d"
	ctpMacchiatoSurface2 = "#5b6078"
	ctpMacchiatoSurface0 = "#363a4f"
	ctpMacchiatoBase     = "#24273a"
)

// CatppuccinMacchiato a soothing mid-saturation, mid-contrast dark pastel theme for the high-spirited
var CatppuccinMacchiato = Register(chroma.MustNewStyle("catppuccin-macchiato", chroma.StyleEntries{
	chroma.TextWhitespace:        ctpMacchiatoSurface0,
	chroma.Comment:               "italic " + ctpMacchiatoSurface2,
	chroma.CommentPreproc:        ctpMacchiatoBlue,
	chroma.Keyword:               ctpMacchiatoMauve,
	chroma.KeywordPseudo:         "bold " + ctpMacchiatoMauve,
	chroma.KeywordType:           ctpMacchiatoFlamingo,
	chroma.KeywordConstant:       "italic " + ctpMacchiatoMauve,
	chroma.Operator:              ctpMacchiatoSky,
	chroma.OperatorWord:          "bold " + ctpMacchiatoSky,
	chroma.Name:                  ctpMacchiatoLavender,
	chroma.NameBuiltin:           "italic " + ctpMacchiatoText,
	chroma.NameFunction:          ctpMacchiatoSky,
	chroma.NameClass:             ctpMacchiatoYellow,
	chroma.NameNamespace:         ctpMacchiatoYellow,
	chroma.NameException:         ctpMacchiatoMaroon,
	chroma.NameVariable:          ctpMacchiatoPeach,
	chroma.NameConstant:          ctpMacchiatoYellow,
	chroma.NameLabel:             ctpMacchiatoYellow,
	chroma.NameEntity:            ctpMacchiatoPink,
	chroma.NameAttribute:         ctpMacchiatoYellow,
	chroma.NameTag:               ctpMacchiatoMauve,
	chroma.NameDecorator:         ctpMacchiatoPink,
	chroma.NameOther:             ctpMacchiatoPeach,
	chroma.Punctuation:           ctpMacchiatoText,
	chroma.LiteralString:         ctpMacchiatoGreen,
	chroma.LiteralStringDoc:      ctpMacchiatoGreen,
	chroma.LiteralStringInterpol: ctpMacchiatoGreen,
	chroma.LiteralStringEscape:   ctpMacchiatoBlue,
	chroma.LiteralStringRegex:    ctpMacchiatoBlue,
	chroma.LiteralStringSymbol:   ctpMacchiatoGreen,
	chroma.LiteralStringOther:    ctpMacchiatoGreen,
	chroma.LiteralNumber:         ctpMacchiatoTeal,
	chroma.GenericHeading:        "bold " + ctpMacchiatoSky,
	chroma.GenericSubheading:     "bold " + ctpMacchiatoSky,
	chroma.GenericDeleted:        ctpMacchiatoMaroon,
	chroma.GenericInserted:       ctpMacchiatoGreen,
	chroma.GenericError:          ctpMacchiatoMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + ctpMacchiatoOverlay0,
	chroma.GenericOutput:         ctpMacchiatoPeach,
	chroma.GenericTraceback:      ctpMacchiatoMaroon,
	chroma.Error:                 ctpMacchiatoRed,
	chroma.Background:            ctpMacchiatoPeach + " bg:" + ctpMacchiatoBase,
}))
