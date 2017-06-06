package styles

import (
	"github.com/alecthomas/chroma"
)

// SwapOff theme.
var SwapOff = Register(New("swapoff", map[chroma.TokenType]string{
	chroma.Number:            "bold #ansiyellow",
	chroma.Comment:           "#ansiteal",
	chroma.CommentPreproc:    "bold #ansigreen",
	chroma.String:            "bold #ansiteal",
	chroma.Keyword:           "bold #ansiwhite",
	chroma.GenericHeading:    "bold",
	chroma.GenericSubheading: "bold",
	chroma.GenericStrong:     "bold",
	chroma.GenericUnderline:  "underline",
	chroma.NameTag:           "bold",
	chroma.NameAttribute:     "#ansiteal",
}))
