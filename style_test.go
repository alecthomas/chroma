package chroma

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyleInherit(t *testing.T) {
	s, err := NewStyle("test", StyleEntries{
		Name:         "bold #f00",
		NameVariable: "#fff",
	})
	require.NoError(t, err)
	require.Equal(t, &StyleEntry{Colour: 0x1000000, Bold: true}, s.Get(NameVariable))
}

func TestColours(t *testing.T) {
	s, err := NewStyle("test", StyleEntries{
		Name: "#f00 bg:#001 border:#ansiblue",
	})
	require.NoError(t, err)
	require.Equal(t, &StyleEntry{Colour: 0xff0001, Background: 0x000012, Border: 0x000100}, s.Get(Name))
}
