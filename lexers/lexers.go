// Package lexers contains the registry of all lexers.
//
// Sub-packages contain lexer implementations.
package lexers

// nolint: golint
import (
	"github.com/colin3dmax/chroma"
	_ "github.com/colin3dmax/chroma/lexers/a"
	_ "github.com/colin3dmax/chroma/lexers/b"
	_ "github.com/colin3dmax/chroma/lexers/c"
	_ "github.com/colin3dmax/chroma/lexers/circular"
	_ "github.com/colin3dmax/chroma/lexers/d"
	_ "github.com/colin3dmax/chroma/lexers/e"
	_ "github.com/colin3dmax/chroma/lexers/f"
	_ "github.com/colin3dmax/chroma/lexers/g"
	_ "github.com/colin3dmax/chroma/lexers/h"
	_ "github.com/colin3dmax/chroma/lexers/i"
	"github.com/colin3dmax/chroma/lexers/internal"
	_ "github.com/colin3dmax/chroma/lexers/j"
	_ "github.com/colin3dmax/chroma/lexers/k"
	_ "github.com/colin3dmax/chroma/lexers/l"
	_ "github.com/colin3dmax/chroma/lexers/m"
	_ "github.com/colin3dmax/chroma/lexers/n"
	_ "github.com/colin3dmax/chroma/lexers/o"
	_ "github.com/colin3dmax/chroma/lexers/p"
	_ "github.com/colin3dmax/chroma/lexers/q"
	_ "github.com/colin3dmax/chroma/lexers/r"
	_ "github.com/colin3dmax/chroma/lexers/s"
	_ "github.com/colin3dmax/chroma/lexers/t"
	_ "github.com/colin3dmax/chroma/lexers/v"
	_ "github.com/colin3dmax/chroma/lexers/w"
	_ "github.com/colin3dmax/chroma/lexers/x"
	_ "github.com/colin3dmax/chroma/lexers/y"
)

// Registry of Lexers.
var Registry = internal.Registry

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string { return internal.Names(withAliases) }

// Get a Lexer by name, alias or file extension.
func Get(name string) chroma.Lexer { return internal.Get(name) }

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer { return internal.MatchMimeType(mimeType) }

// Match returns the first lexer matching filename.
func Match(filename string) chroma.Lexer { return internal.Match(filename) }

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer { return internal.Analyse(text) }

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer { return internal.Register(lexer) }

// Fallback lexer if no other is found.
var Fallback = internal.Fallback
