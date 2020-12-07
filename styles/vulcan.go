package styles

import (
	"github.com/alecthomas/chroma"
)

/*
Inspired by the Doom One Theme
black       : "#282C34"
grey        : "#3E4452"
white       : "#C9C9C9"
red         : "#CF5967"
yellow      : "#E5C07B"
// original doom yellow: #ECBE7B
orange      : "#E69C57"
green       : "#82CC6A"
brightgreen : "#7ED95F"
cyan        : "#56B6C2"
blue        : "#7FBAF5"
purple      : "#BC74C4"
*/

var vulcan = Register(chroma.MustNewStyle("vulcan", chroma.StyleEntries{
	chroma.Comment:                  "#3E4452",
	chroma.CommentHashbang:          "#3E4452 italic",
	chroma.CommentMultiline:         "#3E4452",
	chroma.CommentPreproc:           "#7FBAF5",
	chroma.CommentSingle:            "#3E4452",
	chroma.CommentSpecial:           "#BC74C4 italic",
	chroma.Generic:                  "#C9C9C9",
	chroma.GenericDeleted:           "#CF5967",
	chroma.GenericEmph:              "#C9C9C9 underline",
	chroma.GenericError:             "#CF5967 bold",
	chroma.GenericHeading:           "#E5C07B bold",
	chroma.GenericInserted:          "#E5C07B",
	chroma.GenericOutput:            "#43454f",
	chroma.GenericPrompt:            "#C9C9C9",
	chroma.GenericStrong:            "#CF5967 bold",
	chroma.GenericSubheading:        "#CF5967 italic",
	chroma.GenericTraceback:         "#C9C9C9",
	chroma.GenericUnderline:         "underline",
	chroma.Error:                    "#CF5967",
	chroma.Keyword:                  "#7FBAF5",
	chroma.KeywordConstant:          "#CF5967 bg:#43454f",
	chroma.KeywordDeclaration:       "#7FBAF5",
	chroma.KeywordNamespace:         "#BC74C4",
	chroma.KeywordPseudo:            "#BC74C4",
	chroma.KeywordReserved:          "#7FBAF5",
	chroma.KeywordType:              "#9aedfe italic",
	chroma.Literal:                  "#C9C9C9",
	chroma.LiteralDate:              "#57c7ff",
	chroma.Name:                     "#C9C9C9",
	chroma.NameAttribute:            "#57c7ff",
	chroma.NameBuiltin:              "#7FBAF5",
	chroma.NameBuiltinPseudo:        "#7FBAF5",
	chroma.NameClass:                "#E5C07B",
	chroma.NameConstant:             "#E5C07B",
	chroma.NameDecorator:            "#E5C07B",
	chroma.NameEntity:               "#C9C9C9",
	chroma.NameException:            "#CF5967",
	chroma.NameFunction:             "#57c7ff",
	chroma.NameLabel:                "#CF5967",
	chroma.NameNamespace:            "#C9C9C9",
	chroma.NameOther:                "#C9C9C9",
	chroma.NameTag:                  "#BC74C4",
	chroma.NameVariable:             "#BC74C4 italic",
	chroma.NameVariableClass:        "#57c7ff bold",
	chroma.NameVariableGlobal:       "#E5C07B",
	chroma.NameVariableInstance:     "#57c7ff",
	chroma.LiteralNumber:            "#56B6C2",
	chroma.LiteralNumberBin:         "#BC74C4 underline",
	chroma.LiteralNumberFloat:       "#56B6C2",
	chroma.LiteralNumberHex:         "#BC74C4 underline",
	chroma.LiteralNumberInteger:     "#56B6C2",
	chroma.LiteralNumberIntegerLong: "#56B6C2",
	chroma.LiteralNumberOct:         "#BC74C4 underline",
	chroma.Operator:                 "#BC74C4",
	chroma.OperatorWord:             "#BC74C4",
	chroma.Other:                    "#C9C9C9",
	chroma.Punctuation:              "#56B6C2",
	chroma.LiteralString:            "#82CC6A",
	chroma.LiteralStringBacktick:    "#57c7ff",
	chroma.LiteralStringChar:        "#57c7ff",
	chroma.LiteralStringDoc:         "#82CC6A",
	chroma.LiteralStringDouble:      "#82CC6A",
	chroma.LiteralStringEscape:      "#82CC6A",
	chroma.LiteralStringHeredoc:     "#56B6C2",
	chroma.LiteralStringInterpol:    "#82CC6A",
	chroma.LiteralStringOther:       "#82CC6A",
	chroma.LiteralStringRegex:       "#46D9FF",
	chroma.LiteralStringSingle:      "#82CC6A",
	chroma.LiteralStringSymbol:      "#5af78e",
	chroma.Text:                     "#C9C9C9",
	chroma.TextWhitespace:           "#C9C9C9",
	chroma.Background:               " bg:#282C34",
}))
