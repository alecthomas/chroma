package lexers_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/lexers"
)

func TestRustCStringLiterals(t *testing.T) {
	for _, tt := range []struct {
		name    string
		literal string
		want    []chroma.Token
	}{
		{"c", `c"hello"`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "c"}, {Type: chroma.LiteralString, Value: `"hello"`}}},
		{"c_empty", `c""`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "c"}, {Type: chroma.LiteralString, Value: `""`}}},
		{"c_escapes", `c"\xFF\u{00E6}\""`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "c"}, {Type: chroma.LiteralString, Value: `"`}, {Type: chroma.LiteralStringEscape, Value: `\xFF\u{00E6}\"`}, {Type: chroma.LiteralString, Value: `"`}}},
		{"c_unicode", "c\"æ\u4e2d\"", []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "c"}, {Type: chroma.LiteralString, Value: "\"æ\u4e2d\""}}},
		{"c_multiline", "c\"first\nsecond\"", []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "c"}, {Type: chroma.LiteralString, Value: "\"first\nsecond\""}}},
		{"raw_c", `cr"hello"`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "cr"}, {Type: chroma.LiteralString, Value: `"hello"`}}},
		{"raw_c_hash", `cr#""hello""#`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "cr"}, {Type: chroma.LiteralString, Value: `#""hello""#`}}},
		{"raw_c_two_hashes", `cr##"value "# is not the end; \xFF"##`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "cr"}, {Type: chroma.LiteralString, Value: `##"value "# is not the end; \xFF"##`}}},
		{"raw_c_multiline", "cr#\"first\nsecond\"#", []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "cr"}, {Type: chroma.LiteralString, Value: "#\"first\nsecond\"#"}}},
		{"byte_control", `b"\xFF"`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "b"}, {Type: chroma.LiteralString, Value: `"`}, {Type: chroma.LiteralStringEscape, Value: `\xFF`}, {Type: chroma.LiteralString, Value: `"`}}},
		{"raw_byte_control", `br#"\xFF"#`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "br"}, {Type: chroma.LiteralString, Value: `#"\xFF"#`}}},
		{"raw_control", `r#"\xFF"#`, []chroma.Token{{Type: chroma.LiteralStringAffix, Value: "r"}, {Type: chroma.LiteralString, Value: `#"\xFF"#`}}},
		{"string_control", `"hello"`, []chroma.Token{{Type: chroma.LiteralString, Value: `"hello"`}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			source := tt.literal + " + tail\n"
			it, err := chroma.Coalesce(lexers.Get("rust")).Tokenise(nil, source)
			if err != nil {
				t.Fatal(err)
			}
			tokens := slices.Collect(it)
			if got := chroma.Stringify(tokens...); got != source {
				t.Fatalf("source changed: got %q, want %q", got, source)
			}
			tokens = slices.DeleteFunc(tokens, func(token chroma.Token) bool {
				return token.Type == chroma.TextWhitespace && strings.TrimSpace(token.Value) == ""
			})
			want := append(slices.Clone(tt.want), chroma.Token{Type: chroma.Operator, Value: "+"}, chroma.Token{Type: chroma.Name, Value: "tail"})
			if !slices.Equal(tokens, want) {
				t.Fatalf("tokens = %#v, want %#v", tokens, want)
			}
		})
	}
}
