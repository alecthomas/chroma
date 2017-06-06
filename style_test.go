package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyleAddNoInherit(t *testing.T) {
	s := NewStyle("test", StyleEntries{
		Name:         "bold #f00",
		NameVariable: "noinherit #fff",
	})
	require.Equal(t, &StyleEntry{Colour: 0x1000000}, s.Get(NameVariable))
}

func TestStyleInherit(t *testing.T) {
	s := NewStyle("test", StyleEntries{
		Name:         "bold #f00",
		NameVariable: "#fff",
	})
	require.Equal(t, &StyleEntry{Colour: 0x1000000, Bold: true}, s.Get(NameVariable))
}

func TestColours(t *testing.T) {
	s := NewStyle("test", StyleEntries{
		Name: "#f00 bg:#001 border:#ansiblue",
	})
	require.Equal(t, &StyleEntry{Colour: 0xff0001, Background: 0x000012, Border: 0x000100}, s.Get(Name))
}
