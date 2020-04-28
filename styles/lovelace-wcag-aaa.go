package styles

import (
	"github.com/alecthomas/chroma"
)

// LovelaceHighContrast style.
var LovelaceHighContrast = Register(chroma.MustNewStyle("lovelace-hc", chroma.StyleEntries{
	chroma.TextWhitespace:         "#635202",
	chroma.Comment:                "italic #585858",
	chroma.CommentHashbang:        "#025B72",
	chroma.CommentMultiline:       "#585858",
	chroma.CommentPreproc:         "noitalic #026346",
	chroma.Keyword:                "#2838b0",
	chroma.KeywordConstant:        "italic #444444",
	chroma.KeywordDeclaration:     "italic",
	chroma.KeywordType:            "italic",
	chroma.Operator:               "#535353",
	chroma.OperatorWord:           "#8C188C",
	chroma.Punctuation:            "#585858",
	chroma.NameAttribute:          "#036003",
	chroma.NameBuiltin:            "#036003",
	chroma.NameBuiltinPseudo:      "italic",
	chroma.NameClass:              "#025B72",
	chroma.NameConstant:           "#8A3D00",
	chroma.NameDecorator:          "#025B72",
	chroma.NameEntity:             "#465F00",
	chroma.NameException:          "#5E5803",
	chroma.NameFunction:           "#6E4E35",
	chroma.NameFunctionMagic:      "#8A3D00",
	chroma.NameLabel:              "#026346",
	chroma.NameNamespace:          "#026346",
	chroma.NameTag:                "#2838b0",
	chroma.NameVariable:           "#992828",
	chroma.NameVariableGlobal:     "#5E5803",
	chroma.NameVariableMagic:      "#8A3D00",
	chroma.LiteralString:          "#A11E1E",
	chroma.LiteralStringAffix:     "#444444",
	chroma.LiteralStringChar:      "#8C188C",
	chroma.LiteralStringDelimiter: "#8A3D00",
	chroma.LiteralStringDoc:       "italic #8A3D00",
	chroma.LiteralStringEscape:    "#465F00",
	chroma.LiteralStringInterpol:  "underline",
	chroma.LiteralStringOther:     "#8C188C",
	chroma.LiteralStringRegex:     "#8C188C",
	chroma.LiteralNumber:          "#444444",
	chroma.GenericDeleted:         "#A70202",
	chroma.GenericEmph:            "italic",
	chroma.GenericError:           "#A70202",
	chroma.GenericHeading:         "#535353",
	chroma.GenericSubheading:      "#444444",
	chroma.GenericInserted:        "#036003",
	chroma.GenericOutput:          "#535353",
	chroma.GenericPrompt:          "#444444",
	chroma.GenericStrong:          "bold",
	chroma.GenericTraceback:       "#2838b0",
	chroma.GenericUnderline:       "underline",
	chroma.Error:                  "bg:#a848a8",
	chroma.Background:             " bg:#ffffff",
}))
