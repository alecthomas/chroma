package styles

import (
	"github.com/alecthomas/chroma"
)

var (
	// Inspired by Apple's Xcode "Default (Dark)" Theme
	background                         = "#1F1F24"
	plainText                          = "#FFFFFF"
	comments                           = "#6C7986"
	documentationMarkup                = "#6C7986"
	documentationMarkupKeywords        = "#92A1B1"
	marks                              = "#92A1B1"
	strings                            = "#FC6A5D"
	characters                         = "#D0BF69"
	numbers                            = "#D0BF69"
	keywords                           = "#FC5FA3"
	preprocessorStatements             = "#FD8F3F"
	urls                               = "#5482FF"
	attributes                         = "#BF8555"
	typeDeclarations                   = "#5DD8FF"
	otherDeclarations                  = "#41A1C0"
	projectClassNames                  = "#9EF1DD"
	projectFunctionAndMethodNames      = "#67B7A4"
	projectConstants                   = "#67B7A4"
	projectTypeNames                   = "#9EF1DD"
	projectInstanceVariablesAndGlobals = "#67B7A4"
	projectPreprocessorMacro           = "#FD8F3F"
	otherClassNames                    = "#D0A8FF"
	otherFunctionAndMethodNames        = "#A167E6"
	otherConstants                     = "#A167E6"
	otherTypeNames                     = "#D0A8FF"
	otherInstanceVariablesAndGlobals   = "#A167E6"
	otherPreprocessorMacros            = "#FD8F3F"
	heading                            = "#AA0D91"
)

// Xcode dark style
var XcodeDark = Register(chroma.MustNewStyle("xcode-dark", chroma.StyleEntries{
	chroma.Background: plainText + " bg: " + background,

	chroma.Comment:          comments,
	chroma.CommentMultiline: comments,
	chroma.CommentPreproc:   preprocessorStatements,
	chroma.CommentSingle:    comments,
	chroma.CommentSpecial:   comments + " italic",

	chroma.Error: "#960050",

	chroma.Keyword:            keywords,
	chroma.KeywordConstant:    keywords,
	chroma.KeywordDeclaration: keywords,
	chroma.KeywordReserved:    keywords,

	chroma.LiteralNumber:        numbers,
	chroma.LiteralNumberBin:     numbers,
	chroma.LiteralNumberFloat:   numbers,
	chroma.LiteralNumberHex:     numbers,
	chroma.LiteralNumberInteger: numbers,
	chroma.LiteralNumberOct:     numbers,

	chroma.LiteralString:         strings,
	chroma.LiteralStringEscape:   strings,
	chroma.LiteralStringInterpol: plainText,

	chroma.Name:              plainText,
	chroma.NameBuiltin:       otherTypeNames,
	chroma.NameBuiltinPseudo: otherFunctionAndMethodNames,
	chroma.NameClass:         typeDeclarations,
	chroma.NameFunction:      otherDeclarations,
	chroma.NameVariable:      otherDeclarations,

	chroma.Operator: plainText,

	chroma.Punctuation: plainText,

	chroma.Text: plainText,
}))
