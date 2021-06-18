package g

import (
	"testing"

	"github.com/alecthomas/chroma"
	"github.com/stretchr/testify/assert"
)

func TestGoHTMLTemplateIssue126(t *testing.T) {
	for _, source := range []string{
		`<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{ if eq  .Title  .Site.Title }}{{ .Site.Title }}{{ else }}{{ with .Title }}{{.}} on {{ end }}{{ .Site.Title }}{{ end }}</title>
    <link>{{ .Permalink }}</link>
    <description>Recent content {{ if ne  .Title  .Site.Title }}{{ with .Title }}in {{.}} {{ end }}{{ end }}on {{ .Site.Title }}</description>
    <generator>Hugo -- gohugo.io</generator>{{ with .Site.LanguageCode }}
    <language>{{.}}</language>{{end}}{{ with .Site.Author.email }}
    <managingEditor>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</managingEditor>{{end}}{{ with .Site.Author.email }}
    <webMaster>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</webMaster>{{end}}{{ with .Site.Copyright }}
    <copyright>{{.}}</copyright>{{end}}{{ if not .Date.IsZero }}
    <lastBuildDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</lastBuildDate>{{ end }}
    {{ with .OutputFormats.Get "RSS" }}
        {{ printf "<atom:link href=%q rel=\"self\" type=%q />" .Permalink .MediaType | safeHTML }}
    {{ end }}
	{{/*
		Print all pages
	*/}}
    {{ range .Data.Pages }}
    <item>
      <title>{{ .Title }}</title>
      <link>{{ .Permalink }}</link>
      <pubDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</pubDate>
      {{ with .Site.Author.email }}<author>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</author>{{end}}
      <guid>{{ .Permalink }}</guid>
      <description>{{ .Summary | html }}</description>
    </item>
    {{ end }}
  </channel>
</rss>
`,
		`{{ $headless := .Site.GetPage "page" "some-headless-bundle" }}
{{ $reusablePages := $headless.Resources.Match "author*" }}
<h2>Authors</h2>
{{ range $reusablePages }}
    <h3>{{ .Title }}</h3>
    {{ .Content }}
{{ end }}`} {
		tokens, err := chroma.Tokenise(GoHTMLTemplate, nil, source)
		assert.NoError(t, err)
		assert.Equal(t, source, chroma.Stringify(tokens...))
	}
}

func TestGoHTMLTemplateMultilineComments(t *testing.T) {
	for _, source := range []string{
		`
{{/*
	This is a multiline comment
*/}}
`,
		`
{{- /*
	This is a multiline comment
*/}}
`,
		`
{{/*
	This is a multiline comment
*/ -}}
`,
		`
{{- /*
	This is a multiline comment
*/ -}}
`,
	} {
		tokens, err := chroma.Tokenise(GoHTMLTemplate, nil, source)
		assert.NoError(t, err)
		assert.Equal(t, source, chroma.Stringify(tokens...))

		// Make sure that there are no errors
		for _, token := range tokens {
			assert.NotEqual(t, chroma.Error, token.Type)
		}

		// Make sure that multiline comments are printed
		found := false
		for _, token := range tokens {
			if token.Type == chroma.CommentMultiline {
				found = true
			}
		}
		assert.True(t, found)
	}
}

func TestGoHTMLTemplateNegativeNumber(t *testing.T) {
	for _, source := range []string{
		`
{{ fn -3 }}
`,
	} {
		tokens, err := chroma.Tokenise(GoHTMLTemplate, nil, source)
		assert.NoError(t, err)
		assert.Equal(t, source, chroma.Stringify(tokens...))

		// Make sure that there are no errors
		for _, token := range tokens {
			assert.NotEqual(t, chroma.Error, token.Type)
		}

		// Make sure that negative number is found
		found := false
		for _, token := range tokens {
			if token.Type == chroma.LiteralNumberInteger {
				found = true
			}
		}
		assert.True(t, found)
	}
}
