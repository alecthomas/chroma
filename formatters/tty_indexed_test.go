package formatters

import (
	"errors"
	"io"
	"iter"
	"slices"
	"strings"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/chroma/v3"
)

func TestTTYTableEntriesAreDistinct(t *testing.T) {
	// The 256-colour palette has only 247 distinct RGB values; see the
	// comment on ttyTables. A count below these values means an entry has
	// been dropped or shadowed by a duplicate key.
	tests := []struct {
		colours  int
		expected int
	}{
		{8, 16},
		{16, 16},
		{256, 247},
	}
	for _, test := range tests {
		table := ttyTables[test.colours]
		assert.Equal(t, test.expected, len(table.foreground))
		assert.Equal(t, test.expected, len(table.background))
	}
}

func TestClosestColour(t *testing.T) {
	actual := findClosest(ttyTables[256], chroma.MustParseColour("#e06c75"))
	assert.Equal(t, chroma.MustParseColour("#d75f87"), actual)
}

func TestClosestColourTieBreak(t *testing.T) {
	target := chroma.MustParseColour("#2ac3de")
	low := chroma.MustParseColour("#00afd7")
	high := chroma.MustParseColour("#00d7d7")
	assert.Equal(t, low.Distance(target), high.Distance(target))

	tests := []struct {
		name    string
		colours []chroma.Colour
	}{
		{
			name:    "lower first",
			colours: []chroma.Colour{low, high},
		},
		{
			name:    "higher first",
			colours: []chroma.Colour{high, low},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findClosestColour(slices.Values(test.colours), target)
			assert.Equal(t, low, actual)
		})
	}
	assert.Equal(t, low, findClosest(ttyTables[256], target))
}

func TestNoneColour(t *testing.T) {
	formatter := TTY256
	tokenType := chroma.None

	style, err := chroma.NewStyle("test", chroma.StyleEntries{
		chroma.Background: "#D0ab1e",
	})
	assert.NoError(t, err)

	stringBuilder := strings.Builder{}
	err = formatter.Format(&stringBuilder, style, chroma.Literator(chroma.Token{
		Type:  tokenType,
		Value: "WORD",
	}))
	assert.NoError(t, err)

	// "178" = #d7af00 approximates #d0ab1e
	//
	// 178 color ref: https://jonasjacek.github.io/colors/
	assert.Equal(t, "\033[38;5;178mWORD\033[0m", stringBuilder.String())
}

func TestFormattersRecoverIteratorPanic(t *testing.T) {
	style, err := chroma.NewStyle("test", chroma.StyleEntries{})
	assert.NoError(t, err)

	for _, name := range Names() {
		t.Run(name, func(t *testing.T) {
			err := Get(name).Format(io.Discard, style, iter.Seq[chroma.Token](func(func(chroma.Token) bool) {
				panic("lexer failed")
			}))
			assert.Error(t, err)
			assert.Equal(t, "lexer failed", err.Error())

			sentinel := errors.New("match timeout")
			err = Get(name).Format(io.Discard, style, iter.Seq[chroma.Token](func(func(chroma.Token) bool) {
				panic(sentinel)
			}))
			assert.Equal(t, sentinel, err)
		})
	}
}
