package chroma

import (
	"slices"
	"testing"
	"testing/fstest"

	assert "github.com/alecthomas/assert/v2"
)

func xmlLexerFromString(t *testing.T, config, rules string) *RegexLexer {
	t.Helper()
	fs := fstest.MapFS{"lexer.xml": {Data: []byte("<lexer><config>" + config + "</config><rules><state name=\"root\">" + rules + "</state></rules></lexer>")}}
	lexer, err := NewXMLLexer(fs, "lexer.xml")
	assert.NoError(t, err)
	return lexer
}

func TestWordsMatchLikeRegex(t *testing.T) {
	const symbolSuffix = `(?![\w-])`
	xmlRules := `
		<rule words="car cdr cons ü über" suffix="` + symbolSuffix + `"><token type="NameBuiltin"/></rule>
		<rule words="if iff" prefix="(?&lt;![\w-])" suffix="\b"><token type="Keyword"/></rule>
		<rule words="(" ><bygroups><token type="Punctuation"/></bygroups></rule>
		<rule pattern="[\w-]+"><token type="Name"/></rule>
		<rule pattern="\s+"><token type="Text"/></rule>
		<rule pattern="."><token type="Other"/></rule>`
	goRules := Rules{"root": {
		{Words(``, symbolSuffix, "car", "cdr", "cons", "ü", "über"), NameBuiltin, nil},
		{Words(`(?<![\w-])`, `\b`, "if", "iff"), Keyword, nil},
		{Words(``, ``, "("), ByGroups(Punctuation), nil},
		{`[\w-]+`, Name, nil},
		{`\s+`, Text, nil},
		{`.`, Other, nil},
	}}
	input := "(car x) (CAR y) cars car-like iff if-x Iff über übe ü"
	for _, config := range []string{"", "<case_insensitive>true</case_insensitive>"} {
		t.Run(config, func(t *testing.T) {
			native := xmlLexerFromString(t, config, xmlRules)
			assert.NoError(t, native.needRules())
			assert.NotZero(t, native.rules["root"][0].words, "words rule should be matched natively")

			regex := mustNewLexer(t, native.Config(), goRules)
			nativeTokens, err := native.Tokenise(nil, input)
			assert.NoError(t, err)
			regexTokens, err := regex.Tokenise(nil, input)
			assert.NoError(t, err)
			assert.Equal(t, slices.Collect(regexTokens), slices.Collect(nativeTokens))
		})
	}
}

func TestWordsFollowCaseInsensitiveConfig(t *testing.T) {
	rules := `<rule words="car" suffix="\b"><token type="NameBuiltin"/></rule><rule pattern="\w+|\s+"><token type="Name"/></rule>`
	tests := []struct {
		config   string
		expected TokenType
	}{
		{"", Name},
		{"<case_insensitive>true</case_insensitive>", NameBuiltin},
	}
	for _, test := range tests {
		it, err := xmlLexerFromString(t, test.config, rules).Tokenise(nil, "CAR")
		assert.NoError(t, err)
		assert.Equal(t, []Token{{test.expected, "CAR"}}, slices.Collect(it))
	}
}

func TestWordsAffixesMustBeZeroWidthAssertions(t *testing.T) {
	tests := []struct {
		name  string
		rules string
		err   string
	}{
		{"ConsumingSuffix", `<rule words="a" suffix="\("><token type="Name"/></rule>`, "must be zero-width"},
		{"ConsumingPrefix", `<rule words="a" prefix="\s"><token type="Name"/></rule>`, "must be zero-width"},
		{"CapturingSuffix", `<rule words="a" suffix="(?=(x))"><token type="Name"/></rule>`, "must not capture"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := xmlLexerFromString(t, "", test.rules).Tokenise(nil, "a")
			assert.Error(t, err)
			assert.Contains(t, err.Error(), test.err)
		})
	}
}
