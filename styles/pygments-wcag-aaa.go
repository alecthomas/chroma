package styles

import (
	"github.com/alecthomas/chroma"
)

// Pygments-hc default theme.
var Pygments-hc = Register(chroma.MustNewStyle("pygments-hc", chroma.StyleEntries{
	chroma.Whitespace:     "#555555",
	chroma.Comment:        "italic #075F5F",
	chroma.CommentPreproc: "noitalic #7A4E03",

	chroma.Keyword:       "bold #056705",
	chroma.KeywordPseudo: "nobold",
	chroma.KeywordType:   "nobold #B00040",

	chroma.Operator:     "#535353",
	chroma.OperatorWord: "bold #8100C3",

	chroma.NameBuiltin:   "#056705",
	chroma.NameFunction:  "#0000FF",
	chroma.NameClass:     "bold #0000FF",
	chroma.NameNamespace: "bold #0000FF",
	chroma.NameException: "bold #A91500",
	chroma.NameVariable:  "#19177C",
	chroma.NameConstant:  "#880000",
	chroma.NameLabel:     "#545400",
	chroma.NameEntity:    "bold #585858",
	chroma.NameAttribute: "#505D02",
	chroma.NameTag:       "bold #056705",
	chroma.NameDecorator: "#8100C3",

	chroma.String:         "#AE0D0D",
	chroma.StringDoc:      "italic",
	chroma.StringInterpol: "bold #8B345A",
	chroma.StringEscape:   "bold #7E4002",
	chroma.StringRegex:    "#8B345A",
	chroma.StringSymbol:   "#19177C",
	chroma.StringOther:    "#056705",
	chroma.Number:         "#535353",

	chroma.GenericHeading:    "bold #000080",
	chroma.GenericSubheading: "bold #800080",
	chroma.GenericDeleted:    "#A00000",
	chroma.GenericInserted:   "#056705",
	chroma.GenericError:      "#A80303",
	chroma.GenericEmph:       "italic",
	chroma.GenericStrong:     "bold",
	chroma.GenericPrompt:     "bold #000080",
	chroma.GenericOutput:     "#585858",
	chroma.GenericTraceback:  "#0044DD",
	chroma.GenericUnderline:  "underline",

	chroma.Error: "border:#FF0000",
}))
