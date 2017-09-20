package chroma

//go:generate stringer -type TokenType

// TokenType is the type of token to highlight.
//
// It is also an Emitter, emitting a single token of itself
type TokenType int

// Set of TokenTypes.
//
// Categories of types are grouped in ranges of 1000, while sub-categories are in ranges of 100. For
// example, the literal category is in the range 3000-3999. The sub-category for literal strings is
// in the range 3100-3199.

// Meta token types.
const (
	// Default background style.
	Background TokenType = -1 - iota
	// Line numbers in output.
	LineNumbers
	// Line higlight style.
	LineHighlight
	// Character highlight style.
	Highlight
	// Input that could not be tokenised.
	Error
	// Other is used by the Delegate lexer to indicate which tokens should be handled by the delegate.
	Other
	// No highlighting.
	None
	// Final token.
	EOF
)

// Keywords.
const (
	Keyword TokenType = 1000 + iota
	KeywordConstant
	KeywordDeclaration
	KeywordNamespace
	KeywordPseudo
	KeywordReserved
	KeywordType
)

// Names.
const (
	Name TokenType = 2000 + iota
	NameAttribute
	NameBuiltin
	NameBuiltinPseudo
	NameClass
	NameConstant
	NameDecorator
	NameEntity
	NameException
	NameFunction
	NameFunctionMagic
	NameKeyword
	NameLabel
	NameNamespace
	NameOperator
	NameOther
	NamePseudo
	NameProperty
	NameTag
	NameVariable
	NameVariableAnonymous
	NameVariableClass
	NameVariableGlobal
	NameVariableInstance
	NameVariableMagic
)

// Literals.
const (
	Literal TokenType = 3000 + iota
	LiteralDate
	LiteralOther
)

// Strings.
const (
	LiteralString TokenType = 3100 + iota
	LiteralStringAffix
	LiteralStringAtom
	LiteralStringBacktick
	LiteralStringBoolean
	LiteralStringChar
	LiteralStringDelimiter
	LiteralStringDoc
	LiteralStringDouble
	LiteralStringEscape
	LiteralStringHeredoc
	LiteralStringInterpol
	LiteralStringName
	LiteralStringOther
	LiteralStringRegex
	LiteralStringSingle
	LiteralStringSymbol
)

// Literals.
const (
	LiteralNumber TokenType = 3200 + iota
	LiteralNumberBin
	LiteralNumberFloat
	LiteralNumberHex
	LiteralNumberInteger
	LiteralNumberIntegerLong
	LiteralNumberOct
)

// Operators.
const (
	Operator TokenType = 4000 + iota
	OperatorWord
)

// Punctuation.
const (
	Punctuation TokenType = 5000 + iota
)

// Comments.
const (
	Comment TokenType = 6000 + iota
	CommentHashbang
	CommentMultiline
	CommentSingle
	CommentSpecial
)

// Preprocessor "comments".
const (
	CommentPreproc TokenType = 6100 + iota
	CommentPreprocFile
)

// Generic tokens.
const (
	Generic TokenType = 7000 + iota
	GenericDeleted
	GenericEmph
	GenericError
	GenericHeading
	GenericInserted
	GenericOutput
	GenericPrompt
	GenericStrong
	GenericSubheading
	GenericTraceback
	GenericUnderline
)

// Text.
const (
	Text TokenType = 8000 + iota
	TextWhitespace
	TextSymbol
	TextPunctuation
)

// Aliases.
const (
	Whitespace = TextWhitespace

	Date = LiteralDate

	String          = LiteralString
	StringAffix     = LiteralStringAffix
	StringBacktick  = LiteralStringBacktick
	StringChar      = LiteralStringChar
	StringDelimiter = LiteralStringDelimiter
	StringDoc       = LiteralStringDoc
	StringDouble    = LiteralStringDouble
	StringEscape    = LiteralStringEscape
	StringHeredoc   = LiteralStringHeredoc
	StringInterpol  = LiteralStringInterpol
	StringOther     = LiteralStringOther
	StringRegex     = LiteralStringRegex
	StringSingle    = LiteralStringSingle
	StringSymbol    = LiteralStringSymbol

	Number            = LiteralNumber
	NumberBin         = LiteralNumberBin
	NumberFloat       = LiteralNumberFloat
	NumberHex         = LiteralNumberHex
	NumberInteger     = LiteralNumberInteger
	NumberIntegerLong = LiteralNumberIntegerLong
	NumberOct         = LiteralNumberOct
)

func (t TokenType) Category() TokenType {
	return t / 1000 * 1000
}

func (t TokenType) SubCategory() TokenType {
	return t / 100 * 100
}

func (t TokenType) InCategory(other TokenType) bool {
	return t/1000 == other/1000
}

func (t TokenType) InSubCategory(other TokenType) bool {
	return t/100 == other/100
}

func (t TokenType) Emit(groups []string, lexer Lexer, out func(*Token)) {
	out(&Token{Type: t, Value: groups[0]})
}
