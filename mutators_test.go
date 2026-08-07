package chroma

import (
	"maps"
	"slices"
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestInclude(t *testing.T) {
	include := Include("other")
	actual := CompiledRules{
		"root": {{Rule: include}},
		"other": {
			{Rule: Rule{Pattern: "//.+", Type: Comment}},
			{Rule: Rule{Pattern: `"[^"]*"`, Type: String}},
		},
	}
	lexer := &RegexLexer{rules: actual}
	err := include.Mutator.(LexerMutator).MutateLexer(lexer.rules, "root", 0)
	assert.NoError(t, err)
	expected := []*CompiledRule{
		{Rule: Rule{
			Pattern: "//.+",
			Type:    Comment,
		}},
		{Rule: Rule{
			Pattern: `"[^"]*"`,
			Type:    String,
		}},
	}
	// Include splices the included state's rule pointers into the target state,
	// so "root" and "other" alias the same rules. Assert per-state to keep this
	// independent of how a whole-map comparison renders shared pointers.
	assert.Equal(t, []string{"other", "root"}, slices.Sorted(maps.Keys(actual)))
	assert.Equal(t, expected, actual["other"])
	assert.Equal(t, expected, actual["root"])
}

func TestCombine(t *testing.T) {
	l := mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root":  {{`hello`, String, Combined("world", "bye", "space")}},
		"world": {{`world`, Name, nil}},
		"bye":   {{`bye`, Name, nil}},
		"space": {{`\s+`, Whitespace, nil}},
	})
	it, err := l.Tokenise(nil, "hello world")
	assert.NoError(t, err)
	expected := []Token{{String, `hello`}, {Whitespace, ` `}, {Name, `world`}}
	assert.Equal(t, expected, slices.Collect(it))
}
