//go:build !chroma_no_lexers

package lexers

import (
	"embed"

	"github.com/alecthomas/chroma/v3"
)

//go:embed embedded
var embedded embed.FS

func newLexerRegistry() *chroma.LexerRegistry {
	reg := chroma.NewLexerRegistry()
	for i := range embeddedLexers {
		entry := &embeddedLexers[i]
		lexer, err := chroma.NewXMLLexerFromConfig(&entry.config, embedded, entry.path)
		if err != nil {
			panic(err)
		}
		reg.Register(lexer)
	}
	return reg
}
