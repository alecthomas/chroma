package styles

import (
	"github.com/alecthomas/chroma"
)

// SolarizedLight style.
var SolarizedLight = Register(chroma.MustNewStyle("solarized-light", func() chroma.StyleEntries {
	var (
		base00  = "#657b83"
		base0   = "#839496"
		base1   = "#93a1a1"
		base2   = "#eee8d5"
		base3   = "#fdf6e3"
		yellow  = "#b58900"
		orange  = "#cb4b16"
		red     = "#dc322f"
		magenta = "#d33682"
		blue    = "#268bd2"
		cyan    = "#2aa198"
		green   = "#859900"
	)

	return chroma.StyleEntries{
		chroma.Text:                base00,
		chroma.Keyword:             green,
		chroma.KeywordNamespace:    orange,
		chroma.KeywordDeclaration:  blue,
		chroma.KeywordType:         yellow,
		chroma.Name:                base00,
		chroma.NameBuiltin:         green,
		chroma.NameTag:             "bold " + blue,
		chroma.Literal:             cyan,
		chroma.LiteralStringEscape: orange,
		chroma.LiteralNumber:       magenta,
		chroma.LiteralStringRegex:  orange,
		chroma.OperatorWord:        green,
		chroma.Comment:             "italic " + base1,
		chroma.CommentPreproc:      "noitalic " + orange,
		chroma.CommentPreprocFile:  cyan,
		chroma.CommentSpecial:      "noitalic " + cyan,
		chroma.Generic:             base00,
		chroma.GenericDeleted:      red,
		chroma.GenericEmph:         "italic",
		chroma.GenericError:        red,
		chroma.GenericHeading:      "bold",
		chroma.GenericSubheading:   "bold",
		chroma.GenericInserted:     green,
		chroma.GenericPrompt:       "bold " + blue,
		chroma.GenericStrong:       "bold",
		chroma.GenericUnderline:    "underline",
		chroma.Background:          "bg:" + base3,
		chroma.LineNumbers:         base0 + " bg:" + base2,
		chroma.LineHighlight:       "bg:" + base2,
		chroma.Error:               "bg:" + red,
	}
}()))
