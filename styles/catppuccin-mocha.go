package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	ctpMochaFlamingo = "#f2cdcd"
	ctpMochaPink     = "#f5c2e7"
	ctpMochaMauve    = "#cba6f7"
	ctpMochaRed      = "#f38ba8"
	ctpMochaMaroon   = "#eba0ac"
	ctpMochaPeach    = "#fab387"
	ctpMochaYellow   = "#f9e2af"
	ctpMochaGreen    = "#a6e3a1"
	ctpMochaTeal     = "#94e2d5"
	ctpMochaSky      = "#89dceb"
	ctpMochaBlue     = "#89b4fa"
	ctpMochaLavender = "#b4befe"
	ctpMochaText     = "#cdd6f4"
	ctpMochaOverlay0 = "#6c7086"
	ctpMochaSurface2 = "#585b70"
	ctpMochaSurface0 = "#313244"
	ctpMochaBase     = "#1e1e2e"
)

// CatppuccinMocha a soothing high-saturation, high-contrast dark pastel theme for the high-spirited
var CatppuccinMocha = Register(chroma.MustNewStyle("catppuccin-mocha", chroma.StyleEntries{
	chroma.TextWhitespace:        ctpMochaSurface0,
	chroma.Comment:               "italic " + ctpMochaSurface2,
	chroma.CommentPreproc:        ctpMochaBlue,
	chroma.Keyword:               ctpMochaMauve,
	chroma.KeywordPseudo:         "bold " + ctpMochaMauve,
	chroma.KeywordType:           ctpMochaYellow,
	chroma.KeywordConstant:       "italic " + ctpMochaMauve,
	chroma.Operator:              ctpMochaSky,
	chroma.OperatorWord:          "bold " + ctpMochaSky,
	chroma.Name:                  ctpMochaLavender,
	chroma.NameBuiltin:           "italic " + ctpMochaPeach,
	chroma.NameFunction:          ctpMochaSky,
	chroma.NameClass:             ctpMochaYellow,
	chroma.NameNamespace:         ctpMochaYellow,
	chroma.NameException:         ctpMochaMaroon,
	chroma.NameVariable:          ctpMochaPeach,
	chroma.NameConstant:          ctpMochaYellow,
	chroma.NameLabel:             ctpMochaYellow,
	chroma.NameEntity:            ctpMochaPink,
	chroma.NameAttribute:         ctpMochaYellow,
	chroma.NameTag:               ctpMochaMauve,
	chroma.NameDecorator:         ctpMochaPink,
	chroma.NameOther:             ctpMochaPeach,
	chroma.Punctuation:           ctpMochaText,
	chroma.LiteralString:         ctpMochaGreen,
	chroma.LiteralStringDoc:      ctpMochaGreen,
	chroma.LiteralStringInterpol: ctpMochaGreen,
	chroma.LiteralStringEscape:   ctpMochaBlue,
	chroma.LiteralStringRegex:    ctpMochaBlue,
	chroma.LiteralStringSymbol:   ctpMochaGreen,
	chroma.LiteralStringOther:    ctpMochaGreen,
	chroma.LiteralNumber:         ctpMochaPeach,
	chroma.GenericHeading:        "bold " + ctpMochaSky,
	chroma.GenericSubheading:     "bold " + ctpMochaSky,
	chroma.GenericDeleted:        ctpMochaMaroon,
	chroma.GenericInserted:       ctpMochaGreen,
	chroma.GenericError:          ctpMochaMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + ctpMochaOverlay0,
	chroma.GenericOutput:         ctpMochaPeach,
	chroma.GenericTraceback:      ctpMochaMaroon,
	chroma.Error:                 ctpMochaRed,
	chroma.Background:            ctpMochaPeach + " bg:" + ctpMochaBase,
}))
