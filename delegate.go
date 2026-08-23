package chroma

import (
	"iter"
	"strings"
)

type delegatingLexer struct {
	root     Lexer
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

func (d *delegatingLexer) SetTracing(enable bool) {
	if l, ok := d.language.(TracingLexer); ok {
		l.SetTracing(enable)
	}
	if l, ok := d.root.(TracingLexer); ok {
		l.SetTracing(enable)
	}
}

func (d *delegatingLexer) AnalyseText(text string) float32 {
	return d.root.AnalyseText(text)
}

func (d *delegatingLexer) SetAnalyser(analyser func(text string) float32) Lexer {
	d.root.SetAnalyser(analyser)
	return d
}

func (d *delegatingLexer) SetRegistry(r *LexerRegistry) Lexer {
	d.root.SetRegistry(r)
	d.language.SetRegistry(r)
	return d
}

func (d *delegatingLexer) Config() *Config {
	return d.language.Config()
}

func (d *delegatingLexer) Tokenise(options *TokeniseOptions, text string) (iter.Seq[Token], error) { // nolint: gocognit
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
		return d.root.Tokenise(options, text)
	}

	rootTokens, err := Tokenise(Coalesce(d.root), options, rootText.String())
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
