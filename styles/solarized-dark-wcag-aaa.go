package styles

import (
	"github.com/alecthomas/chroma"
)

// SolarizedDarkHighContrast style.
var SolarizedDarkHighContrast = Register(chroma.MustNewStyle("solarized-dark-hc", chroma.StyleEntries{
	chroma.Keyword:               "#97C44F",
	chroma.KeywordConstant:       "#FF9783",
	chroma.KeywordDeclaration:    "#77BAFF",
	chroma.KeywordReserved:       "#77BAFF",
	chroma.KeywordType:           "#FF9594",
	chroma.NameAttribute:         "#ACBABA",
	chroma.NameBuiltin:           "#DBAF4E",
	chroma.NameBuiltinPseudo:     "#77BAFF",
	chroma.NameClass:             "#77BAFF",
	chroma.NameConstant:          "#FF9783",
	chroma.NameDecorator:         "#77BAFF",
	chroma.NameEntity:            "#FF9783",
	chroma.NameException:         "#FF9783",
	chroma.NameFunction:          "#77BAFF",
	chroma.NameTag:               "#77BAFF",
	chroma.NameVariable:          "#77BAFF",
	chroma.LiteralString:         "#60C7BD",
	chroma.LiteralStringBacktick: "#A5BAC1",
	chroma.LiteralStringChar:     "#60C7BD",
	chroma.LiteralStringDoc:      "#ACBABA",
	chroma.LiteralStringEscape:   "#FF9783",
	chroma.LiteralStringHeredoc:  "#ACBABA",
	chroma.LiteralStringRegex:    "#FF9594",
	chroma.LiteralNumber:         "#60C7BD",
	chroma.Operator:              "#97C44F",
	chroma.Comment:               "#A5BAC1",
	chroma.CommentPreproc:        "#97C44F",
	chroma.CommentSpecial:        "#97C44F",
	chroma.GenericDeleted:        "#FF9594",
	chroma.GenericEmph:           "italic",
	chroma.GenericError:          "#FF9594 bold",
	chroma.GenericHeading:        "#FF9783",
	chroma.GenericInserted:       "#97C44F",
	chroma.GenericStrong:         "bold",
	chroma.GenericSubheading:     "#77BAFF",
	chroma.Background:            "#93A1A1 bg:#002B36",
	chroma.Other:                 "#FF9783",
}))
