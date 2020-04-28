package styles

import (
	"github.com/alecthomas/chroma"
)

// Colorful-hc style.
var Colorful-hc = Register(chroma.MustNewStyle("colorful-hc", chroma.StyleEntries{
	chroma.TextWhitespace:        "#555555",
	chroma.Comment:               "#585858",
	chroma.CommentPreproc:        "#2E577A",
	chroma.CommentSpecial:        "bold #A70202",
	chroma.Keyword:               "bold #026302",
	chroma.KeywordPseudo:         "#003388",
	chroma.KeywordType:           "#333399",
	chroma.Operator:              "#333333",
	chroma.OperatorWord:          "bold #000000",
	chroma.NameBuiltin:           "#03651D",
	chroma.NameFunction:          "bold #005399",
	chroma.NameClass:             "bold #A9005C",
	chroma.NameNamespace:         "bold #025F85",
	chroma.NameException:         "bold #A80303",
	chroma.NameVariable:          "#754701",
	chroma.NameVariableInstance:  "#3333BB",
	chroma.NameVariableClass:     "#225C8F",
	chroma.NameVariableGlobal:    "bold #824300",
	chroma.NameConstant:          "bold #003366",
	chroma.NameLabel:             "bold #6F5504",
	chroma.NameEntity:            "bold #880000",
	chroma.NameAttribute:         "#0000CC",
	chroma.NameTag:               "#056005",
	chroma.NameDecorator:         "bold #555555",
	chroma.LiteralString:         "bg:#fff0f0",
	chroma.LiteralStringChar:     "#0044DD bg:",
	chroma.LiteralStringDoc:      "#A62900 bg:",
	chroma.LiteralStringInterpol: "bg:#eee",
	chroma.LiteralStringEscape:   "bold #535353",
	chroma.LiteralStringRegex:    "bg:#fff0ff #000000",
	chroma.LiteralStringSymbol:   "#7C4900 bg:",
	chroma.LiteralStringOther:    "#A31600",
	chroma.LiteralNumber:         "bold #6600EE",
	chroma.LiteralNumberInteger:  "bold #0000DD",
	chroma.LiteralNumberFloat:    "bold #6600EE",
	chroma.LiteralNumberHex:      "bold #005588",
	chroma.LiteralNumberOct:      "bold #4400EE",
	chroma.GenericHeading:        "bold #000080",
	chroma.GenericSubheading:     "bold #800080",
	chroma.GenericDeleted:        "#A00000",
	chroma.GenericInserted:       "#056705",
	chroma.GenericError:          "#A80303",
	chroma.GenericEmph:           "italic",
	chroma.GenericStrong:         "bold",
	chroma.GenericPrompt:         "bold #914101",
	chroma.GenericOutput:         "#585858",
	chroma.GenericTraceback:      "#0044DD",
	chroma.GenericUnderline:      "underline",
	chroma.Error:                 "#6B0000 bg:#FAA",
	chroma.Background:            " bg:#ffffff",
}))
