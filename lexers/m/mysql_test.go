package m_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/m"
)

func TestMySQL_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/mysql_backtick.sql")
	assert.NoError(t, err)

	analyser, ok := m.MySQL.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.5), analyser.AnalyseText(string(data)))
}
