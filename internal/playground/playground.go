// Package playground implements the highlighting pipeline shared by the
// chromad server and the WASM playground library.
package playground

import (
	"strings"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/lexers"
	"github.com/alecthomas/chroma/v3/styles"
)

// Result of highlighting a playground snippet.
type Result struct {
	HTML string
	// Language actually used, or empty if the fallback lexer was applied.
	Language   string
	Background string
}

// Highlight source as HTML with the named lexer and style, autodetecting the
// language if lexerName is unknown.
func Highlight(source, lexerName, styleName string, classes bool) (*Result, error) {
	language := lexers.Get(lexerName)
	if language == nil {
		language = lexers.Analyse(source)
	}
	if language == nil {
		language = lexers.Fallback
	}

	tokens, err := chroma.Coalesce(language).Tokenise(source)
	if err != nil {
		return nil, err
	}

	style := styles.Get(styleName)
	if style == nil {
		style = styles.Fallback
	}

	options := []html.Option{}
	if classes {
		options = append(options, html.WithClasses(true), html.Standalone(true))
	}
	buf := &strings.Builder{}
	if err := html.New(options...).Format(buf, style, tokens); err != nil {
		return nil, err
	}

	lang := language.Config().Name
	if language == lexers.Fallback {
		lang = ""
	}
	return &Result{
		HTML:       buf.String(),
		Language:   lang,
		Background: html.StyleEntryToCSS(style.Get(chroma.Background)),
	}, nil
}
