package chroma

import (
	"fmt"
	"iter"
	"strings"
)

type delegatingLexer struct {
	root     Lexer
	rootName string
	registry *LexerRegistry
	language Lexer
}

// DelegatingLexer combines two lexers to handle the common case of a language embedded inside another, such as PHP
// inside HTML or PHP inside plain text.
//
// It takes two lexers as arguments: a root lexer and a language lexer. First everything is scanned using the language
// lexer, which must return "Other" for unrecognised tokens. Then all "Other" tokens are lexed using the root lexer.
// Finally, these two sets of tokens are merged.
//
// The lexers from the template lexer package use this base lexer.
func DelegatingLexer(root Lexer, language Lexer) Lexer {
	return &delegatingLexer{
		root:     root,
		language: language,
	}
}

// delegateByName is DelegatingLexer with the root looked up by name in the
// registry when first needed, as XML definitions can only refer to it by name.
func delegateByName(rootName string, language Lexer) Lexer {
	return &delegatingLexer{rootName: rootName, language: language}
}

func (d *delegatingLexer) rootLexer() (Lexer, error) {
	if d.root != nil {
		return d.root, nil
	}
	if d.registry == nil {
		return nil, fmt.Errorf("no LexerRegistry available to delegate to %q", d.rootName)
	}
	root := d.registry.Get(d.rootName)
	if root == nil {
		return nil, fmt.Errorf("no such lexer %q", d.rootName)
	}
	return root, nil
}

func (d *delegatingLexer) SetTracing(enable bool) {
	if l, ok := d.language.(TracingLexer); ok {
		l.SetTracing(enable)
	}
	root, err := d.rootLexer()
	if err != nil {
		return
	}
	if l, ok := root.(TracingLexer); ok {
		l.SetTracing(enable)
	}
}

func (d *delegatingLexer) AnalyseText(text string) float32 {
	return d.language.AnalyseText(text)
}

func (d *delegatingLexer) SetAnalyser(analyser func(text string) float32) Lexer {
	d.language.SetAnalyser(analyser)
	return d
}

func (d *delegatingLexer) SetRegistry(r *LexerRegistry) Lexer {
	d.registry = r
	if d.root != nil {
		d.root.SetRegistry(r)
	}
	d.language.SetRegistry(r)
	return d
}

func (d *delegatingLexer) Config() *Config {
	return d.language.Config()
}

func (d *delegatingLexer) Tokenise(options *TokeniseOptions, text string) (tokens iter.Seq[Token], err error) { // nolint: gocognit
	// Recover panics from the eager language pass before any Formatter is on the stack.
	defer func() {
		if perr := recover(); perr != nil {
			if e, ok := perr.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", perr)
			}
		}
	}()
	rootLexer, err := d.rootLexer()
	if err != nil {
		return nil, err
	}
	languageTokens, err := Tokenise(Coalesce(d.language), options, text)
	if err != nil {
		return nil, err
	}
	var rootText strings.Builder
	hasLanguageTokens := false
	for token := range languageTokens {
		if token.Type == Other {
			rootText.WriteString(token.Value)
		} else {
			hasLanguageTokens = true
		}
	}
	if !hasLanguageTokens {
		return rootLexer.Tokenise(options, text)
	}

	rootTokens, err := Tokenise(Coalesce(rootLexer), options, rootText.String())
	if err != nil {
		return nil, err
	}
	languageTokens, err = Tokenise(Coalesce(d.language), options, text)
	if err != nil {
		return nil, err
	}
	return func(yield func(Token) bool) {
		nextRoot, stopRoot := iter.Pull(rootTokens)
		defer stopRoot()

		var root Token
		var pendingRoot Token
		for language := range languageTokens {
			if language.Value == "" {
				continue
			}
			if language.Type != Other {
				if pendingRoot.Value != "" {
					if !yield(pendingRoot) {
						return
					}
					pendingRoot = Token{}
				}
				if !yield(language) {
					return
				}
				continue
			}

			// Consume only Other spans because embedded language spans were removed from the root input.
			remaining := len(language.Value)
			for remaining > 0 {
				if root.Value == "" {
					var ok bool
					root, ok = nextRoot()
					if !ok {
						break
					}
					continue
				}
				consumed, rest := splitToken(root, min(remaining, len(root.Value)))
				root = rest
				remaining -= len(consumed.Value)
				if pendingRoot.Value == "" {
					pendingRoot = consumed
				} else {
					pendingRoot.Value += consumed.Value
				}
				if root.Value == "" {
					if !yield(pendingRoot) {
						return
					}
					pendingRoot = Token{}
				}
			}
		}
		if pendingRoot.Value != "" {
			yield(pendingRoot)
		}
	}, nil
}

func splitToken(t Token, offset int) (l Token, r Token) {
	if offset <= 0 {
		return Token{}, t
	}
	if offset >= len(t.Value) {
		return t, Token{}
	}
	l = t.Clone()
	r = t.Clone()
	l.Value = l.Value[:offset]
	r.Value = r.Value[offset:]
	return
}
