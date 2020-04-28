package styles

import (
	"github.com/alecthomas/chroma"
)

// PerldocHighContrast style.
var PerldocHighContrast = Register(chroma.MustNewStyle("perldoc-hc", chroma.StyleEntries{
	chroma.TextWhitespace:       "#4C4C4C",
	chroma.Comment:              "#025C02",
	chroma.CommentPreproc:       "#014F5C",
	chroma.CommentSpecial:       "#8B008B bold",
	chroma.LiteralString:        "#911515",
	chroma.LiteralStringHeredoc: "#04544A italic",
	chroma.LiteralStringRegex:   "#04544A",
	chroma.LiteralStringOther:   "#7A3D02",
	chroma.LiteralNumber:        "#7E0293",
	chroma.OperatorWord:         "#8B008B",
	chroma.Keyword:              "#8B008B bold",
	chroma.KeywordType:          "#005471",
	chroma.NameClass:            "#015B2C bold",
	chroma.NameException:        "#015B2C bold",
	chroma.NameFunction:         "#015B2C",
	chroma.NameNamespace:        "#015B2C underline",
	chroma.NameVariable:         "#005471",
	chroma.NameConstant:         "#005471",
	chroma.NameDecorator:        "#454F51",
	chroma.NameTag:              "#8B008B bold",
	chroma.NameAttribute:        "#395100",
	chroma.NameBuiltin:          "#395100",
	chroma.GenericHeading:       "bold #000080",
	chroma.GenericSubheading:    "bold #800080",
	chroma.GenericDeleted:       "#9A0202",
	chroma.GenericInserted:      "#045904",
	chroma.GenericError:         "#9A0202",
	chroma.GenericEmph:          "italic",
	chroma.GenericStrong:        "bold",
	chroma.GenericPrompt:        "#4C4C4C",
	chroma.GenericOutput:        "#474747",
	chroma.GenericTraceback:     "#9A0202",
	chroma.GenericUnderline:     "underline",
	chroma.Error:                "bg:#e3d2d2 #820101",
	chroma.Background:           " bg:#eeeedd",
}))
