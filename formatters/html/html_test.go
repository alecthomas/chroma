package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressStyle(t *testing.T) {
	style := "color: #888888; background-color: #ffffff"
	actual := compressStyle(style)
	expected := "color:#888;background-color:#fff"
	require.Equal(t, expected, actual)
}
