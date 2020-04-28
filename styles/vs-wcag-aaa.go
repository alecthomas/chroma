package styles

import (
	"github.com/alecthomas/chroma"
)

// VisualStudioHighContrast style.
var VisualStudioHighContrast = Register(chroma.MustNewStyle("vs-hc", chroma.StyleEntries{
	chroma.Comment:           "#056705",
	chroma.CommentPreproc:    "#0000ff",
	chroma.Keyword:           "#0000ff",
	chroma.OperatorWord:      "#0000ff",
	chroma.KeywordType:       "#005F75",
	chroma.NameClass:         "#005F75",
	chroma.LiteralString:     "#a31515",
	chroma.GenericHeading:    "bold",
	chroma.GenericSubheading: "bold",
	chroma.GenericEmph:       "italic",
	chroma.GenericStrong:     "bold",
	chroma.GenericPrompt:     "bold",
	chroma.Error:             "border:#FF0000",
	chroma.Background:        " bg:#ffffff",
}))
