package styles

import (
	"github.com/alecthomas/chroma"
)

// SolarizedDark256-hc style.
var SolarizedDark256-hc = Register(chroma.MustNewStyle("solarized-dark256-hc", chroma.StyleEntries{
	chroma.Keyword:               "#89B351",
	chroma.KeywordConstant:       "#FF8D5E",
	chroma.KeywordDeclaration:    "#77A6FF",
	chroma.KeywordNamespace:      "#FF8D5E",
	chroma.KeywordReserved:       "#77A6FF",
	chroma.KeywordType:           "#FF8585",
	chroma.NameAttribute:         "#A8A8A8",
	chroma.NameBuiltin:           "#77A6FF",
	chroma.NameBuiltinPseudo:     "#77A6FF",
	chroma.NameClass:             "#77A6FF",
	chroma.NameConstant:          "#FF8D5E",
	chroma.NameDecorator:         "#77A6FF",
	chroma.NameEntity:            "#FF8D5E",
	chroma.NameException:         "#CDA543",
	chroma.NameFunction:          "#77A6FF",
	chroma.NameTag:               "#77A6FF",
	chroma.NameVariable:          "#77A6FF",
	chroma.LiteralString:         "#27B8B8",
	chroma.LiteralStringBacktick: "#AEAEAE",
	chroma.LiteralStringChar:     "#27B8B8",
	chroma.LiteralStringDoc:      "#27B8B8",
	chroma.LiteralStringEscape:   "#FF8585",
	chroma.LiteralStringHeredoc:  "#27B8B8",
	chroma.LiteralStringRegex:    "#FF8585",
	chroma.LiteralNumber:         "#27B8B8",
	chroma.Operator:              "#A8A8A8",
	chroma.OperatorWord:          "#89B351",
	chroma.Comment:               "#AEAEAE",
	chroma.CommentPreproc:        "#89B351",
	chroma.CommentSpecial:        "#89B351",
	chroma.GenericDeleted:        "#FF8585",
	chroma.GenericEmph:           "italic",
	chroma.GenericError:          "#FF8585 bold",
	chroma.GenericHeading:        "#FF8D5E",
	chroma.GenericInserted:       "#89B351",
	chroma.GenericStrong:         "bold",
	chroma.GenericSubheading:     "#77A6FF",
	chroma.Background:            "#8a8a8a bg:#1c1c1c",
	chroma.Other:                 "#FF8D5E",
}))
