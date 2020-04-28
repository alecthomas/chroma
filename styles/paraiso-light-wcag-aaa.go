package styles

import (
	"github.com/alecthomas/chroma"
)

// ParaisoLight-hc style.
var ParaisoLight-hc = Register(chroma.MustNewStyle("paraiso-light-hc", chroma.StyleEntries{
	chroma.Text:                  "#2f1e2e",
	chroma.Error:                 "#931A00",
	chroma.Comment:               "#4C4546",
	chroma.Keyword:               "#5E3082",
	chroma.KeywordNamespace:      "#005451",
	chroma.KeywordType:           "#5F4700",
	chroma.Operator:              "#005451",
	chroma.Punctuation:           "#2f1e2e",
	chroma.Name:                  "#2f1e2e",
	chroma.NameAttribute:         "#054D67",
	chroma.NameClass:             "#5F4700",
	chroma.NameConstant:          "#931A00",
	chroma.NameDecorator:         "#005451",
	chroma.NameException:         "#931A00",
	chroma.NameFunction:          "#054D67",
	chroma.NameNamespace:         "#5F4700",
	chroma.NameOther:             "#054D67",
	chroma.NameTag:               "#005451",
	chroma.NameVariable:          "#931A00",
	chroma.LiteralNumber:         "#6C4001",
	chroma.Literal:               "#6C4001",
	chroma.LiteralDate:           "#015739",
	chroma.LiteralString:         "#015739",
	chroma.LiteralStringChar:     "#2f1e2e",
	chroma.LiteralStringDoc:      "#4C4546",
	chroma.LiteralStringEscape:   "#6C4001",
	chroma.LiteralStringInterpol: "#6C4001",
	chroma.GenericDeleted:        "#931A00",
	chroma.GenericEmph:           "italic",
	chroma.GenericHeading:        "bold #2f1e2e",
	chroma.GenericInserted:       "#015739",
	chroma.GenericPrompt:         "bold #4C4546",
	chroma.GenericStrong:         "bold",
	chroma.GenericSubheading:     "bold #005451",
	chroma.Background:            "bg:#e7e9db",
}))
