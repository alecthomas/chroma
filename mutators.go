package chroma

import (
	"fmt"
	"strings"
)

// A Mutator modifies the behaviour of the lexer.
type Mutator interface {
	// Mutate the lexer state machine as it is processing.
	Mutate(state *LexerState) error
}

// A LexerMutator is an additional interface that a Mutator can implement
// to modify the lexer when it is compiled.
type LexerMutator interface {
	MutateLexer(lexer *RegexLexer, rule *CompiledRule) error
}

// A MutatorFunc is a Mutator that mutates the lexer state machine as it is processing.
type MutatorFunc func(state *LexerState) error

func (m MutatorFunc) Mutate(state *LexerState) error { return m(state) }

// Mutators applies a set of Mutators in order.
func Mutators(modifiers ...Mutator) MutatorFunc {
	return func(state *LexerState) error {
		for _, modifier := range modifiers {
			if err := modifier.Mutate(state); err != nil {
				return err
			}
		}
		return nil
	}
}

// Include the given state.
func Include(state string) Rule {
	return Rule{
		Mutator: MutatorFunc(func(ls *LexerState) error {
			includedRules, ok := ls.Rules[state]
			if !ok {
				return fmt.Errorf("invalid include state %q", state)
			}
			stateRules := ls.Rules[ls.State]
			stateRules = append(stateRules[:ls.Rule], append(includedRules, stateRules[ls.Rule+1:]...)...)
			ls.Rules[ls.State] = stateRules
			return nil
		}),
	}
}

type combinedMutator struct {
	states []string
}

func (c *combinedMutator) Mutate(s *LexerState) error { return nil }

func (c *combinedMutator) MutateLexer(lexer *RegexLexer, rule *CompiledRule) error {
	name := "__combined_" + strings.Join(c.states, "__")
	if _, ok := lexer.rules[name]; !ok {
		combined := []*CompiledRule{}
		for _, state := range c.states {
			rules, ok := lexer.rules[state]
			if !ok {
				return fmt.Errorf("invalid combine state %q", state)
			}
			combined = append(combined, rules...)
		}
		lexer.rules[name] = combined
	}
	rule.Mutator = Push(name)
	return nil
}

// Combined creates a new anonymous state from the given states, and pushes that state.
func Combined(states ...string) Mutator {
	return &combinedMutator{states}
}

// Push states onto the stack.
func Push(states ...string) MutatorFunc {
	return func(s *LexerState) error {
		if len(states) == 0 {
			s.Stack = append(s.Stack, s.State)
		} else {
			for _, state := range states {
				if state == "#pop" {
					s.Stack = s.Stack[:len(s.Stack)-1]
				} else {
					s.Stack = append(s.Stack, state)
				}
			}
		}
		return nil
	}
}

// Pop state from the stack when rule matches.
func Pop(n int) MutatorFunc {
	return func(state *LexerState) error {
		if len(state.Stack) == 0 {
			return fmt.Errorf("nothing to pop")
		}
		state.Stack = state.Stack[:len(state.Stack)-n]
		return nil
	}
}

func Default(mutators ...Mutator) Rule {
	return Rule{Mutator: Mutators(mutators...)}
}
