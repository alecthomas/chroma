package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	ctpLattePink     = "#ea76cb"
	ctpLatteMauve    = "#8839ef"
	ctpLatteRed      = "#d20f39"
	ctpLatteMaroon   = "#e64553"
	ctpLattePeach    = "#fe640b"
	ctpLatteYellow   = "#df8e1d"
	ctpLatteGreen    = "#40a02b"
	ctpLatteSky      = "#04a5e5"
	ctpLatteBlue     = "#1e66f5"
	ctpLatteLavender = "#7287fd"
	ctpLatteText     = "#4c4f69"
	ctpLatteOverlay0 = "#9ca0b0"
	ctpLatteSurface2 = "#acb0be"
	ctpLatteSurface0 = "#ccd0da"
	ctpLatteBase     = "#eff1f5"
)

// CatppuccinLatte a soothing light pastel theme for the high-spirited
var CatppuccinLatte = Register(chroma.MustNewStyle("catppuccin-latte", chroma.StyleEntries{
	chroma.TextWhitespace:        ctpLatteSurface0,
	chroma.Comment:               "italic " + ctpLatteSurface2,
	chroma.CommentPreproc:        ctpLatteBlue,
	chroma.Keyword:               ctpLatteMauve,
	chroma.KeywordPseudo:         "bold " + ctpLatteMauve,
	chroma.KeywordType:           ctpLatteYellow,
	chroma.KeywordConstant:       "italic " + ctpLatteMauve,
	chroma.Operator:              ctpLatteSky,
	chroma.OperatorWord:          "bold " + ctpLatteSky,
	chroma.Name:                  ctpLatteLavender,
	chroma.NameBuiltin:           "italic " + ctpLattePeach,
	chroma.NameFunction:          ctpLatteSky,
	chroma.NameClass:             ctpLatteYellow,
	chroma.NameNamespace:         ctpLatteYellow,
	chroma.NameException:         ctpLatteMaroon,
	chroma.NameVariable:          ctpLattePeach,
	chroma.NameConstant:          ctpLatteYellow,
	chroma.NameLabel:             ctpLatteYellow,
	chroma.NameEntity:            ctpLattePink,
	chroma.NameAttribute:         ctpLatteYellow,
	chroma.NameTag:               ctpLatteMauve,
	chroma.NameDecorator:         ctpLattePink,
	chroma.NameOther:             ctpLattePeach,
	chroma.Punctuation:           ctpLatteText,
	chroma.LiteralString:         ctpLatteGreen,
	chroma.LiteralStringDoc:      ctpLatteGreen,
	chroma.LiteralStringInterpol: ctpLatteGreen,
	chroma.LiteralStringEscape:   ctpLatteBlue,
	chroma.LiteralStringRegex:    ctpLatteBlue,
	chroma.LiteralStringSymbol:   ctpLatteGreen,
	chroma.LiteralStringOther:    ctpLatteGreen,
	chroma.LiteralNumber:         ctpLattePeach,
	chroma.GenericHeading:        "bold " + ctpLatteSky,
	chroma.GenericSubheading:     "bold " + ctpLatteSky,
	chroma.GenericDeleted:        ctpLatteMaroon,
	chroma.GenericInserted:       ctpLatteGreen,
	chroma.GenericError:          ctpLatteMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + ctpLatteOverlay0,
	chroma.GenericOutput:         ctpLattePeach,
	chroma.GenericTraceback:      ctpLatteMaroon,
	chroma.Error:                 ctpLatteRed,
	chroma.Background:            ctpLattePeach + " bg:" + ctpLatteBase,
}))
