package styles

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alecthomas/chroma"
)

func TestStyleAddNoInherit(t *testing.T) {
	s := New("test", Entries{
		chroma.Name:         "bold #f00",
		chroma.NameVariable: "noinherit #fff",
	})
	require.Equal(t, &Entry{Colour: 0x1000000}, s.Get(chroma.NameVariable))
}

func TestStyleInherit(t *testing.T) {
	s := New("test", Entries{
		chroma.Name:         "bold #f00",
		chroma.NameVariable: "#fff",
	})
	require.Equal(t, &Entry{Colour: 0x1000000, Bold: true}, s.Get(chroma.NameVariable))
}

func TestColours(t *testing.T) {
	s := New("test", Entries{
		chroma.Name: "#f00 bg:#001 border:#ansiblue",
	})
	require.Equal(t, &Entry{Colour: 0xff0001, Background: 0x000012, Border: 0x000100}, s.Get(chroma.Name))
}
