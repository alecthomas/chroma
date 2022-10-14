package styles

import "github.com/alecthomas/chroma/v2"

var (
	ctpFlamingo = "#F2CDCD"
	ctpPink     = "#F5C2E7"
	ctpMauve    = "#CBA6F7"
	ctpRed      = "#F38BA8"
	ctpMaroon   = "#EBA0AC"
	ctpPeach    = "#FAB387"
	ctpYellow   = "#F9E2AF"
	ctpGreen    = "#A6E3A1"
	ctpTeal     = "#94E2D5"
	ctpSky      = "#89DCEB"
	ctpBlue     = "#87B0F9"
	ctpLavender = "#B4BEFE"
	ctpText     = "#C6D0F5"
	ctpOverlay0 = "#696D86"
	ctpSurface2 = "#565970"
	ctpSurface0 = "#313244"
	ctpBase     = "#1E1E2E"
)

// Catppuccin a soothing pastel theme for the high-spirited (mocha variant)
var Catppuccin = Register(chroma.MustNewStyle("catppuccin", chroma.StyleEntries{
	chroma.TextWhitespace:        ctpSurface0,
	chroma.Comment:               "italic " + ctpSurface2,
	chroma.CommentPreproc:        ctpBlue,
	chroma.Keyword:               ctpMauve,
	chroma.KeywordPseudo:         "bold " + ctpMauve,
	chroma.KeywordType:           ctpFlamingo,
	chroma.KeywordConstant:       "italic " + ctpMauve,
	chroma.Operator:              ctpSky,
	chroma.OperatorWord:          "bold " + ctpSky,
	chroma.Name:                  ctpLavender,
	chroma.NameBuiltin:           "italic " + ctpText,
	chroma.NameFunction:          ctpSky,
	chroma.NameClass:             ctpYellow,
	chroma.NameNamespace:         ctpYellow,
	chroma.NameException:         ctpMaroon,
	chroma.NameVariable:          ctpPeach,
	chroma.NameConstant:          ctpYellow,
	chroma.NameLabel:             ctpYellow,
	chroma.NameEntity:            ctpPink,
	chroma.NameAttribute:         ctpYellow,
	chroma.NameTag:               ctpMauve,
	chroma.NameDecorator:         ctpPink,
	chroma.NameOther:             ctpPeach,
	chroma.Punctuation:           ctpText,
	chroma.LiteralString:         ctpGreen,
	chroma.LiteralStringDoc:      ctpGreen,
	chroma.LiteralStringInterpol: ctpGreen,
	chroma.LiteralStringEscape:   ctpBlue,
	chroma.LiteralStringRegex:    ctpBlue,
	chroma.LiteralStringSymbol:   ctpGreen,
	chroma.LiteralStringOther:    ctpGreen,
	chroma.LiteralNumber:         ctpTeal,
	chroma.GenericHeading:        "bold " + ctpSky,
	chroma.GenericSubheading:     "bold " + ctpSky,
	chroma.GenericDeleted:        ctpMaroon,
	chroma.GenericInserted:       ctpGreen,
	chroma.GenericError:          ctpMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + ctpOverlay0,
	chroma.GenericOutput:         ctpPeach,
	chroma.GenericTraceback:      ctpMaroon,
	chroma.Error:                 ctpRed,
	chroma.Background:            ctpPeach + " bg:" + ctpBase,
}))
