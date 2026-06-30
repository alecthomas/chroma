package lexers_test

import (
	"testing"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

func TestAstroLexerMatchesFilenames(t *testing.T) {
	if lexer := lexers.Match("component.astro"); lexer == nil || lexer.Config().Name != "Astro" {
		t.Fatalf("expected Astro lexer for .astro files, got %v", lexer)
	}
}

func TestAstroDynamicAttributeExpressions(t *testing.T) {
	lexer := lexers.Get("astro")
	if lexer == nil {
		t.Fatal("expected Astro lexer to be registered")
	}

	tokens, err := chroma.Tokenise(lexer, nil, `<Component title={title}>Hello {title}</Component>`)
	if err != nil {
		t.Fatal(err)
	}

	for _, token := range tokens {
		if token.Type == chroma.Error {
			t.Fatalf("unexpected error token: %#v", token)
		}
	}
}

func TestAstroFrontmatterAllowsCRLF(t *testing.T) {
	tokens := mustTokeniseAstro(t, "---\r\ninterface Props {\r\n\ttitle: string;\r\n}\r\n---\r\n<h1>{title}</h1>")

	if !hasToken(tokens, chroma.KeywordReserved, "interface") {
		t.Fatalf("expected CRLF frontmatter to be tokenised as TypeScript, got %#v", tokens)
	}
}

func TestAstroDefaultScriptUsesTypeScript(t *testing.T) {
	tokens := mustTokeniseAstro(t, `<script>
const message: string = "loaded";
</script>`)

	if !hasToken(tokens, chroma.KeywordType, "string") {
		t.Fatalf("expected default script blocks to be tokenised as TypeScript, got %#v", tokens)
	}
}

func mustTokeniseAstro(t *testing.T, source string) []chroma.Token {
	t.Helper()

	lexer := lexers.Get("astro")
	if lexer == nil {
		t.Fatal("expected Astro lexer to be registered")
	}

	tokens, err := chroma.Tokenise(lexer, nil, source)
	if err != nil {
		t.Fatal(err)
	}

	for _, token := range tokens {
		if token.Type == chroma.Error {
			t.Fatalf("unexpected error token: %#v", token)
		}
	}

	return tokens
}

func hasToken(tokens []chroma.Token, tokenType chroma.TokenType, value string) bool {
	for _, token := range tokens {
		if token.Type == tokenType && token.Value == value {
			return true
		}
	}
	return false
}
