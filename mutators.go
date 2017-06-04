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

// Combined creates a new anonymous state from the given states, and pushes that state.
func Combined(states ...string) MutatorFunc {
	return func(s *LexerState) error {
		name := "__combined_" + strings.Join(states, "__")
		if _, ok := s.Rules[name]; !ok {
			combined := []CompiledRule{}
			for _, state := range states {
				rules, ok := s.Rules[state]
				if !ok {
					return fmt.Errorf("invalid combine state %q", state)
				}
				combined = append(combined, rules...)
			}
			s.Rules[name] = combined
		}
		s.Rules[s.State][s.Rule].Mutator = Push(name)
		s.Stack = append(s.Stack, name)
		return nil
	}
}

// Push states onto the stack.
func Push(states ...string) MutatorFunc {
	return func(s *LexerState) error {
		if len(states) == 0 {
			s.Stack = append(s.Stack, s.State)
		} else {
			s.Stack = append(s.Stack, states...)
		}
		return nil
	}
}

// Pop state from the stack when rule matches.
func Pop(n int) MutatorFunc {
	return func(state *LexerState) error {
		state.Stack = state.Stack[:len(state.Stack)-n]
		return nil
	}
}

func Default(mutator Mutator) Rule {
	return Rule{Mutator: mutator}
}
