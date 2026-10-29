package lexers_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

func TestTopazLispexTokens(t *testing.T) {
	tests := []struct {
		language, filename, source string
		want                       map[string]chroma.TokenType
	}{
		{
			language: "Topaz",
			filename: "example.tpz",
			source:   "let name = \"Topaz\"\nlet let💎: int = 42\nlet hex = 0x2A\nlet binary = 0b10_10\nlet text = \"Hi {name}\"\nlet multiline = \"\"\"\n// string body\n\"\"\"\n/* outer /* inner */ let done = true\n",
			want: map[string]chroma.TokenType{
				"let":   chroma.KeywordDeclaration,
				"let💎": chroma.Name,
				"int":   chroma.KeywordType,
				"42":    chroma.LiteralNumberInteger,
				"0x2A": chroma.LiteralNumberHex,
				"0b10_10": chroma.LiteralNumberBin,
				"\n// string body\n": chroma.LiteralStringDouble,
				"done":  chroma.Name,
				"true":  chroma.KeywordConstant,
			},
		},
		{
			language: "Lispex",
			filename: "example.lspx",
			source:   "(define +5 1/3) #| outer #| inner |# end |#\n(list #\\\" #\\; #f \"text;\")\n",
			want: map[string]chroma.TokenType{
				"define": chroma.Keyword,
				"+5":     chroma.NameVariable,
				"1/3":    chroma.LiteralNumber,
				"#\\\"": chroma.LiteralStringChar,
				"#\\;":  chroma.LiteralStringChar,
				"#f":     chroma.NameConstant,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.language, func(t *testing.T) {
			lexer := lexers.Get(tt.language)
			if lexer == nil || lexers.Match(tt.filename) != lexer {
				t.Fatalf("missing lexer or extension mapping for %s", tt.language)
			}
			iterator, err := lexer.Tokenise(nil, tt.source)
			if err != nil {
				t.Fatal(err)
			}
			seen := map[string]bool{}
			var reconstructed strings.Builder
			for token := range iterator {
				reconstructed.WriteString(token.Value)
				if token.Type == chroma.Error {
					t.Errorf("unexpected error token %q", token.Value)
				}
				if want, ok := tt.want[token.Value]; ok {
					if token.Type != want {
						t.Errorf("%q: got %s, want %s", token.Value, token.Type, want)
					}
					seen[token.Value] = true
				}
			}
			if reconstructed.String() != tt.source {
				t.Error("tokenization changed the source text")
			}
			for value := range tt.want {
				if !seen[value] {
					t.Errorf("missing expected token %q", value)
				}
			}
			t.Run("CompleteProgram", func(t *testing.T) {
				source, err := os.ReadFile(filepath.Join("testdata", strings.ToLower(tt.language)+".actual"))
				if err != nil {
					t.Fatal(err)
				}
				iterator, err := lexer.Tokenise(nil, string(source))
				if err != nil {
					t.Fatal(err)
				}
				var reconstructed strings.Builder
				for token := range iterator {
					reconstructed.WriteString(token.Value)
					if token.Type == chroma.Error {
						t.Errorf("unexpected error token in complete example: %q", token.Value)
					}
				}
				if reconstructed.String() != string(source) {
					t.Error("tokenization changed the complete example")
				}
			})
		})
	}
}
