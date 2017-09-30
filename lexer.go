package chroma

import (
	"fmt"
)

var (
	defaultOptions = &TokeniseOptions{
		State: "root",
	}
)

// Config for a lexer.
type Config struct {
	// Name of the lexer.
	Name string

	// Shortcuts for the lexer
	Aliases []string

	// File name globs
	Filenames []string

	// Secondary file name globs
	AliasFilenames []string

	// MIME types
	MimeTypes []string

	// Regex matching is case-insensitive.
	CaseInsensitive bool

	// Regex matches all characters.
	DotAll bool

	// Regex does not match across lines ($ matches EOL).
	//
	// Defaults to multiline.
	NotMultiline bool

	// Don't strip leading and trailing newlines from the input.
	// DontStripNL bool

	// Strip all leading and trailing whitespace from the input
	// StripAll bool

	// Make sure that the input ends with a newline. This
	// is required for some lexers that consume input linewise.
	EnsureNL bool

	// If given and greater than 0, expand tabs in the input.
	// TabSize int
}

// Token output to formatter.
type Token struct {
	Type  TokenType
	Value string
}

func (t *Token) String() string   { return t.Value }
func (t *Token) GoString() string { return fmt.Sprintf("Token{%s, %q}", t.Type, t.Value) }

func (t *Token) Clone() *Token {
	clone := &Token{}
	*clone = *t
	return clone
}

type TokeniseOptions struct {
	// State to start tokenisation in. Defaults to "root".
	State string
}

// A Lexer for tokenising source code.
type Lexer interface {
	// Config describing the features of the Lexer.
	Config() *Config
	// Tokenise returns an Iterator over tokens in text.
	Tokenise(options *TokeniseOptions, text string) (Iterator, error)
}

type Lexers []Lexer

// Pick attempts to pick the best Lexer for a piece of source code. May return nil.
func (l Lexers) Pick(text string) Lexer {
	if len(l) == 0 {
		return nil
	}
	var picked Lexer
	highest := float32(-1)
	for _, lexer := range l {
		if analyser, ok := lexer.(Analyser); ok {
			score := analyser.AnalyseText(text)
			if score > highest {
				highest = score
				picked = lexer
				continue
			}
		}
	}
	return picked
}

func (l Lexers) Len() int           { return len(l) }
func (l Lexers) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Lexers) Less(i, j int) bool { return l[i].Config().Name < l[j].Config().Name }

// Analyser determines how appropriate this lexer is for the given text.
type Analyser interface {
	AnalyseText(text string) float32
}
