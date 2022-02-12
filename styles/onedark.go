package styles

import (
	"github.com/alecthomas/chroma/v2"
)

// OneDark style.
var OneDark = Register(chroma.MustNewStyle("one-dark", chroma.StyleEntries{
	chroma.Background:         "#ABB2BF bg:#282C34",
	chroma.Punctuation:        "#ABB2BF",
	chroma.PunctuationMarker:  "#ABB2BF",
	chroma.Keyword:            "#C678DD",
	chroma.KeywordConstant:    "#E5C07B",
	chroma.KeywordDeclaration: "#C678DD",
	chroma.KeywordNamespace:   "#C678DD",
	chroma.KeywordReserved:    "#C678DD",
	chroma.KeywordType:        "#E5C07B",
	chroma.Name:               "#E06C75",
	chroma.NameAttribute:      "#E06C75",
	chroma.NameBuiltin:        "#E5C07B",
	chroma.NameClass:          "#E5C07B",
	chroma.NameFunction:       "bold #61AFEF",
	chroma.NameFunctionMagic:  "bold #56B6C2",
	chroma.NameOther:          "#E06C75",
	chroma.NameTag:            "#E06C75",
	chroma.NameDecorator:      "#61AFEF",
	chroma.LiteralString:      "#98C379",
	chroma.LiteralNumber:      "#D19A66",
	chroma.Operator:           "#56B6C2",
	chroma.Comment:            "#7F848E",
}))
