package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	ctpFrappeFlamingo = "#eebebe"
	ctpFrappePink     = "#f4b8e4"
	ctpFrappeMauve    = "#ca9ee6"
	ctpFrappeRed      = "#e78284"
	ctpFrappeMaroon   = "#ea999c"
	ctpFrappePeach    = "#ef9f76"
	ctpFrappeYellow   = "#e5c890"
	ctpFrappeGreen    = "#a6d189"
	ctpFrappeTeal     = "#81c8be"
	ctpFrappeSky      = "#99d1db"
	ctpFrappeBlue     = "#8caaee"
	ctpFrappeLavender = "#babbf1"
	ctpFrappeText     = "#c6d0f5"
	ctpFrappeOverlay0 = "#737994"
	ctpFrappeSurface2 = "#626880"
	ctpFrappeSurface0 = "#414559"
	ctpFrappeBase     = "#303446"
)

// CatppuccinFrappe a soothing low-saturation, low-contrast dark pastel theme for the high-spirited
var CatppuccinFrappe = Register(chroma.MustNewStyle("catppuccin-frappe", chroma.StyleEntries{
	chroma.TextWhitespace:        ctpFrappeSurface0,
	chroma.Comment:               "italic " + ctpFrappeSurface2,
	chroma.CommentPreproc:        ctpFrappeBlue,
	chroma.Keyword:               ctpFrappeMauve,
	chroma.KeywordPseudo:         "bold " + ctpFrappeMauve,
	chroma.KeywordType:           ctpFrappeFlamingo,
	chroma.KeywordConstant:       "italic " + ctpFrappeMauve,
	chroma.Operator:              ctpFrappeSky,
	chroma.OperatorWord:          "bold " + ctpFrappeSky,
	chroma.Name:                  ctpFrappeLavender,
	chroma.NameBuiltin:           "italic " + ctpFrappeText,
	chroma.NameFunction:          ctpFrappeSky,
	chroma.NameClass:             ctpFrappeYellow,
	chroma.NameNamespace:         ctpFrappeYellow,
	chroma.NameException:         ctpFrappeMaroon,
	chroma.NameVariable:          ctpFrappePeach,
	chroma.NameConstant:          ctpFrappeYellow,
	chroma.NameLabel:             ctpFrappeYellow,
	chroma.NameEntity:            ctpFrappePink,
	chroma.NameAttribute:         ctpFrappeYellow,
	chroma.NameTag:               ctpFrappeMauve,
	chroma.NameDecorator:         ctpFrappePink,
	chroma.NameOther:             ctpFrappePeach,
	chroma.Punctuation:           ctpFrappeText,
	chroma.LiteralString:         ctpFrappeGreen,
	chroma.LiteralStringDoc:      ctpFrappeGreen,
	chroma.LiteralStringInterpol: ctpFrappeGreen,
	chroma.LiteralStringEscape:   ctpFrappeBlue,
	chroma.LiteralStringRegex:    ctpFrappeBlue,
	chroma.LiteralStringSymbol:   ctpFrappeGreen,
	chroma.LiteralStringOther:    ctpFrappeGreen,
	chroma.LiteralNumber:         ctpFrappeTeal,
	chroma.GenericHeading:        "bold " + ctpFrappeSky,
	chroma.GenericSubheading:     "bold " + ctpFrappeSky,
	chroma.GenericDeleted:        ctpFrappeMaroon,
	chroma.GenericInserted:       ctpFrappeGreen,
	chroma.GenericError:          ctpFrappeMaroon,
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold " + ctpFrappeOverlay0,
	chroma.GenericOutput:         ctpFrappePeach,
	chroma.GenericTraceback:      ctpFrappeMaroon,
	chroma.Error:                 ctpFrappeRed,
	chroma.Background:            ctpFrappePeach + " bg:" + ctpFrappeBase,
}))
