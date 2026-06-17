package lexers

import (
	"iter"
	"strings"

	. "github.com/alecthomas/chroma/v3" // nolint
)

// HTTP lexer.
var HTTP = Register(httpBodyContentTypeLexer(MustNewLexer(
	&Config{
		Name:         "HTTP",
		Aliases:      []string{"http"},
		Filenames:    []string{},
		MimeTypes:    []string{},
		NotMultiline: true,
		DotAll:       true,
	},
	httpRules,
)))

func httpRules() Rules {
	return Rules{
		"root": {
			{`(GET|POST|PUT|DELETE|HEAD|OPTIONS|TRACE|PATCH|CONNECT)( +)([^ ]+)( +)(HTTP)(/)([123](?:\.[01])?)(\r?\n|\Z)`, ByGroups(NameFunction, Text, NameNamespace, Text, KeywordReserved, Operator, LiteralNumber, Text), Push("headers")},
			{`(HTTP)(/)([123](?:\.[01])?)( +)(\d{3})( *)([^\r\n]*)(\r?\n|\Z)`, ByGroups(KeywordReserved, Operator, LiteralNumber, Text, LiteralNumber, Text, NameException, Text), Push("headers")},
		},
		"headers": {
			{`([^\s:]+)( *)(:)( *)([^\r\n]+)(\r?\n|\Z)`, EmitterFunc(httpHeaderBlock), nil},
			{`([\t ]+)([^\r\n]+)(\r?\n|\Z)`, EmitterFunc(httpContinuousHeaderBlock), nil},
			{`\r?\n`, Text, Push("content")},
		},
		"content": {
			{`.+`, EmitterFunc(httpContentBlock), nil},
		},
	}
}

func httpContentBlock(groups []string, state *LexerState) iter.Seq[Token] {
	return Literator(Token{Generic, groups[0]})
}

func httpHeaderBlock(groups []string, state *LexerState) iter.Seq[Token] {
	return Literator(
		Token{Name, groups[1]},
		Token{Text, groups[2]},
		Token{Operator, groups[3]},
		Token{Text, groups[4]},
		Token{Literal, groups[5]},
		Token{Text, groups[6]},
	)
}

func httpContinuousHeaderBlock(groups []string, state *LexerState) iter.Seq[Token] {
	return Literator(
		Token{Text, groups[1]},
		Token{Literal, groups[2]},
		Token{Text, groups[3]},
	)
}

func httpBodyContentTypeLexer(lexer Lexer) Lexer { return &httpBodyContentTyper{lexer} }

type httpBodyContentTyper struct{ Lexer }

func (d *httpBodyContentTyper) Tokenise(options *TokeniseOptions, text string) (iter.Seq[Token], error) { // nolint: gocognit
	it, err := d.Lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}

	return func(yield func(Token) bool) {
		var contentType string
		var isContentType bool

		for token := range it {
			switch {
			case token.Type == Name && strings.ToLower(token.Value) == "content-type":
				isContentType = true

			case token.Type == Literal && isContentType:
				isContentType = false
				contentType = strings.TrimSpace(token.Value)
				if pos := strings.Index(contentType, ";"); pos > 0 {
					contentType = strings.TrimSpace(contentType[:pos])
				}

			case token.Type == Generic && contentType != "":
				lexer := MatchMimeType(contentType)

				if lexer == nil && strings.Contains(contentType, "+") {
					slashPos := strings.Index(contentType, "/")
					plusPos := strings.LastIndex(contentType, "+")
					contentType = contentType[:slashPos+1] + contentType[plusPos+1:]
					lexer = MatchMimeType(contentType)
				}

				if lexer == nil {
					token.Type = Text
				} else {
					subIt, err := lexer.Tokenise(nil, token.Value)
					if err != nil {
						panic(err)
					}
					for subToken := range subIt {
						if !yield(subToken) {
							return
						}
					}
					continue
				}
			}
			if !yield(token) {
				return
			}
		}
	}, nil
}
