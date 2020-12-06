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
orange      : "#E69C57"
green       : "#82CC6A"
brightgreen : "#7ed95f"
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
	chroma.Generic:                  "#e2e4e5",
	chroma.GenericDeleted:           "#ff5c57",
	chroma.GenericEmph:              "#e2e4e5 underline",
	chroma.GenericError:             "#ff5c57",
	chroma.GenericHeading:           "#e2e4e5 bold",
	chroma.GenericInserted:          "#e2e4e5 bold",
	chroma.GenericOutput:            "#43454f",
	chroma.GenericPrompt:            "#e2e4e5",
	chroma.GenericStrong:            "#e2e4e5 italic",
	chroma.GenericSubheading:        "#e2e4e5 bold",
	chroma.GenericTraceback:         "#e2e4e5",
	chroma.GenericUnderline:         "underline",
	chroma.Error:                    "#ff5c57",
	chroma.Keyword:                  "#ff6ac1",
	chroma.KeywordConstant:          "#ff6ac1",
	chroma.KeywordDeclaration:       "#ff5c57",
	chroma.KeywordNamespace:         "#ff6ac1",
	chroma.KeywordPseudo:            "#ff6ac1",
	chroma.KeywordReserved:          "#ff6ac1",
	chroma.KeywordType:              "#9aedfe",
	chroma.Literal:                  "#e2e4e5",
	chroma.LiteralDate:              "#e2e4e5",
	chroma.Name:                     "#e2e4e5",
	chroma.NameAttribute:            "#57c7ff",
	chroma.NameBuiltin:              "#ff5c57",
	chroma.NameBuiltinPseudo:        "#e2e4e5",
	chroma.NameClass:                "#f3f99d",
	chroma.NameConstant:             "#ff9f43",
	chroma.NameDecorator:            "#ff9f43",
	chroma.NameEntity:               "#e2e4e5",
	chroma.NameException:            "#e2e4e5",
	chroma.NameFunction:             "#57c7ff",
	chroma.NameLabel:                "#ff5c57",
	chroma.NameNamespace:            "#e2e4e5",
	chroma.NameOther:                "#e2e4e5",
	chroma.NameTag:                  "#ff6ac1",
	chroma.NameVariable:             "#ff5c57",
	chroma.NameVariableClass:        "#ff5c57",
	chroma.NameVariableGlobal:       "#ff5c57",
	chroma.NameVariableInstance:     "#ff5c57",
	chroma.LiteralNumber:            "#ff9f43",
	chroma.LiteralNumberBin:         "#ff9f43",
	chroma.LiteralNumberFloat:       "#ff9f43",
	chroma.LiteralNumberHex:         "#ff9f43",
	chroma.LiteralNumberInteger:     "#ff9f43",
	chroma.LiteralNumberIntegerLong: "#ff9f43",
	chroma.LiteralNumberOct:         "#ff9f43",
	chroma.Operator:                 "#ff6ac1",
	chroma.OperatorWord:             "#ff6ac1",
	chroma.Other:                    "#e2e4e5",
	chroma.Punctuation:              "#e2e4e5",
	chroma.LiteralString:            "#5af78e",
	chroma.LiteralStringBacktick:    "#5af78e",
	chroma.LiteralStringChar:        "#5af78e",
	chroma.LiteralStringDoc:         "#5af78e",
	chroma.LiteralStringDouble:      "#5af78e",
	chroma.LiteralStringEscape:      "#5af78e",
	chroma.LiteralStringHeredoc:     "#5af78e",
	chroma.LiteralStringInterpol:    "#5af78e",
	chroma.LiteralStringOther:       "#5af78e",
	chroma.LiteralStringRegex:       "#5af78e",
	chroma.LiteralStringSingle:      "#5af78e",
	chroma.LiteralStringSymbol:      "#5af78e",
	chroma.Text:                     "#e2e4e5",
	chroma.TextWhitespace:           "#e2e4e5",
	chroma.Background:               " bg:#282a36",
}))
