package chroma

import (
	"encoding/xml"
	"testing"

	assert "github.com/alecthomas/assert/v2"
)

func TestXMLEmitterUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected Emitter
	}{
		{"ByGroups", `<bygroups><token type="Name"/><using lexer="Go"/></bygroups>`, ByGroups(Name, Using("Go"))},
		{"UsingSelf", `<usingself state="root"/>`, UsingSelf("root")},
		{"Using", `<using lexer="Go"/>`, Using("Go")},
		{
			"UsingByGroup",
			`<usingbygroup><sublexer_name_group>1</sublexer_name_group><code_group>2</code_group><emitters><token type="Name"/></emitters></usingbygroup>`,
			UsingByGroup(1, 2, Name),
		},
		{"TokenType", `<token type="Name"/>`, Name},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Rule{}
			err := xml.Unmarshal([]byte(`<rule>`+test.xml+`</rule>`), &actual)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual.Type)
		})
	}
}

func TestXMLMutatorUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected Mutator
	}{
		{"Include", `<include state="string"/>`, Include("string").Mutator},
		{"Combined", `<combined state="a" state="b" state="c"/>`, Combined("a", "b", "c")},
		{
			"Multi",
			`<mutators><include state="string"/><push state="quote"/></mutators>`,
			Mutators(Include("string").Mutator, Push("quote")),
		},
		{"Push", `<push state="include"/>`, Push("include")},
		{"Pop", `<pop depth="1"/>`, Pop(1)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Rule{}
			err := xml.Unmarshal([]byte(`<rule>`+test.xml+`</rule>`), &actual)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual.Mutator)
		})
	}
}

func TestXMLRuleUnmarshal(t *testing.T) {
	tests := []struct {
		xml      string
		expected Rule
	}{
		{`<rule><include state="String"/></rule>`, Include("String")},
		{`<rule pattern="\d+"><token type="Text"/></rule>`, Rule{`\d+`, Text, nil}},
		{`<rule pattern="&#34;"><token type="LiteralString"/><push state="String"/></rule>`, Rule{`"`, String, Push("String")}},
	}
	for _, test := range tests {
		actual := Rule{}
		err := xml.Unmarshal([]byte(test.xml), &actual)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, actual)
	}
}

func TestXMLRulesUnmarshal(t *testing.T) {
	data := []byte(`<rules>
  <state name="root">
    <rule pattern="#.*$"><token type="CommentSingle"/></rule>
    <rule><include state="interp"/></rule>
  </state>
  <state name="interp">
    <rule pattern="\$\{"><token type="LiteralStringInterpol"/><push state="curly"/></rule>
  </state>
</rules>`)
	expected := Rules{
		"root": {
			{`#.*$`, CommentSingle, nil},
			Include("interp"),
		},
		"interp": {
			{`\$\{`, LiteralStringInterpol, Push("curly")},
		},
	}
	actual := Rules{}
	err := xml.Unmarshal(data, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
