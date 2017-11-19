package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
	"strings"
)

var httpBodyContentType string

// Http lexer.
var Http = Register(MustNewLexer(
	&Config{
		Name:         "HTTP",
		Aliases:      []string{"http"},
		Filenames:    []string{},
		MimeTypes:    []string{},
		NotMultiline: true,
		DotAll:       true,
	},
	Rules{
		"root": {
			{`(GET|POST|PUT|DELETE|HEAD|OPTIONS|TRACE|PATCH)( +)([^ ]+)( +)(HTTP)(/)(1\.[01])(\r?\n|\Z)`, ByGroups(NameFunction, Text, NameNamespace, Text, KeywordReserved, Operator, LiteralNumber, Text), Push("headers")},
			{`(HTTP)(/)(1\.[01])( +)(\d{3})( +)([^\r\n]+)(\r?\n|\Z)`, ByGroups(KeywordReserved, Operator, LiteralNumber, Text, LiteralNumber, Text, NameException, Text), Push("headers")},
		},
		"headers": {
			{`([^\s:]+)( *)(:)( *)([^\r\n]+)(\r?\n|\Z)`, EmitterFunc(httpHeaderBlock), nil},
			{`([\t ]+)([^\r\n]+)(\r?\n|\Z)`, EmitterFunc(httpContinuousHeaderBlock), nil},
			{`\r?\n`, Text, Push("content")},
		},
		"content": {
			{`.+`, EmitterFunc(httpContentBlock), nil},
		},
	},
))

func httpContentBlock(groups []string, lexer Lexer) Iterator {
	iterators := []Iterator{}
	code := groups[0]

	if len(httpBodyContentType) > 0 {
		lexer := MatchMimeType(httpBodyContentType)

		// application/calendar+xml can be treated as application/xml
		// if there's not a better match.
		if lexer == nil && strings.Contains(httpBodyContentType, "+") {
			slashPos := strings.Index(httpBodyContentType, "/")
			plusPos := strings.LastIndex(httpBodyContentType, "+")
			httpBodyContentType = httpBodyContentType[:slashPos+1] + httpBodyContentType[plusPos+1:]
			lexer = MatchMimeType(httpBodyContentType)
		}

		if lexer != nil {
			sub, err := lexer.Tokenise(nil, code)
			if err != nil {
				panic(err)
			}
			iterators = append(iterators, sub)
		} else {
			tokens := []*Token{
				{Text, code},
			}
			iterators = append(iterators, Literator(tokens...))
		}
	}
	return Concaterator(iterators...)
}

func httpHeaderBlock(groups []string, lexer Lexer) Iterator {
	if strings.ToLower(groups[1]) == "content-type" {
		contentType := strings.TrimSpace(groups[5])
		pos := strings.Index(contentType, ";")
		if pos > 0 {
			contentType = strings.TrimSpace(contentType[:pos])
		}

		httpBodyContentType = contentType
	}

	iterators := []Iterator{}
	tokens := []*Token{
		{Name, groups[1]},
		{Text, groups[2]},
		{Operator, groups[3]},
		{Text, groups[4]},
		{Literal, groups[5]},
		{Text, groups[6]},
	}
	iterators = append(iterators, Literator(tokens...))
	return Concaterator(iterators...)
}

func httpContinuousHeaderBlock(groups []string, lexer Lexer) Iterator {
	iterators := []Iterator{}
	tokens := []*Token{
		{Text, groups[1]},
		{Literal, groups[2]},
		{Text, groups[3]},
	}
	iterators = append(iterators, Literator(tokens...))
	return Concaterator(iterators...)
}
