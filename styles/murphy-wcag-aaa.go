package styles

import (
	"github.com/alecthomas/chroma"
)

// Murphy-hc style.
var Murphy-hc = Register(chroma.MustNewStyle("murphy-hc", chroma.StyleEntries{
	chroma.TextWhitespace:        "#555555",
	chroma.Comment:               "#535353 italic",
	chroma.CommentPreproc:        "#2E577A noitalic",
	chroma.CommentSpecial:        "#A70202 bold",
	chroma.Keyword:               "bold #065965",
	chroma.KeywordPseudo:         "#0458A7",
	chroma.KeywordType:           "#3939EA",
	chroma.Operator:              "#333333",
	chroma.OperatorWord:          "bold #000000",
	chroma.NameBuiltin:           "#02601B",
	chroma.NameFunction:          "bold #096359",
	chroma.NameClass:             "bold #8B248B",
	chroma.NameNamespace:         "bold #025F85",
	chroma.NameException:         "bold #A80303",
	chroma.NameVariable:          "#003366",
	chroma.NameVariableInstance:  "#4747A2",
	chroma.NameVariableClass:     "#505083",
	chroma.NameVariableGlobal:    "#8C4101",
	chroma.NameConstant:          "bold #096359",
	chroma.NameLabel:             "bold #6F5504",
	chroma.NameEntity:            "#880000",
	chroma.NameAttribute:         "#000077",
	chroma.NameTag:               "#056005",
	chroma.NameDecorator:         "bold #555555",
	chroma.LiteralString:         "bg:#e0e0ff",
	chroma.LiteralStringChar:     "#4141C8 bg:",
	chroma.LiteralStringDoc:      "#A62900 bg:",
	chroma.LiteralStringInterpol: "bg:#eee",
	chroma.LiteralStringEscape:   "bold #535353",
	chroma.LiteralStringRegex:    "bg:#e0e0ff #000000",
	chroma.LiteralStringSymbol:   "#725100 bg:",
	chroma.LiteralStringOther:    "#9F2E2E",
	chroma.LiteralNumber:         "bold #6600EE",
	chroma.LiteralNumberInteger:  "bold #3939EA",
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
