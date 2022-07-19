package styles

import (
	"fmt"

	"github.com/alecthomas/chroma/v2"
)

var (
	// colors used from https://github.com/primer/primitives
	ghRed2      = "#ffa198"
	ghRed3      = "#ff7b72"
	ghRed9      = "#490202"
	ghOrange3   = "#f0883e"
	ghOrange2   = "#ffa657"
	ghGreen1    = "#7ee787"
	ghGreen2    = "#56d364"
	ghBlue1     = "#a5d6ff"
	ghBlue2     = "#79c0ff"
	ghPurple2   = "#d2a8ff"
	ghGray3     = "#8b949e"
	ghFgSubtle  = "#6e7681"
	ghFgDefault = "#c9d1d9"
	ghBgDefault = "#0d1117"
	ghDangerFg  = "#f85149"
)

// GitHub Dark style.
var GitHubDark = Register(chroma.MustNewStyle("github-dark", chroma.StyleEntries{
	chroma.Comment:             "italic " + ghGray3,
	chroma.CommentPreproc:      "bold " + ghGray3,
	chroma.CommentSpecial:      "bold italic " + ghGray3,
	chroma.Error:               ghDangerFg,
	chroma.GenericDeleted:      fmt.Sprintf("bg:%s %s", ghRed9, ghRed2),
	chroma.Generic:             ghFgDefault,
	chroma.GenericEmph:         "italic " + ghFgDefault,
	chroma.GenericError:        ghRed2,
	chroma.GenericHeading:      ghBlue2,
	chroma.GenericInserted:     fmt.Sprintf("bg:%s %s", ghGreen2, ghGreen1),
	chroma.GenericOutput:       ghGray3,
	chroma.GenericPrompt:       ghGray3,
	chroma.GenericStrong:       "bold",
	chroma.GenericSubheading:   ghBlue2,
	chroma.GenericTraceback:    ghRed2,
	chroma.GenericUnderline:    "underline",
	chroma.Keyword:             "bold " + ghRed3,
	chroma.KeywordConstant:     "bold " + ghBlue2,
	chroma.Literal:             ghBlue1,
	chroma.LiteralDate:         ghBlue2,
	chroma.LiteralStringRegex:  ghBlue2,
	chroma.LiteralStringAffix:  ghBlue2,
	chroma.LiteralStringEscape: ghBlue2,
	chroma.Name:                ghFgDefault,
	chroma.NameProperty:        ghBlue2,
	chroma.NameClass:           ghOrange3,
	chroma.NameConstant:        ghBlue2,
	chroma.NameDecorator:       "bold " + ghPurple2,
	chroma.NameEntity:          ghOrange2,
	chroma.NameException:       "bold " + ghOrange3,
	chroma.NameFunction:        "bold " + ghPurple2,
	chroma.NameLabel:           "bold " + ghBlue2,
	chroma.NameNamespace:       ghRed3,
	chroma.NameTag:             ghGreen1,
	chroma.NameVariable:        ghBlue2,
	chroma.Operator:            "bold " + ghRed3,
	chroma.TextWhitespace:      ghFgSubtle,
	chroma.Background:          fmt.Sprintf("bg:%s %s", ghBgDefault, ghFgDefault),
}))
