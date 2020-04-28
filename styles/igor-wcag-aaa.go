package styles

import (
	"github.com/alecthomas/chroma"
)

// IgorHighContrast style.
var IgorHighContrast = Register(chroma.MustNewStyle("igor-hc", chroma.StyleEntries{
	chroma.Comment:       "italic #A80303",
	chroma.Keyword:       "#0000FF",
	chroma.NameFunction:  "#903600",
	chroma.NameDecorator: "#A60384",
	chroma.NameClass:     "#045F5F",
	chroma.LiteralString: "#036503",
	chroma.Background:    " bg:#ffffff",
}))
