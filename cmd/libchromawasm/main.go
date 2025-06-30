//go:build wasm

// Package main is an experimental WASM library intended for TinyGO.
package main

import (
	"strings"
	"syscall/js"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

func main() {
	// Register the highlight function with the JavaScript global object
	js.Global().Set("highlight", js.FuncOf(highlight))

	// Keep the program running
	select {}
}

// Highlight source code using Chroma.
//
// Equivalent to the JS function:
//
//	function highlight(source, lexer, styleName, classes)
//
// If the "lexer" is unknown, this will attempt to autodetect the language type.
func highlight(this js.Value, args []js.Value) any {
	source := args[0].String()
	lexer := args[1].String()
	styleName := args[2].String()
	classes := args[3].Bool()

	language := lexers.Get(lexer)
	if language == nil {
		language = lexers.Analyse(source)
		if language != nil {
			lexer = language.Config().Name
		}
	}
	if language == nil {
		language = lexers.Fallback
	}

	tokens, err := chroma.Coalesce(language).Tokenise(nil, source)
	if err != nil {
		panic(err)
	}

	style := styles.Get(styleName)
	if style == nil {
		style = styles.Fallback
	}

	buf := &strings.Builder{}
	options := []html.Option{}
	if classes {
		options = append(options, html.WithClasses(true), html.Standalone(true))
	}
	formatter := html.New(options...)
	err = formatter.Format(buf, style, tokens)
	if err != nil {
		panic(err)
	}
	lang := language.Config().Name
	if language == lexers.Fallback {
		lang = ""
	}
	return js.ValueOf(map[string]any{
		"html":       buf.String(),
		"language":   lang,
		"background": html.StyleEntryToCSS(style.Get(chroma.Background)),
	})
}
