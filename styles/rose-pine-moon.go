package styles

import (
	"github.com/alecthomas/chroma/v2"
)

var (
	moonBase    = "#232136"
	moonOverlay = "#393552"
	moonMuted   = "#6e6a86"
	moonSubtle  = "#908caa"
	moonText    = "#e0def4"
	moonLove    = "#eb6f92"
	moonGold    = "#f6c177"
	moonRose    = "#ea9a97"
	moonPine    = "#3e8fb0"
	moonFoam    = "#9ccfd8"
	moonIris    = "#c4a7e7"
)

// RosePine (moon) style.
var RosePineMoon = Register(chroma.MustNewStyle("rose-pine-moon", chroma.StyleEntries{
	chroma.Text:                moonText,
	chroma.Error:               moonLove,
	chroma.Comment:             moonMuted,
	chroma.Keyword:             moonPine,
	chroma.KeywordNamespace:    moonIris,
	chroma.Operator:            moonSubtle,
	chroma.Punctuation:         moonSubtle,
	chroma.Name:                moonRose,
	chroma.NameAttribute:       moonRose,
	chroma.NameClass:           moonFoam,
	chroma.NameConstant:        moonGold,
	chroma.NameDecorator:       moonSubtle,
	chroma.NameException:       moonPine,
	chroma.NameFunction:        moonRose,
	chroma.NameOther:           moonText,
	chroma.NameTag:             moonRose,
	chroma.LiteralNumber:       moonGold,
	chroma.Literal:             moonGold,
	chroma.LiteralDate:         moonGold,
	chroma.LiteralString:       moonGold,
	chroma.LiteralStringEscape: moonPine,
	chroma.GenericDeleted:      moonLove,
	chroma.GenericEmph:         "italic",
	chroma.GenericInserted:     moonFoam,
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   moonOverlay,
	chroma.Background:          "bg:" + moonBase,
}))
