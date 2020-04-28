package styles

import (
	"github.com/alecthomas/chroma"
)

// RainbowDash-hc style.
var RainbowDash-hc = Register(chroma.MustNewStyle("rainbow_dash-hc", chroma.StyleEntries{
	chroma.Comment:             "italic #0452A8",
	chroma.CommentPreproc:      "noitalic",
	chroma.CommentSpecial:      "bold",
	chroma.Error:               "bg:#cc0000 #ffffff",
	chroma.GenericDeleted:      "border:#c5060b bg:#ffcccc",
	chroma.GenericEmph:         "italic",
	chroma.GenericError:        "#A80303",
	chroma.GenericHeading:      "bold #0249B5",
	chroma.GenericInserted:     "border:#00cc00 bg:#ccffcc",
	chroma.GenericOutput:       "#575757",
	chroma.GenericPrompt:       "bold #0249B5",
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   "bold #0249B5",
	chroma.GenericTraceback:    "#B30003",
	chroma.GenericUnderline:    "underline",
	chroma.Keyword:             "bold #0249B5",
	chroma.KeywordPseudo:       "nobold",
	chroma.KeywordType:         "#5918bb",
	chroma.NameAttribute:       "italic #0249B5",
	chroma.NameBuiltin:         "bold #5918bb",
	chroma.NameClass:           "underline",
	chroma.NameConstant:        "#08606F",
	chroma.NameDecorator:       "bold #873F00",
	chroma.NameEntity:          "bold #5918bb",
	chroma.NameException:       "bold #5918bb",
	chroma.NameFunction:        "bold #873F00",
	chroma.NameTag:             "bold #0249B5",
	chroma.LiteralNumber:       "bold #5918bb",
	chroma.Operator:            "#0249B5",
	chroma.OperatorWord:        "bold",
	chroma.LiteralString:       "#035F2D",
	chroma.LiteralStringDoc:    "italic",
	chroma.LiteralStringEscape: "bold #B30003",
	chroma.LiteralStringOther:  "#08606F",
	chroma.LiteralStringSymbol: "bold #B30003",
	chroma.Text:                "#4d4d4d",
	chroma.TextWhitespace:      "#535353",
	chroma.Background:          " bg:#ffffff",
}))
