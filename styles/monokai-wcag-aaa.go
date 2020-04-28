package styles

import (
	"github.com/alecthomas/chroma"
)

// Monokai-hc style.
var Monokai-hc = Register(chroma.MustNewStyle("monokai-hc", chroma.StyleEntries{
	chroma.Text:                "#f8f8f2",
	chroma.Error:               "#EE72A1 bg:#1e0010",
	chroma.Comment:             "#BCB8A5",
	chroma.Keyword:             "#66d9ef",
	chroma.KeywordNamespace:    "#FF94AB",
	chroma.Operator:            "#FF94AB",
	chroma.Punctuation:         "#f8f8f2",
	chroma.Name:                "#f8f8f2",
	chroma.NameAttribute:       "#a6e22e",
	chroma.NameClass:           "#a6e22e",
	chroma.NameConstant:        "#66d9ef",
	chroma.NameDecorator:       "#a6e22e",
	chroma.NameException:       "#a6e22e",
	chroma.NameFunction:        "#a6e22e",
	chroma.NameOther:           "#a6e22e",
	chroma.NameTag:             "#FF94AB",
	chroma.LiteralNumber:       "#C2A5FF",
	chroma.Literal:             "#C2A5FF",
	chroma.LiteralDate:         "#e6db74",
	chroma.LiteralString:       "#e6db74",
	chroma.LiteralStringEscape: "#C2A5FF",
	chroma.GenericDeleted:      "#FF94AB",
	chroma.GenericEmph:         "italic",
	chroma.GenericInserted:     "#a6e22e",
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   "#BCB8A5",
	chroma.Background:          "bg:#272822",
}))
