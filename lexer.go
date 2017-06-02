package chroma

import (
	"fmt"
	"regexp"
	"strings"
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

	// Don't strip leading and trailing newlines from the input.
	DontStripNL bool

	// Strip all leading and trailing whitespace from the input
	StripAll bool

	// Make sure that the input does not end with a newline. This
	// is required for some lexers that consume input linewise.
	DontEnsureNL bool

	// If given and greater than 0, expand tabs in the input.
	TabSize int

	// If given, must be an encoding name. This encoding will be used to
	// convert the input string to Unicode, if it is not already a Unicode
	// string.
	Encoding string
}

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string   { return fmt.Sprintf("Token{%s, %q}", t.Type, t.Value) }
func (t Token) GoString() string { return t.String() }

type Lexer interface {
	Config() *Config
	Tokenise(text string, out func(Token)) error
}

// Analyser determines if this lexer is appropriate for the given text.
type Analyser interface {
	AnalyseText(text string) float32
}

type Rule struct {
	Pattern  string
	Type     Emitter
	Modifier Modifier
}

// An Emitter takes group matches and returns tokens.
type Emitter interface {
	// Emit tokens for the given regex groups.
	Emit(groups []string, out func(Token))
}

// EmitterFunc is a function that is an Emitter.
type EmitterFunc func(groups []string, out func(Token))

// Emit tokens for groups.
func (e EmitterFunc) Emit(groups []string, out func(Token)) { e(groups, out) }

// ByGroups emits a token for each matching group in the rule's regex.
func ByGroups(emitters ...Emitter) Emitter {
	return EmitterFunc(func(groups []string, out func(Token)) {
		for i, group := range groups[1:] {
			emitters[i].Emit([]string{group}, out)
		}
		return
	})
}

// Using uses a given Lexer for parsing and emitting.
func Using(lexer Lexer) Emitter {
	return EmitterFunc(func(groups []string, out func(Token)) {
		if err := lexer.Tokenise(groups[0], out); err != nil {
			// TODO: Emitters should return an error, though it's not clear what one would do with
			// it.
			panic(err)
		}
	})
}

// Words creates a regex that matches any of the given literal words.
func Words(words ...string) string {
	for i, word := range words {
		words[i] = regexp.QuoteMeta(word)
	}
	return "\\b(?:" + strings.Join(words, "|") + ")\\b"
}

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
			re, err := regexp.Compile("^(?m)" + rule.Pattern)
			if err != nil {
				return nil, fmt.Errorf("invalid regex %q for state %q: %s", rule.Pattern, state, err)
			}
			crule.Regexp = re
			compiledRules[state] = append(compiledRules[state], crule)
		}
	}
	// Apply any pre-processor modifiers.
	for state, rules := range compiledRules {
		for index, rule := range rules {
			if rule.Modifier != nil {
				err := rule.Modifier.Preprocess(compiledRules, state, index)
				if err != nil {
					return nil, err
				}
			}
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

type regexLexer struct {
	config *Config
	rules  map[string][]CompiledRule
}

func (r *regexLexer) Config() *Config {
	return r.config
}

type LexerState struct {
	Text  string
	Pos   int
	Stack []string
	Rules map[string][]CompiledRule
	State string
}

func (r *regexLexer) Tokenise(text string, out func(Token)) error {
	state := &LexerState{
		Text:  text,
		Stack: []string{"root"},
		Rules: r.rules,
	}
	for state.Pos < len(text) && len(state.Stack) > 0 {
		state.State = state.Stack[len(state.Stack)-1]
		rule, index := matchRules(state.Text[state.Pos:], state.Rules[state.State])
		// No match.
		if index == nil {
			out(Token{Error, state.Text[state.Pos : state.Pos+1]})
			state.Pos++
			continue
		}

		groups := make([]string, len(index)/2)
		for i := 0; i < len(index); i += 2 {
			groups[i/2] = text[state.Pos+index[i] : state.Pos+index[i+1]]
		}
		state.Pos += index[1]
		if rule.Modifier != nil {
			if err := rule.Modifier.Mutate(state); err != nil {
				return err
			}
		} else {
			rule.Type.Emit(groups, out)
		}
	}
	return nil
}

func matchRules(text string, rules []CompiledRule) (CompiledRule, []int) {
	for _, rule := range rules {
		if index := rule.Regexp.FindStringSubmatchIndex(text); index != nil {
			return rule, index
		}
	}
	return CompiledRule{}, nil
}
