package chroma

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInclude(t *testing.T) {
	include := Include("other")
	actual := CompiledRules{
		"root": {
			{Rule: include},
		},
		"other": {
			{Rule: Rule{
				Pattern: "//.+",
				Type:    Comment,
			}},
			{Rule: Rule{
				Pattern: `"[^"]*"`,
				Type:    String,
			}},
		},
	}
	state := &LexerState{
		State: "root",
		Rules: actual,
	}
	err := include.Mutator.Mutate(state)
	require.NoError(t, err)
	expected := CompiledRules{
		"root": {
			{Rule: Rule{
				Pattern: "//.+",
				Type:    Comment,
			}},
			{Rule: Rule{
				Pattern: `"[^"]*"`,
				Type:    String,
			}},
		},
		"other": {
			{Rule: Rule{
				Pattern: "//.+",
				Type:    Comment,
			}},
			{Rule: Rule{
				Pattern: `"[^"]*"`,
				Type:    String,
			}},
		},
	}
	require.Equal(t, expected, actual)
}

func TestCombine(t *testing.T) {
	l := MustNewLexer(nil, Rules{
		"root":  {{`hello`, String, Combined("world", "bye", "space")}},
		"world": {{`world`, Name, nil}},
		"bye":   {{`bye`, Name, nil}},
		"space": {{`\s+`, Whitespace, nil}},
	})
	it, err := l.Tokenise(nil, "hello world")
	require.NoError(t, err)
	expected := []*Token{{String, `hello`}, {Whitespace, ` `}, {Name, `world`}}
	assert.Equal(t, expected, it.Tokens())
}
