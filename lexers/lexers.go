package lexers

import (
	"embed"

	"github.com/alecthomas/chroma/v3"
)

//go:generate go run github.com/alecthomas/chroma/v3/internal/metadatagen -dir embedded -out lexers_gen.go

//go:embed embedded
var embedded embed.FS

// GlobalLexerRegistry is the global LexerRegistry of Lexers.
//
// It is populated from generated metadata mirroring the embedded XML
// definitions, so that registering every lexer does not require parsing any
// XML; each lexer's rules are still loaded from its XML on first use.
var GlobalLexerRegistry = func() *chroma.LexerRegistry {
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
}()

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string {
	return GlobalLexerRegistry.Names(withAliases)
}

// Aliases of all the lexers, and skip those lexers who do not have any aliases,
// or show their name instead
func Aliases(skipWithoutAliases bool) []string {
	return GlobalLexerRegistry.Aliases(skipWithoutAliases)
}

// Get a Lexer by name, alias or file extension.
//
// Note that this if there isn't an exact match on name or alias, this will
// call Match(), so it is not efficient.
func Get(name string) chroma.Lexer {
	return GlobalLexerRegistry.Get(name)
}

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer {
	return GlobalLexerRegistry.MatchMimeType(mimeType)
}

// Match returns the first lexer matching filename.
//
// Note that this iterates over all file patterns in all lexers, so it's not
// particularly efficient.
func Match(filename string) chroma.Lexer {
	return GlobalLexerRegistry.Match(filename)
}

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer {
	return GlobalLexerRegistry.Register(lexer)
}

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer {
	return GlobalLexerRegistry.Analyse(text)
}

// PlaintextRules is used for the fallback lexer as well as the explicit
// plaintext lexer.
func PlaintextRules() chroma.Rules {
	return chroma.Rules{
		"root": []chroma.Rule{
			{`.+`, chroma.Text, nil},
			{`\n`, chroma.Text, nil},
		},
	}
}

// Fallback lexer if no other is found.
var Fallback chroma.Lexer = chroma.MustNewLexer(&chroma.Config{
	Name:      "fallback",
	Filenames: []string{"*"},
	Priority:  -1,
}, PlaintextRules)
