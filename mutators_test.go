package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInclude(t *testing.T) {
	include := Include("other")
	actual := CompiledRules{
		"root": {
			CompiledRule{Rule: include},
		},
		"other": {
			CompiledRule{Rule: Rule{
				Pattern: "//.+",
				Type:    Comment,
			}},
			CompiledRule{Rule: Rule{
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
			CompiledRule{Rule: Rule{
				Pattern: "//.+",
				Type:    Comment,
			}},
			CompiledRule{Rule: Rule{
				Pattern: `"[^"]*"`,
				Type:    String,
			}},
		},
		"other": {
			CompiledRule{Rule: Rule{
				Pattern: "//.+",
				Type:    Comment,
			}},
			CompiledRule{Rule: Rule{
				Pattern: `"[^"]*"`,
				Type:    String,
			}},
		},
	}
	require.Equal(t, expected, actual)
}
