package chroma

import "fmt"

// A Modifier modifies the behaviour of the lexer.
type Modifier interface {
	// Preprocess the lexer rules.
	//
	// "self" and "rule" are the rule name and index this Modifier is associated with.
	Preprocess(rules map[string][]CompiledRule, self string, rule int) error
	// Mutate the lexer state machine as it is processing.
	Mutate(state *LexerState) error
}

// A MutatorFunc is a Modifier that mutates the lexer state machine as it is processing.
type MutatorFunc func(state *LexerState) error

func (m MutatorFunc) Preprocess(rules map[string][]CompiledRule, self string, rule int) error {
	return nil
}

func (m MutatorFunc) Mutate(state *LexerState) error {
	return m(state)
}

// A PreprocessorFunc is a Modifier that pre-processes the lexer rules.
type PreprocessorFunc func(rules map[string][]CompiledRule, self string, rule int) error

func (p PreprocessorFunc) Preprocess(rules map[string][]CompiledRule, self string, rule int) error {
	return p(rules, self, rule)
}

func (p PreprocessorFunc) Mutate(state *LexerState) error {
	return nil
}

// Modifiers applies a set of Modifiers in order.
func Modifiers(modifiers ...Modifier) MutatorFunc {
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
		Modifier: PreprocessorFunc(func(rules map[string][]CompiledRule, self string, rule int) error {
			includedRules, ok := rules[state]
			if !ok {
				return fmt.Errorf("invalid include state %q", state)
			}
			stateRules := rules[self]
			stateRules = append(stateRules[:rule], append(includedRules, stateRules[rule+1:]...)...)
			rules[self] = stateRules
			return nil
		}),
	}
}

// Push states onto the stack.
func Push(states ...string) MutatorFunc {
	return func(s *LexerState) error {
		s.Stack = append(s.Stack, states...)
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
