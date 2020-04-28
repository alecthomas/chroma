package styles

import (
	"github.com/alecthomas/chroma"
)

// ParaisoDark-hc style.
var ParaisoDark-hc = Register(chroma.MustNewStyle("paraiso-dark-hc", chroma.StyleEntries{
	chroma.Text:                  "#e7e9db",
	chroma.Error:                 "#FF9792",
	chroma.Comment:               "#B6ADB0",
	chroma.Keyword:               "#C1A0E8",
	chroma.KeywordNamespace:      "#5bc4bf",
	chroma.KeywordType:           "#fec418",
	chroma.Operator:              "#5bc4bf",
	chroma.Punctuation:           "#e7e9db",
	chroma.Name:                  "#e7e9db",
	chroma.NameAttribute:         "#2BBEF7",
	chroma.NameClass:             "#fec418",
	chroma.NameConstant:          "#FF9792",
	chroma.NameDecorator:         "#5bc4bf",
	chroma.NameException:         "#FF9792",
	chroma.NameFunction:          "#2BBEF7",
	chroma.NameNamespace:         "#fec418",
	chroma.NameOther:             "#2BBEF7",
	chroma.NameTag:               "#5bc4bf",
	chroma.NameVariable:          "#FF9792",
	chroma.LiteralNumber:         "#f99b15",
	chroma.Literal:               "#f99b15",
	chroma.LiteralDate:           "#5EC796",
	chroma.LiteralString:         "#5EC796",
	chroma.LiteralStringChar:     "#e7e9db",
	chroma.LiteralStringDoc:      "#B6ADB0",
	chroma.LiteralStringEscape:   "#f99b15",
	chroma.LiteralStringInterpol: "#f99b15",
	chroma.GenericDeleted:        "#FF9792",
	chroma.GenericEmph:           "italic",
	chroma.GenericHeading:        "bold #e7e9db",
	chroma.GenericInserted:       "#5EC796",
	chroma.GenericPrompt:         "bold #B6ADB0",
	chroma.GenericStrong:         "bold",
	chroma.GenericSubheading:     "bold #5bc4bf",
	chroma.Background:            "bg:#2f1e2e",
}))
