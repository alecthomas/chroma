package c

import (
	"github.com/alecthomas/chroma"
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Cuda lexer.
var Cuda = internal.Register(MustNewLexer(
	&Config{
		Name:      "CUDA",
		Aliases:   []string{"cuda", "cu"},
		Filenames: []string{"*.cu", "*.cuh"},
		MimeTypes: []string{"text/x-cuda"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if analyser, ok := C.(chroma.Analyser); ok {
		return analyser.AnalyseText(text)
	}

	return 0
}))
