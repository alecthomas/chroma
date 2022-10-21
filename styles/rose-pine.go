package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	base    = "#191724"
	overlay = "#26233a"
	muted   = "#6e6a86"
	subtle  = "#908caa"
	text    = "#e0def4"
	love    = "#eb6f92"
	gold    = "#f6c177"
	rose    = "#ebbcba"
	pine    = "#31748f"
	foam    = "#9ccfd8"
	iris    = "#c4a7e7"
)

// RosePine (Main) style.
var RosePine = Register(chroma.MustNewStyle("rose-pine", chroma.StyleEntries{
	chroma.Text:                text,
	chroma.Error:               love,
	chroma.Comment:             muted,
	chroma.Keyword:             pine,
	chroma.KeywordNamespace:    iris,
	chroma.Operator:            subtle,
	chroma.Punctuation:         subtle,
	chroma.Name:                rose,
	chroma.NameAttribute:       rose,
	chroma.NameClass:           foam,
	chroma.NameConstant:        gold,
	chroma.NameDecorator:       subtle,
	chroma.NameException:       pine,
	chroma.NameFunction:        rose,
	chroma.NameOther:           text,
	chroma.NameTag:             rose,
	chroma.LiteralNumber:       gold,
	chroma.Literal:             gold,
	chroma.LiteralDate:         gold,
	chroma.LiteralString:       gold,
	chroma.LiteralStringEscape: pine,
	chroma.GenericDeleted:      love,
	chroma.GenericEmph:         "italic",
	chroma.GenericInserted:     foam,
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   overlay,
	chroma.Background:          "bg:" + base,
}))
