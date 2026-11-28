package chroma

import (
	"cmp"
	"fmt"
	"iter"
	"path/filepath"
	"strings"
)

// Config for a lexer.
type Config struct {
	// Name of the lexer.
	Name string `xml:"name,omitempty"`

	// Shortcuts for the lexer
	Aliases []string `xml:"alias,omitempty"`

	// File name globs
	Filenames []string `xml:"filename,omitempty"`

	// Secondary file name globs
	AliasFilenames []string `xml:"alias_filename,omitempty"`

	// MIME types
	MimeTypes []string `xml:"mime_type,omitempty"`

	// Regex matching is case-insensitive.
	CaseInsensitive bool `xml:"case_insensitive,omitempty"`

	// Regex matches all characters.
	DotAll bool `xml:"dot_all,omitempty"`

	// Regex does not match across lines ($ matches EOL).
	//
	// Defaults to multiline.
	NotMultiline bool `xml:"not_multiline,omitempty"`

	// Make sure that the input ends with a newline. This
	// is required for some lexers that consume input linewise.
	EnsureNL bool `xml:"ensure_nl,omitempty"`

	// Priority of lexer.
	//
	// If this is 0 it will be treated as a default of 1.
	Priority float32 `xml:"priority,omitempty"`

	// Delegate names the lexer that tokenises the text this lexer emits as
	// Other. See DelegatingLexer.
	Delegate string `xml:"delegate,omitempty"`

	// Analyse is a list of regexes to match against the input.
	//
	// If a match is found, the score is returned if single attribute is set to true,
	// otherwise the sum of all the score of matching patterns will be
	// used as the final score.
	Analyse *AnalyseConfig `xml:"analyse,omitempty"`
}

// validateGlobs returns an error if any of the config's filename globs are
// invalid.
func (c *Config) validateGlobs() error {
	for _, glob := range append(c.Filenames, c.AliasFilenames...) {
		if _, err := filepath.Match(glob, ""); err != nil {
			return fmt.Errorf("%s: %q is not a valid glob: %w", c.Name, glob, err)
		}
	}
	return nil
}

// AnalyseConfig defines the list of regexes analysers.
type AnalyseConfig struct {
	Regexes []RegexConfig `xml:"regex,omitempty"`
	// If true, the first matching score is returned.
	First bool `xml:"first,attr"`
}

// RegexConfig defines a single regex pattern and its score in case of match.
type RegexConfig struct {
	Pattern string  `xml:"pattern,attr"`
	Score   float32 `xml:"score,attr"`
}

// Token output to formatter.
type Token struct {
	Type  TokenType `json:"type"`
	Value string    `json:"value"`
}

func (t Token) String() string   { return t.Value }
func (t Token) GoString() string { return fmt.Sprintf("Token{%s, %q}", t.Type, t.Value) }

// TokeniseOptions are the resolved options for a call to Lexer.Tokenise.
//
// Lexer implementations obtain them with NewTokeniseOptions.
type TokeniseOptions struct {
	// State to start tokenisation in.
	State string
	// Nested is true when tokenising a fragment on behalf of another lexer.
	Nested bool
	// EnsureLF converts CRLF and CR line endings to LF.
	EnsureLF bool
	// Trace, if set, is called after each attempt to match a rule.
	Trace func(Trace)
}

// A TokeniseOption modifies the default TokeniseOptions.
type TokeniseOption func(*TokeniseOptions)

// NewTokeniseOptions applies options to the defaults, which are the "root"
// state with EnsureLF enabled.
func NewTokeniseOptions(options ...TokeniseOption) TokeniseOptions {
	resolved := TokeniseOptions{State: "root", EnsureLF: true}
	for _, option := range options {
		option(&resolved)
	}
	return resolved
}

// WithState sets the state to start tokenisation in. Defaults to "root".
func WithState(state string) TokeniseOption {
	return func(o *TokeniseOptions) { o.State = state }
}

// WithEnsureLF sets whether CRLF and CR line endings are converted to LF.
// Defaults to true.
func WithEnsureLF(ensure bool) TokeniseOption {
	return func(o *TokeniseOptions) { o.EnsureLF = ensure }
}

// WithTrace calls trace after each attempt to match a rule, including in
// lexers that tokenise embedded fragments.
func WithTrace(trace func(Trace)) TokeniseOption {
	return func(o *TokeniseOptions) { o.Trace = trace }
}

// A Lexer for tokenising source code.
type Lexer interface {
	// Config describing the features of the Lexer.
	Config() *Config
	// Tokenise returns an iterator over tokens in text. The iterator panics if
	// lexing fails mid-stream; Formatters recover this, direct consumers must.
	Tokenise(text string, options ...TokeniseOption) (iter.Seq[Token], error)
	// SetRegistry sets the registry this Lexer is associated with.
	//
	// The registry should be used by the Lexer if it needs to look up other
	// lexers.
	SetRegistry(registry *LexerRegistry) Lexer
	// SetAnalyser sets a function the Lexer should use for scoring how
	// likely a fragment of text is to match this lexer, between 0.0 and 1.0.
	// A value of 1 indicates high confidence.
	//
	// Lexers may ignore this if they implement their own analysers.
	SetAnalyser(analyser func(text string) float32) Lexer
	// AnalyseText scores how likely a fragment of text is to match
	// this lexer, between 0.0 and 1.0. A value of 1 indicates high confidence.
	AnalyseText(text string) float32
}

// Trace records one attempt to match a rule. Length is -1 if no rule matched.
type Trace struct {
	Lexer   string  `json:"lexer"`
	State   string  `json:"state"`
	Rule    int     `json:"rule"`
	Pattern string  `json:"pattern"`
	Pos     int     `json:"pos"`
	Length  int     `json:"length"`
	Elapsed float64 `json:"elapsedMs"` // Elapsed time spent matching for this rule.
}

// Lexers is a slice of lexers sortable by name.
type Lexers []Lexer

func (l Lexers) Len() int      { return len(l) }
func (l Lexers) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l Lexers) Less(i, j int) bool {
	return strings.ToLower(l[i].Config().Name) < strings.ToLower(l[j].Config().Name)
}

// PrioritisedLexers is a slice of lexers sortable by priority.
type PrioritisedLexers []Lexer

func (l PrioritisedLexers) Len() int      { return len(l) }
func (l PrioritisedLexers) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l PrioritisedLexers) Less(i, j int) bool {
	return compareLexersByPriority(l[i], l[j]) < 0
}

// compareLexersByPriority orders lexers by descending priority, treating an
// unset priority as 1.
func compareLexersByPriority(a, b Lexer) int {
	ap := a.Config().Priority
	if ap == 0 {
		ap = 1
	}
	bp := b.Config().Priority
	if bp == 0 {
		bp = 1
	}
	return cmp.Compare(bp, ap)
}
