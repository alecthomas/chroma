package chroma

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/dlclark/regexp2"
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

	// Make sure that the input does not end with a newline. This
	// is required for some lexers that consume input linewise.
	// DontEnsureNL bool

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
	// Tokenise text and call out for each generated token.
	//
	// A token of type EOF will be passed to out() to signify the end of the stream.
	Tokenise(options *TokeniseOptions, text string, out func(*Token)) error
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

type Rule struct {
	Pattern string
	Type    Emitter
	Mutator Mutator
}

// An Emitter takes group matches and returns tokens.
type Emitter interface {
	// Emit tokens for the given regex groups.
	Emit(groups []string, lexer Lexer, out func(*Token))
}

// EmitterFunc is a function that is an Emitter.
type EmitterFunc func(groups []string, lexer Lexer, out func(*Token))

// Emit tokens for groups.
func (e EmitterFunc) Emit(groups []string, lexer Lexer, out func(*Token)) { e(groups, lexer, out) }

// ByGroups emits a token for each matching group in the rule's regex.
func ByGroups(emitters ...Emitter) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer, out func(*Token)) {
		// NOTE: If this line panics, there is a mismatch with groups. Uncomment the following line to debug.
		// fmt.Printf("%s %#v\n", emitters, groups[1:])
		for i, group := range groups[1:] {
			emitters[i].Emit([]string{group}, lexer, out)
		}
		return
	})
}

// Using returns an Emitter that uses a given Lexer for parsing and emitting.
func Using(lexer Lexer, options *TokeniseOptions) Emitter {
	return EmitterFunc(func(groups []string, _ Lexer, out func(*Token)) {
		if err := lexer.Tokenise(options, groups[0], out); err != nil {
			panic(err)
		}
	})
}

// UsingSelf is like Using, but uses the current Lexer.
func UsingSelf(state string) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer, out func(*Token)) {
		if err := lexer.Tokenise(&TokeniseOptions{State: state}, groups[0], out); err != nil {
			panic(err)
		}
	})
}

// Words creates a regex that matches any of the given literal words.
func Words(prefix, suffix string, words ...string) string {
	for i, word := range words {
		words[i] = regexp.QuoteMeta(word)
	}
	return prefix + `(` + strings.Join(words, `|`) + `)` + suffix
}

// Rules maps from state to a sequence of Rules.
type Rules map[string][]Rule

// MustNewLexer creates a new Lexer or panics.
func MustNewLexer(config *Config, rules Rules) *RegexLexer {
	lexer, err := NewLexer(config, rules)
	if err != nil {
		panic(err)
	}
	return lexer
}

// NewLexer creates a new regex-based Lexer.
//
// "rules" is a state machine transitition map. Each key is a state. Values are sets of rules
// that match input, optionally modify lexer state, and output tokens.
func NewLexer(config *Config, rules Rules) (*RegexLexer, error) {
	if config == nil {
		config = &Config{}
	}
	if _, ok := rules["root"]; !ok {
		return nil, fmt.Errorf("no \"root\" state")
	}
	compiledRules := map[string][]CompiledRule{}
	for state, rules := range rules {
		for _, rule := range rules {
			flags := ""
			if !config.NotMultiline {
				flags += "m"
			}
			if config.CaseInsensitive {
				flags += "i"
			}
			if config.DotAll {
				flags += "s"
			}
			compiledRules[state] = append(compiledRules[state], CompiledRule{Rule: rule, flags: flags})
		}
	}
	return &RegexLexer{
		config: config,
		rules:  compiledRules,
	}, nil
}

// A CompiledRule is a Rule with a pre-compiled regex.
//
// Note that regular expressions are lazily compiled on first use of the lexer.
type CompiledRule struct {
	Rule
	Regexp *regexp2.Regexp
	flags  string
}

type CompiledRules map[string][]CompiledRule

type LexerState struct {
	Text  []rune
	Pos   int
	Rules map[string][]CompiledRule
	Stack []string
	State string
	Rule  int
	// Group matches.
	Groups []string
	// Custum context for mutators.
	MutatorContext map[interface{}]interface{}
}

func (l *LexerState) Set(key interface{}, value interface{}) {
	l.MutatorContext[key] = value
}

func (l *LexerState) Get(key interface{}) interface{} {
	return l.MutatorContext[key]
}

type RegexLexer struct {
	config   *Config
	analyser func(text string) float32

	mu       sync.Mutex
	compiled bool
	rules    map[string][]CompiledRule
}

// SetAnalyser sets the analyser function used to perform content inspection.
func (r *RegexLexer) SetAnalyser(analyser func(text string) float32) *RegexLexer {
	r.analyser = analyser
	return r
}

func (r *RegexLexer) AnalyseText(text string) float32 {
	if r.analyser != nil {
		return r.analyser(text)
	}
	return 0.0
}

func (r *RegexLexer) Config() *Config {
	return r.config
}

// Regex compilation is deferred until the lexer is used. This is to avoid significant init() time costs.
func (r *RegexLexer) maybeCompile() (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.compiled {
		return nil
	}
	for state, rules := range r.rules {
		for i, rule := range rules {
			if rule.Regexp == nil {
				rule.Regexp, err = regexp2.Compile("^(?"+rule.flags+")(?:"+rule.Pattern+")", 0)
				if err != nil {
					return fmt.Errorf("failed to compile rule %s.%d: %s", state, i, err)
				}
			}
			rules[i] = rule
		}
	}
	r.compiled = true
	return nil
}

func (r *RegexLexer) Tokenise(options *TokeniseOptions, text string, out func(*Token)) error {
	if err := r.maybeCompile(); err != nil {
		return err
	}
	if options == nil {
		options = defaultOptions
	}
	state := &LexerState{
		Text:           []rune(text),
		Stack:          []string{options.State},
		Rules:          r.rules,
		MutatorContext: map[interface{}]interface{}{},
	}
	for state.Pos < len(state.Text) && len(state.Stack) > 0 {
		state.State = state.Stack[len(state.Stack)-1]
		ruleIndex, rule, groups := matchRules(state.Text[state.Pos:], state.Rules[state.State])
		// No match.
		if groups == nil {
			out(&Token{Error, string(state.Text[state.Pos : state.Pos+1])})
			state.Pos++
			continue
		}
		state.Rule = ruleIndex

		state.Groups = groups
		state.Pos += len(groups[0])
		if rule.Mutator != nil {
			if err := rule.Mutator.Mutate(state); err != nil {
				return err
			}
		}
		if rule.Type != nil {
			rule.Type.Emit(state.Groups, r, out)
		}
	}
	out(&Token{Type: EOF})
	return nil
}

// Tokenise text using lexer, returning tokens as a slice.
func Tokenise(lexer Lexer, options *TokeniseOptions, text string) ([]*Token, error) {
	out := []*Token{}
	return out, lexer.Tokenise(options, text, func(token *Token) { out = append(out, token) })
}

func matchRules(text []rune, rules []CompiledRule) (int, CompiledRule, []string) {
	for i, rule := range rules {
		match, err := rule.Regexp.FindRunesMatch(text)
		if match != nil && err == nil {
			groups := []string{}
			for _, g := range match.Groups() {
				groups = append(groups, g.String())
			}
			return i, rule, groups
		}
	}
	return 0, CompiledRule{}, nil
}
