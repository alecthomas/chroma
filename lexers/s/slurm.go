package s

import (
	"github.com/alecthomas/chroma"
	. "github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/b"
	"github.com/alecthomas/chroma/lexers/internal"
)

// Slurm lexer. Lexer for (ba|k|z|)sh Slurm scripts.
var Slurm = internal.Register(MustNewLexer(
	&Config{
		Name:      "Slurm",
		Aliases:   []string{"slurm", "sbatch"},
		Filenames: []string{"*.sl"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if analyser, ok := b.Bash.(chroma.Analyser); ok {
		return analyser.AnalyseText(text)
	}

	return 0
}))
