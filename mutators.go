package chroma

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// A Mutator modifies the behaviour of the lexer.
type Mutator interface {
	// Mutate the lexer state machine as it is processing.
	Mutate(state *LexerState) error
}

// ValidatingMutator is a Mutator that can validate itself against the compiled rules.
//
// Validation occurs once, after all LexerMutators have been applied.
type ValidatingMutator interface {
	Mutator
	ValidateMutator(rules CompiledRules) error
}

// A LexerMutator is an additional interface that a Mutator can implement
// to modify the lexer when it is compiled.
type LexerMutator interface {
	// MutateLexer can be implemented to mutate the lexer itself.
	//
	// Rules are the lexer rules, state is the state key for the rule the mutator is associated with.
	MutateLexer(rules CompiledRules, state string, rule int) error
}

// A MutatorFunc is a Mutator that mutates the lexer state machine as it is processing.
type MutatorFunc func(state *LexerState) error

func (m MutatorFunc) Mutate(state *LexerState) error { return m(state) } // nolint

type multiMutator struct {
	Mutators []Mutator `xml:"mutator"`
}

func (m *multiMutator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch token := token.(type) {
		case xml.StartElement:
			mutator, err := unmarshalMutator(d, token)
			if err != nil {
				return err
			}
			m.Mutators = append(m.Mutators, mutator)

		case xml.EndElement:
			return nil
		}
	}
}

func (m *multiMutator) ValidateMutator(rules CompiledRules) error {
	for _, mutator := range m.Mutators {
		if validate, ok := mutator.(ValidatingMutator); ok {
			if err := validate.ValidateMutator(rules); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *multiMutator) Mutate(state *LexerState) error {
	for _, modifier := range m.Mutators {
		if err := modifier.Mutate(state); err != nil {
			return err
		}
	}
	return nil
}

// Mutators applies a set of Mutators in order.
func Mutators(modifiers ...Mutator) Mutator {
	return &multiMutator{modifiers}
}

type includeMutator struct {
	State string `xml:"state,attr"`
}

// Include the given state.
func Include(state string) Rule {
	return Rule{Mutator: &includeMutator{state}}
}

func (i *includeMutator) Mutate(s *LexerState) error {
	return fmt.Errorf("should never reach here Include(%q)", i.State)
}

func (i *includeMutator) MutateLexer(rules CompiledRules, state string, rule int) error {
	includedRules, ok := rules[i.State]
	if !ok {
		return fmt.Errorf("invalid include state %q", i.State)
	}
	rules[state] = append(rules[state][:rule], append(includedRules, rules[state][rule+1:]...)...)
	return nil
}

type combinedMutator struct {
	States []string `xml:"state,attr"`
}

// Combined creates a new anonymous state from the given states, and pushes that state.
func Combined(states ...string) Mutator {
	return &combinedMutator{states}
}

func (c *combinedMutator) Mutate(s *LexerState) error {
	return fmt.Errorf("should never reach here Combined(%v)", c.States)
}

func (c *combinedMutator) MutateLexer(rules CompiledRules, state string, rule int) error {
	name := "__combined_" + strings.Join(c.States, "__")
	if _, ok := rules[name]; !ok {
		combined := []*CompiledRule{}
		for _, state := range c.States {
			rules, ok := rules[state]
			if !ok {
				return fmt.Errorf("invalid combine state %q", state)
			}
			combined = append(combined, rules...)
		}
		rules[name] = combined
	}
	rules[state][rule].Mutator = Push(name)
	return nil
}

type pushMutator struct {
	States []string `xml:"state,attr"`
}

var (
	_ ValidatingMutator = (*pushMutator)(nil)
	_ ValidatingMutator = (*multiMutator)(nil)
)

func (p *pushMutator) ValidateMutator(rules CompiledRules) error {
	for _, state := range p.States {
		if state == "#pop" {
			continue
		}
		if _, ok := rules[state]; !ok {
			return fmt.Errorf("invalid push state %q", state)
		}
	}
	return nil
}

func (p *pushMutator) Mutate(s *LexerState) error {
	if len(p.States) == 0 {
		s.Stack = append(s.Stack, s.State)
	} else {
		for _, state := range p.States {
			if state == "#pop" {
				if len(s.Stack) > 0 {
					s.Stack = s.Stack[:len(s.Stack)-1]
				}
			} else {
				s.Stack = append(s.Stack, state)
			}
		}
	}
	return nil
}

// Push states onto the stack.
func Push(states ...string) Mutator {
	return &pushMutator{states}
}

type popMutator struct {
	Depth int `xml:"depth,attr"`
}

func (p *popMutator) Mutate(state *LexerState) error {
	if len(state.Stack) == 0 {
		return fmt.Errorf("nothing to pop")
	}
	state.Stack = state.Stack[:len(state.Stack)-min(p.Depth, len(state.Stack))]
	return nil
}

// Pop state from the stack when rule matches.
func Pop(n int) Mutator {
	return &popMutator{n}
}

// Default returns a Rule that applies a set of Mutators.
func Default(mutators ...Mutator) Rule {
	return Rule{Mutator: Mutators(mutators...)}
}

// Stringify returns the raw string for a set of tokens.
func Stringify(tokens ...Token) string {
	out := []string{}
	for _, t := range tokens {
		out = append(out, t.Value)
	}
	return strings.Join(out, "")
}
