package chroma

import (
	"fmt"
	"regexp"
	"strings"
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

	// Priority, should multiple lexers match and no content is provided
	Priority int

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

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string   { return fmt.Sprintf("Token{%s, %q}", t.Type, t.Value) }
func (t Token) GoString() string { return t.String() }

type TokeniseOptions struct {
	// State to start tokenisation in. Defaults to "root".
	State string
}

type Lexer interface {
	Config() *Config
	Tokenise(options *TokeniseOptions, text string, out func(Token)) error
}

// Analyser determines if this lexer is appropriate for the given text.
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
	Emit(groups []string, lexer Lexer, out func(Token))
}

// EmitterFunc is a function that is an Emitter.
type EmitterFunc func(groups []string, lexer Lexer, out func(Token))

// Emit tokens for groups.
func (e EmitterFunc) Emit(groups []string, lexer Lexer, out func(Token)) { e(groups, lexer, out) }

// ByGroups emits a token for each matching group in the rule's regex.
func ByGroups(emitters ...Emitter) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer, out func(Token)) {
		for i, group := range groups[1:] {
			emitters[i].Emit([]string{group}, lexer, out)
		}
		return
	})
}

// Using returns an Emitter that uses a given Lexer for parsing and emitting.
func Using(lexer Lexer, options *TokeniseOptions) Emitter {
	return EmitterFunc(func(groups []string, _ Lexer, out func(Token)) {
		if err := lexer.Tokenise(options, groups[0], out); err != nil {
			panic(err)
		}
	})
}

// UsingSelf is like Using, but uses the current Lexer.
func UsingSelf(state string) Emitter {
	return EmitterFunc(func(groups []string, lexer Lexer, out func(Token)) {
		if err := lexer.Tokenise(&TokeniseOptions{State: state}, groups[0], out); err != nil {
			panic(err)
		}
	})
}

// Words creates a regex that matches any of the given literal words.
func Words(words ...string) string {
	for i, word := range words {
		words[i] = regexp.QuoteMeta(word)
	}
	return `\b(?:` + strings.Join(words, `|`) + `)\b`
}

// Rules maps from state to a sequence of Rules.
type Rules map[string][]Rule

// MustNewLexer creates a new Lexer or panics.
func MustNewLexer(config *Config, rules Rules) Lexer {
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
func NewLexer(config *Config, rules Rules) (Lexer, error) {
	if _, ok := rules["root"]; !ok {
		return nil, fmt.Errorf("no \"root\" state")
	}
	compiledRules := map[string][]CompiledRule{}
	for state, rules := range rules {
		for _, rule := range rules {
			crule := CompiledRule{Rule: rule}
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
			re, err := regexp.Compile("^(?" + flags + ")(?:" + rule.Pattern + ")")
			if err != nil {
				return nil, fmt.Errorf("invalid regex %q for state %q: %s", rule.Pattern, state, err)
			}
			crule.Regexp = re
			compiledRules[state] = append(compiledRules[state], crule)
		}
	}
	return &regexLexer{
		config: config,
		rules:  compiledRules,
	}, nil
}

// A CompiledRule is a Rule with a pre-compiled regex.
type CompiledRule struct {
	Rule
	Regexp *regexp.Regexp
}

type CompiledRules map[string][]CompiledRule

type LexerState struct {
	Text  string
	Pos   int
	Rules map[string][]CompiledRule
	Stack []string
	State string
	Rule  int
}

type regexLexer struct {
	config *Config
	rules  map[string][]CompiledRule
}

func (r *regexLexer) Config() *Config {
	return r.config
}

func (r *regexLexer) Tokenise(options *TokeniseOptions, text string, out func(Token)) error {
	if options == nil {
		options = defaultOptions
	}
	state := &LexerState{
		Text:  text,
		Stack: []string{options.State},
		Rules: r.rules,
	}
	for state.Pos < len(text) && len(state.Stack) > 0 {
		state.State = state.Stack[len(state.Stack)-1]
		ruleIndex, rule, index := matchRules(state.Text[state.Pos:], state.Rules[state.State])
		// fmt.Println(text[state.Pos:state.Pos+1], rule, state.Text[state.Pos:state.Pos+1])
		// No match.
		if index == nil {
			out(Token{Error, state.Text[state.Pos : state.Pos+1]})
			state.Pos++
			continue
		}
		state.Rule = ruleIndex

		groups := make([]string, len(index)/2)
		for i := 0; i < len(index); i += 2 {
			start := state.Pos + index[i]
			end := state.Pos + index[i+1]
			if start == -1 || end == -1 {
				continue
			}
			groups[i/2] = text[start:end]
		}
		state.Pos += index[1]
		if rule.Type != nil {
			rule.Type.Emit(groups, r, out)
		}
		if rule.Mutator != nil {
			if err := rule.Mutator.Mutate(state); err != nil {
				return err
			}
		}
	}
	return nil
}

// Tokenise text using lexer, returning tokens as a slice.
func Tokenise(lexer Lexer, options *TokeniseOptions, text string) ([]Token, error) {
	out := []Token{}
	return out, lexer.Tokenise(options, text, func(token Token) { out = append(out, token) })
}

func matchRules(text string, rules []CompiledRule) (int, CompiledRule, []int) {
	for i, rule := range rules {
		if index := rule.Regexp.FindStringSubmatchIndex(text); index != nil {
			return i, rule, index
		}
	}
	return 0, CompiledRule{}, nil
}
