package styles

import (
	"github.com/alecthomas/chroma"
)

// MonokaiLight-hc style.
var MonokaiLight-hc = Register(chroma.MustNewStyle("monokailight-hc", chroma.StyleEntries{
	chroma.Text:                "#272822",
	chroma.Error:               "#EE72A1 bg:#1e0010",
	chroma.Comment:             "#56523D",
	chroma.Keyword:             "#015768",
	chroma.KeywordNamespace:    "#AA0048",
	chroma.Operator:            "#AA0048",
	chroma.Punctuation:         "#111111",
	chroma.Name:                "#111111",
	chroma.NameAttribute:       "#3A5B03",
	chroma.NameClass:           "#3A5B03",
	chroma.NameConstant:        "#015768",
	chroma.NameDecorator:       "#3A5B03",
	chroma.NameException:       "#3A5B03",
	chroma.NameFunction:        "#3A5B03",
	chroma.NameOther:           "#3A5B03",
	chroma.NameTag:             "#AA0048",
	chroma.LiteralNumber:       "#6A20BF",
	chroma.Literal:             "#6A20BF",
	chroma.LiteralDate:         "#7E4A00",
	chroma.LiteralString:       "#7E4A00",
	chroma.LiteralStringEscape: "#6400E2",
	chroma.GenericEmph:         "italic",
	chroma.GenericStrong:       "bold",
	chroma.Background:          " bg:#fafafa",
}))
