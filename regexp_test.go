package chroma

import (
	"fmt"
	"iter"
	"runtime"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	assert "github.com/alecthomas/assert/v2"
	"github.com/dlclark/regexp2/v2"
)

func mustNewLexer(t *testing.T, config *Config, rules Rules) *RegexLexer { // nolint: forbidigo
	t.Helper()
	lexer, err := NewLexer(config, func() Rules {
		return rules
	})
	assert.NoError(t, err)
	return lexer
}

func TestWords(t *testing.T) {
	tests := []struct {
		name           string
		prefix, suffix string
		words          []string
		input          string
		match, word    string
	}{
		{"PrefersLongest", ``, ``, []string{"fo", "foobar", "foo"}, "foobarx", "foobar", "foobar"},
		{"BacktracksToShorterWord", ``, `b`, []string{"a", "ab"}, "ab", "ab", "a"},
		{"EscapesMetacharacters", ``, ``, []string{"1+", "*", "<="}, "1+2", "1+", "1+"},
		{"Multibyte", ``, ``, []string{"über", "übel"}, "übel", "übel", "übel"},
		{"PrefixAndSuffix", `\(`, `\b`, []string{"if", "iff"}, "(iff)", "(iff", "iff"},
		{"NoMatch", ``, `\b`, []string{"car", "cdr"}, "cart", "", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			re := regexp2.MustCompile(`\G`+Words(test.prefix, test.suffix, test.words...), regexp2.None)
			match, err := re.FindStringMatch(test.input)
			assert.NoError(t, err)
			if test.match == "" {
				assert.Zero(t, match)
				return
			}
			assert.NotZero(t, match)
			assert.Equal(t, test.match, match.String())
			assert.Equal(t, test.word, match.GroupByNumber(1).String())
		})
	}
}

func TestNewlineAtEndOfFile(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{EnsureNL: true}, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err := l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Keyword, "hello"}, {Whitespace, "\n"}}, slices.Collect(it))

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = l.Tokenise(nil, `hello`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, "hello"}}, slices.Collect(it))
}

func TestMatchingAtStart(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`\s+`, Whitespace, nil},
			{`^-`, Punctuation, Push("directive")},
			{`->`, Operator, nil},
		},
		"directive": {
			{"module", NameEntity, Pop(1)},
		},
	}))
	it, err := l.Tokenise(nil, `-module ->`)
	assert.NoError(t, err)
	assert.Equal(t,
		[]Token{{Punctuation, "-"}, {NameEntity, "module"}, {Whitespace, " "}, {Operator, "->"}},
		slices.Collect(it))
}

func TestZeroWidthMatchTerminates(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`a*`, Text, nil},
		},
	}))
	it, err := l.Tokenise(nil, "b")
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, "b"}}, slices.Collect(it))
}

func TestZeroWidthPushPopCycleTerminates(t *testing.T) {
	lexer := mustNewLexer(t, &Config{Name: "loopy"}, Rules{ // nolint: forbidigo
		"root": {
			{`(?=\S)`, None, Push("a")},
		},
		"a": {
			{`(?=\S)`, None, Push("b")},
		},
		"b": {
			{``, None, Pop(1)},
		},
	})
	it, err := lexer.Tokenise(nil, "x")
	assert.NoError(t, err)
	tokens := collectTokensTimeout(t, it, 2*time.Second)
	var got []Token
	for _, tok := range tokens {
		if tok.Value != "" {
			got = append(got, tok)
		}
	}
	assert.Equal(t, []Token{{Error, "x"}}, got)
}

func collectTokensTimeout(t *testing.T, seq iter.Seq[Token], timeout time.Duration) []Token {
	t.Helper()
	done := make(chan []Token, 1)
	go func() {
		done <- slices.Collect(seq)
	}()
	select {
	case tokens := <-done:
		return tokens
	case <-time.After(timeout):
		t.Fatal("Iterator never returned: zero-width push/pop cycle")
		return nil
	}
}

func TestEnsureLFOption(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\r?\n|\r)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err := l.Tokenise(&TokeniseOptions{
		State:    "root",
		EnsureLF: true,
	}, "hello\r\nworld\r")
	assert.NoError(t, err)
	assert.Equal(t, []Token{
		{Keyword, "hello"},
		{Whitespace, "\n"},
		{Keyword, "world"},
		{Whitespace, "\n"},
	}, slices.Collect(it))

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{`(\w+)(\r?\n|\r)`, ByGroups(Keyword, Whitespace), nil},
		},
	}))
	it, err = l.Tokenise(&TokeniseOptions{
		State:    "root",
		EnsureLF: false,
	}, "hello\r\nworld\r")
	assert.NoError(t, err)
	assert.Equal(t, []Token{
		{Keyword, "hello"},
		{Whitespace, "\r\n"},
		{Keyword, "world"},
		{Whitespace, "\r"},
	}, slices.Collect(it))
}

func TestEnsureLFFunc(t *testing.T) {
	tests := []struct{ in, out string }{
		{in: "", out: ""},
		{in: "abc", out: "abc"},
		{in: "\r", out: "\n"},
		{in: "a\r", out: "a\n"},
		{in: "\rb", out: "\nb"},
		{in: "a\rb", out: "a\nb"},
		{in: "\r\n", out: "\n"},
		{in: "a\r\n", out: "a\n"},
		{in: "\r\nb", out: "\nb"},
		{in: "a\r\nb", out: "a\nb"},
		{in: "\r\r\r\n\r", out: "\n\n\n\n"},
	}
	for _, test := range tests {
		out := ensureLF(test.in)
		assert.Equal(t, out, test.out)
	}
}

func TestByGroupNames(t *testing.T) {
	l := Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)(?<operator>=)(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:      String,
					`operator`: Operator,
					`value`:    String,
				}),
				nil,
			},
		},
	}))
	it, err := l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc`}, {Operator, `=`}, {String, `123`}}, slices.Collect(it))

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)(?<operator>=)(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:   String,
					`value`: String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc`}, {Error, `=`}, {String, `123`}}, slices.Collect(it))

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)=(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:   String,
					`value`: String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc123`}}, slices.Collect(it))

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`(?<key>\w+)(?<op>=)(?<value>\w+)`,
				ByGroupNames(map[string]Emitter{
					`key`:      String,
					`operator`: Operator,
					`value`:    String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{String, `abc`}, {Error, `=`}, {String, `123`}}, slices.Collect(it))

	l = Coalesce(mustNewLexer(t, nil, Rules{ // nolint: forbidigo
		"root": {
			{
				`\w+=\w+`,
				ByGroupNames(map[string]Emitter{
					`key`:      String,
					`operator`: Operator,
					`value`:    String,
				}),
				nil,
			},
		},
	}))
	it, err = l.Tokenise(nil, `abc=123`)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Error, `abc=123`}}, slices.Collect(it))
}

func TestIgnoreToken(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{EnsureNL: true}, Rules{ // nolint: forbidigo
		"root": {
			{`(\s*)(\w+)(?:\1)(\n)`, ByGroups(Ignore, Keyword, Whitespace), nil},
		},
	}))
	it, err := l.Tokenise(nil, `  hello  `)
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Keyword, "hello"}, {TextWhitespace, "\n"}}, slices.Collect(it))
}

func TestPushToUndefinedStateIsAnError(t *testing.T) {
	lexer, err := NewLexer(&Config{Name: "test"}, func() Rules {
		return Rules{
			"root": {
				{`a`, Text, Push("nonexistent")},
			},
		}
	})
	assert.NoError(t, err)
	_, err = lexer.Tokenise(nil, "a")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `invalid push state "nonexistent"`)

	// #pop is a valid pseudo-state and must not be rejected.
	lexer = mustNewLexer(t, &Config{Name: "test"}, Rules{ // nolint: forbidigo
		"root": {
			{`a`, Text, Push("other")},
		},
		"other": {
			{`b`, Text, Push("#pop")},
		},
	})
	it, err := lexer.Tokenise(nil, "ab")
	assert.NoError(t, err)
	assert.Equal(t, []Token{{Text, "a"}, {Text, "b"}}, slices.Collect(it))
}

func TestLongTokensMatch(t *testing.T) {
	l := Coalesce(mustNewLexer(t, &Config{}, Rules{ // nolint: forbidigo
		"root": {
			{`"(\\\\|\\"|[^"])*"`, LiteralStringDouble, nil},
			{`/\*(.|\n)*?\*/`, CommentMultiline, nil},
		},
	}))
	for _, token := range []Token{
		{LiteralStringDouble, `"` + strings.Repeat("a", 100_000) + `"`},
		{CommentMultiline, "/*" + strings.Repeat("a\n", 50_000) + "*/"},
	} {
		it, err := l.Tokenise(nil, token.Value)
		assert.NoError(t, err)
		assert.Equal(t, []Token{token}, slices.Collect(it))
	}
}

func TestTokeniseDefaultsEmptyState(t *testing.T) {
	for _, entry := range []struct {
		name string
		call func(Lexer, *TokeniseOptions, string) ([]Token, error)
	}{
		{"method", func(l Lexer, options *TokeniseOptions, text string) ([]Token, error) {
			it, err := l.Tokenise(options, text)
			if err != nil {
				return nil, err
			}
			return slices.Collect(it), nil
		}},
		{"helper", func(l Lexer, options *TokeniseOptions, text string) ([]Token, error) {
			it, err := Tokenise(l, options, text)
			if err != nil {
				return nil, err
			}
			return slices.Collect(it), nil
		}},
	} {
		t.Run(entry.name, func(t *testing.T) {
			for _, test := range []struct {
				name    string
				options *TokeniseOptions
				text    string
				want    []Token
			}{
				{"nil", nil, "word\r\n", []Token{{Keyword, "word"}, {Whitespace, "\n"}}},
				{"zero", &TokeniseOptions{}, "word\r\n", []Token{{Keyword, "word"}, {Whitespace, "\r\n"}}},
				{"normalize", &TokeniseOptions{EnsureLF: true}, "word\rword\r\n", []Token{{Keyword, "word"}, {Whitespace, "\n"}, {Keyword, "word"}, {Whitespace, "\n"}}},
				{"preserve", &TokeniseOptions{EnsureLF: false}, "word\rword\r\n", []Token{{Keyword, "word"}, {Whitespace, "\r"}, {Keyword, "word"}, {Whitespace, "\r\n"}}},
				{"explicit-root", &TokeniseOptions{State: "root"}, "word", []Token{{Keyword, "word"}}},
				{"explicit-other", &TokeniseOptions{State: "other"}, "word", []Token{{Name, "word"}}},
				{"empty-input", &TokeniseOptions{}, "", nil},
			} {
				t.Run(test.name, func(t *testing.T) {
					defer func() {
						if p := recover(); p != nil {
							t.Fatalf("Tokenise panicked: %v", p)
						}
					}()
					l := mustNewLexer(t, &Config{Name: "default-state"}, Rules{
						"root":  {{`[a-z]+`, Keyword, nil}, {`\r\n|\r|\n`, Whitespace, nil}},
						"other": {{`[a-z]+`, Name, nil}},
					})
					var before TokeniseOptions
					if test.options != nil {
						before = *test.options
					}
					got, err := entry.call(l, test.options, test.text)
					if err != nil {
						t.Fatal(err)
					}
					if !slices.Equal(got, test.want) {
						t.Fatalf("tokens: got %#v, want %#v", got, test.want)
					}
					if test.options != nil && *test.options != before {
						t.Fatalf("caller options changed: got %+v, want %+v", *test.options, before)
					}
				})
			}
		})
	}
}

func TestTokeniseEmptyStatePreservesNested(t *testing.T) {
	for _, nested := range []bool{false, true} {
		t.Run(fmt.Sprint(nested), func(t *testing.T) {
			defer func() {
				if p := recover(); p != nil {
					t.Fatalf("Tokenise panicked: %v", p)
				}
			}()
			l := mustNewLexer(t, &Config{Name: "nested-default-state", EnsureNL: true}, Rules{
				"root": {{`word\n`, Keyword, nil}, {`word`, Name, nil}},
			})
			options := TokeniseOptions{Nested: nested}
			it, err := l.Tokenise(&options, "word")
			if err != nil {
				t.Fatal(err)
			}
			want := []Token{{Keyword, "word\n"}}
			if nested {
				want = []Token{{Name, "word"}}
			}
			if got := slices.Collect(it); !slices.Equal(got, want) {
				t.Fatalf("tokens: got %#v, want %#v", got, want)
			}
			if options != (TokeniseOptions{Nested: nested}) {
				t.Fatalf("caller options changed: %+v", options)
			}
		})
	}
}

func TestTokeniseEmptyStateResetsAfterUnmatchedNewline(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("Tokenise panicked: %v", p)
		}
	}()
	l := mustNewLexer(t, &Config{Name: "reset-default-state"}, Rules{
		"root":   {{`"`, Punctuation, Push("string")}, {`[a-z]+`, Keyword, nil}, {`\n`, Whitespace, nil}},
		"string": {{`[a-z]+`, LiteralString, nil}},
	})
	options := TokeniseOptions{}
	it, err := l.Tokenise(&options, "\"abc\nword")
	if err != nil {
		t.Fatal(err)
	}
	want := []Token{{Punctuation, "\""}, {LiteralString, "abc"}, {Whitespace, "\n"}, {Keyword, "word"}}
	if got := slices.Collect(it); !slices.Equal(got, want) {
		t.Fatalf("tokens: got %#v, want %#v", got, want)
	}
	if options != (TokeniseOptions{}) {
		t.Fatalf("caller options changed: %+v", options)
	}
}

func TestTokeniseSharedEmptyStateOptions(t *testing.T) {
	l := mustNewLexer(t, &Config{Name: "shared-default-state"}, Rules{
		"root": {{`[a-z]+`, Keyword, nil}},
	})
	options := TokeniseOptions{}
	results := make(chan error, 16)
	var wg sync.WaitGroup
	for range cap(results) {
		wg.Go(func() {
			defer func() {
				if p := recover(); p != nil {
					results <- fmt.Errorf("Tokenise panicked: %v", p)
				}
			}()
			it, err := l.Tokenise(&options, "word")
			if err != nil {
				results <- err
				return
			}
			if got := slices.Collect(it); !slices.Equal(got, []Token{{Keyword, "word"}}) {
				results <- fmt.Errorf("unexpected tokens: %#v", got)
			}
		})
	}
	wg.Wait()
	close(results)
	for err := range results {
		t.Error(err)
	}
	if options != (TokeniseOptions{}) {
		t.Fatalf("shared caller options changed: %+v", options)
	}
}

func TestEnsureNLEmptyStack(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		input   string
		want    []Token
	}{
		{
			name:    "PopAtEndWithoutTrailingNewline",
			pattern: `hello`,
			input:   "hello",
			want:    []Token{{Keyword, "hello"}},
		},
		{
			name:    "PopMidInputWithoutTrailingNewline",
			pattern: `hel`,
			input:   "hello",
			want:    []Token{{Keyword, "hel"}, {Error, "lo"}},
		},
		{
			name:    "PopConsumingSyntheticNewline",
			pattern: `hello\n`,
			input:   "hello",
			want:    []Token{{Keyword, "hello\n"}},
		},
		{
			name:    "PopMidInputWithRealTrailingNewline",
			pattern: `hel`,
			input:   "hello\n",
			want:    []Token{{Keyword, "hel"}, {Error, "lo\n"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := mustNewLexer(t, &Config{Name: "ensure-nl-empty-stack", EnsureNL: true}, Rules{
				"root": {
					{test.pattern, Keyword, Pop(1)},
				},
			})
			it, err := l.Tokenise(nil, test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.want, slices.Collect(it))
		})
	}
}

func TestMatchContentionPreemption(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping contention test in short mode")
	}
	old := runtime.GOMAXPROCS(1)
	defer runtime.GOMAXPROCS(old)

	done := make(chan struct{})
	defer close(done)
	for range 64 {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					runtime.Gosched()
				}
			}
		}()
	}

	l := mustNewLexer(t, &Config{Name: "test-contention"}, Rules{
		"root": {
			{`\n`, Whitespace, nil},
			{`\s+`, Whitespace, nil},
			{`//(\\n|[\s\S])*?\n`, Comment, nil},
			{`(?s)/\*.*?\*/`, Comment, nil},
			{`<<=|>>=|<<|>>|<=|>=|&\^=|&\^|--|\+\+|&&|\|\||<-|->|==|!=|:=|[-+/*%&|^<>=!~]`, Punctuation, nil},
			{`[(){}\[\],.;]`, Punctuation, nil},
			{`(?i)\b(?:break|default|func|interface|select|case|defer|go|map|struct|else|goto|package|switch|fallthrough|if|range|type|continue|for|import|return|var)\b`, Keyword, nil},
			{`"[^"\\]*(?:\\.[^"\\]*)*"`, LiteralString, nil},
			{`[a-zA-Z_]\w*`, Name, nil},
		},
	})

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("Tokenise panicked: %v", p)
		}
	}()

	it, err := l.Tokenise(nil, strings.Repeat("s := \"a\" + \"b\"\n", 8000))
	assert.NoError(t, err)
	count := 0
	for range it {
		count++
	}
	assert.True(t, count > 0)
}

func TestMatchTimeoutCatastrophicBacktracking(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping catastrophic backtracking test in short mode")
	}
	l := mustNewLexer(t, &Config{Name: "catastrophic"}, Rules{
		"root": {
			{`(a+)+$`, Keyword, nil},
		},
	})
	defer func() {
		p := recover()
		assert.NotZero(t, p)
		err, ok := p.(error)
		assert.True(t, ok)
		assert.Contains(t, err.Error(), "match timeout")
	}()

	it, err := l.Tokenise(nil, strings.Repeat("a", 30)+"!")
	assert.NoError(t, err)
	for range it {
	}
	t.Fatal("expected match timeout panic on catastrophic backtracking, but got none")
}
