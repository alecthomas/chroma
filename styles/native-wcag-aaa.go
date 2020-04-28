package styles

import (
	"github.com/alecthomas/chroma"
)

// NativeHighContrast style.
var NativeHighContrast = Register(chroma.MustNewStyle("native-hc", chroma.StyleEntries{
	chroma.Background:         "#d0d0d0 bg:#202020",
	chroma.TextWhitespace:     "#ABABAB",
	chroma.Comment:            "italic #ACACAC",
	chroma.CommentPreproc:     "noitalic bold #FF8B8B",
	chroma.CommentSpecial:     "noitalic bold #FF9292 bg:#520000",
	chroma.Keyword:            "bold #73C136",
	chroma.KeywordPseudo:      "nobold",
	chroma.OperatorWord:       "bold #73C136",
	chroma.LiteralString:      "#ed9d13",
	chroma.LiteralStringOther: "#ffa500",
	chroma.LiteralNumber:      "#82B6E8",
	chroma.NameBuiltin:        "#61BAC7",
	chroma.NameVariable:       "#40ffff",
	chroma.NameConstant:       "#40ffff",
	chroma.NameClass:          "underline #85B1FF",
	chroma.NameFunction:       "#85B1FF",
	chroma.NameNamespace:      "underline #85B1FF",
	chroma.NameException:      "#bbbbbb",
	chroma.NameTag:            "bold #73C136",
	chroma.NameAttribute:      "#bbbbbb",
	chroma.NameDecorator:      "#ffa500",
	chroma.GenericHeading:     "bold #ffffff",
	chroma.GenericSubheading:  "underline #ffffff",
	chroma.GenericDeleted:     "#FF8D8D",
	chroma.GenericInserted:    "#82C056",
	chroma.GenericError:       "#FF8D8D",
	chroma.GenericEmph:        "italic",
	chroma.GenericStrong:      "bold",
	chroma.GenericPrompt:      "#aaaaaa",
	chroma.GenericOutput:      "#cccccc",
	chroma.GenericTraceback:   "#FF8D8D",
	chroma.GenericUnderline:   "underline",
	chroma.Error:              "bg:#e3d2d2 #820101",
}))
