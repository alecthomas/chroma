package chroma

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmitterSerialisationRoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		emitter Emitter
	}{
		{"ByGroups", ByGroups(Name, Using("Go"))},
		{"UsingSelf", UsingSelf("root")},
		{"Using", Using("Go")},
		{"UsingByGroup", UsingByGroup(1, 2, Name)},
		{"TokenType", Name},
	}
	for _, test := range tests {
		// nolint: scopelint
		t.Run(test.name, func(t *testing.T) {
			data, err := xml.Marshal(test.emitter)
			require.NoError(t, err)
			t.Logf("%s", data)
			value, target := newFromTemplate(test.emitter)
			err = xml.Unmarshal(data, target)
			require.NoError(t, err)
			require.Equal(t, test.emitter, value())
		})
	}
}

func TestMutatorSerialisationRoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		mutator Mutator
	}{
		{"Include", Include("string").Mutator},
		{"Combined", Combined("a", "b", "c")},
		{"Multi", Mutators(Include("string").Mutator, Push("quote"))},
		{"Push", Push("include")},
		{"Pop", Pop(1)},
	}
	for _, test := range tests {
		// nolint: scopelint
		t.Run(test.name, func(t *testing.T) {
			data, err := xml.Marshal(test.mutator)
			require.NoError(t, err)
			t.Logf("%s", data)
			value, target := newFromTemplate(test.mutator)
			err = xml.Unmarshal(data, target)
			require.NoError(t, err)
			require.Equal(t, test.mutator, value())
		})
	}
}

func TestMarshal(t *testing.T) {
	actual := MustNewLexer(&Config{
		Name:      "PkgConfig",
		Aliases:   []string{"pkgconfig"},
		Filenames: []string{"*.pc"},
	}, func() Rules {
		return Rules{
			"root": {
				{`#.*$`, CommentSingle, nil},
				{`^(\w+)(=)`, ByGroups(NameAttribute, Operator), nil},
				{`^([\w.]+)(:)`, ByGroups(NameTag, Punctuation), Push("spvalue")},
				Include("interp"),
				{`[^${}#=:\n.]+`, Text, nil},
				{`.`, Text, nil},
			},
			"interp": {
				{`\$\$`, Text, nil},
				{`\$\{`, LiteralStringInterpol, Push("curly")},
			},
			"curly": {
				{`\}`, LiteralStringInterpol, Pop(1)},
				{`\w+`, NameAttribute, nil},
			},
			"spvalue": {
				Include("interp"),
				{`#.*$`, CommentSingle, Pop(1)},
				{`\n`, Text, Pop(1)},
				{`[^${}#\n]+`, Text, nil},
				{`.`, Text, nil},
			},
		}
	})
	data, err := Marshal(actual)
	require.NoError(t, err)
	expected, err := Unmarshal(data)
	require.NoError(t, err)
	require.Equal(t, expected.Config(), actual.Config())
	require.Equal(t, mustRules(t, expected), mustRules(t, actual))
}

func mustRules(t testing.TB, r *RegexLexer) Rules {
	t.Helper()
	rules, err := r.Rules()
	require.NoError(t, err)
	return rules
}

func TestRuleSerialisation(t *testing.T) {
	tests := []Rule{
		Include("String"),
		{`\d+`, Text, nil},
		{`"`, String, Push("String")},
	}
	for _, test := range tests {
		data, err := xml.Marshal(test)
		require.NoError(t, err)
		t.Log(string(data))
		actual := Rule{}
		err = xml.Unmarshal(data, &actual)
		require.NoError(t, err)
		require.Equal(t, test, actual)
	}
}

func TestRulesSerialisation(t *testing.T) {
	expected := Rules{
		"root": {
			{`#.*$`, CommentSingle, nil},
			{`^(\w+)(=)`, ByGroups(NameAttribute, Operator), nil},
			{`^([\w.]+)(:)`, ByGroups(NameTag, Punctuation), Push("spvalue")},
			Include("interp"),
			{`[^${}#=:\n.]+`, Text, nil},
			{`.`, Text, nil},
		},
		"interp": {
			{`\$\$`, Text, nil},
			{`\$\{`, LiteralStringInterpol, Push("curly")},
		},
		"curly": {
			{`\}`, LiteralStringInterpol, Pop(1)},
			{`\w+`, NameAttribute, nil},
		},
		"spvalue": {
			Include("interp"),
			{`#.*$`, CommentSingle, Pop(1)},
			{`\n`, Text, Pop(1)},
			{`[^${}#\n]+`, Text, nil},
			{`.`, Text, nil},
		},
	}
	data, err := xml.MarshalIndent(expected, "  ", "  ")
	require.NoError(t, err)
	re := regexp.MustCompile(`></[a-zA-Z]+>`)
	data = re.ReplaceAll(data, []byte(`/>`))
	b := &bytes.Buffer{}
	w := gzip.NewWriter(b)
	fmt.Fprintln(w, string(data))
	w.Close()
	actual := Rules{}
	err = xml.Unmarshal(data, &actual)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}
