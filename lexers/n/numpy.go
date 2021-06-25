package n

import (
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/shebang"
)

// NumPy lexer.
var NumPy = internal.Register(MustNewLexer(
	&Config{
		Name:    "NumPy",
		Aliases: []string{"numpy"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	hasPythonShebang, _ := shebang.MatchString(text, `pythonw?(3(\.\d)?)?`)
	containsNumpyImport := strings.Contains(text, "import numpy")
	containsFromNumpyImport := strings.Contains(text, "from numpy import")

	var containsImport bool

	if len(text) > 1000 {
		containsImport = strings.Contains(text[:1000], "import ")
	} else {
		containsImport = strings.Contains(text, "import ")
	}

	if (hasPythonShebang || containsImport) && (containsNumpyImport || containsFromNumpyImport) {
		return 1.0
	}

	return 0
}))
