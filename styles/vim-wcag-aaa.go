package styles

import (
	"github.com/alecthomas/chroma"
)

// Vim-hc style.
var Vim-hc = Register(chroma.MustNewStyle("vim-hc", chroma.StyleEntries{
	chroma.Background:         "#cccccc bg:#000000",
	chroma.Comment:            "#9191D0",
	chroma.CommentSpecial:     "bold #FF6262",
	chroma.Keyword:            "#cdcd00",
	chroma.KeywordDeclaration: "#00cd00",
	chroma.KeywordNamespace:   "#F054F0",
	chroma.KeywordType:        "#00cd00",
	chroma.Operator:           "#45A3D7",
	chroma.OperatorWord:       "#cdcd00",
	chroma.NameClass:          "#00cdcd",
	chroma.NameBuiltin:        "#F054F0",
	chroma.NameException:      "bold #9797CB",
	chroma.NameVariable:       "#00cdcd",
	chroma.LiteralString:      "#FF6262",
	chroma.LiteralNumber:      "#F054F0",
	chroma.GenericHeading:     "bold #9191D0",
	chroma.GenericSubheading:  "bold #CE79CE",
	chroma.GenericDeleted:     "#FF6262",
	chroma.GenericInserted:    "#00cd00",
	chroma.GenericError:       "#FF6060",
	chroma.GenericEmph:        "italic",
	chroma.GenericStrong:      "bold",
	chroma.GenericPrompt:      "bold #9191D0",
	chroma.GenericOutput:      "#9D9D9D",
	chroma.GenericTraceback:   "#7E8EFF",
	chroma.GenericUnderline:   "underline",
	chroma.Error:              "border:#FF0000",
}))
