package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	dawnBase    = "#faf4ed"
	dawnOverlay = "#f2e9e1"
	dawnMuted   = "#9893a5"
	dawnSubtle  = "#797593"
	dawnText    = "#575279"
	dawnLove    = "#b4637a"
	dawnGold    = "#ea9d34"
	dawnRose    = "#d7827e"
	dawnPine    = "#286983"
	dawnFoam    = "#56949f"
	dawnIris    = "#907aa9"
)

// RosePine (dawn) style.
var RosePineDawn = Register(chroma.MustNewStyle("rose-pine-dawn", chroma.StyleEntries{
	chroma.Text:                dawnText,
	chroma.Error:               dawnLove,
	chroma.Comment:             dawnMuted,
	chroma.Keyword:             dawnPine,
	chroma.KeywordNamespace:    dawnIris,
	chroma.Operator:            dawnSubtle,
	chroma.Punctuation:         dawnSubtle,
	chroma.Name:                dawnRose,
	chroma.NameAttribute:       dawnRose,
	chroma.NameClass:           dawnFoam,
	chroma.NameConstant:        dawnGold,
	chroma.NameDecorator:       dawnSubtle,
	chroma.NameException:       dawnPine,
	chroma.NameFunction:        dawnRose,
	chroma.NameOther:           dawnText,
	chroma.NameTag:             dawnRose,
	chroma.LiteralNumber:       dawnGold,
	chroma.Literal:             dawnGold,
	chroma.LiteralDate:         dawnGold,
	chroma.LiteralString:       dawnGold,
	chroma.LiteralStringEscape: dawnPine,
	chroma.GenericDeleted:      dawnLove,
	chroma.GenericEmph:         "italic",
	chroma.GenericInserted:     dawnFoam,
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   dawnOverlay,
	chroma.Background:          "bg:" + dawnBase,
}))
