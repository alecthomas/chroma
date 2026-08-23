package chroma

import (
	"compress/gzip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"sync"

	"github.com/dlclark/regexp2/v2"
)

// XML lexer rules use the following format:
//
//	<rules>
//	  <state name="$STATE">
//	    <rule [pattern="$PATTERN"]>
//	      [<$EMITTER ...>]
//	      [<$MUTATOR ...>]
//	    </rule>
//	  </state>
//	</rules>
//
// Include("String"):
//
//	<rule>
//	  <include state="String" />
//	</rule>
//
// Rule{`\d+`, Text, nil}:
//
//	<rule pattern="\\d+">
//	  <token type="Text"/>
//	</rule>
//
// Rule{`"`, String, Push("String")}:
//
//	<rule pattern="\"">
//	  <token type="String" />
//	  <push state="String" />
//	</rule>
//
// Rule{`(\w+)(\n)`, ByGroups(Keyword, Whitespace), nil}:
//
//	<rule pattern="(\\w+)(\\n)">
//	  <bygroups token="Keyword" token="Whitespace" />
//	  <push state="String" />
//	</rule>
var (
	errUnknownXMLRuleElement = errors.New("unknown XML rule element")
	emitterDecoders          = map[string]xmlDecoder[Emitter]{
		"bygroups":     decodeXML(func(value *byGroupsEmitter) Emitter { return value }),
		"usingself":    decodeXML(func(value *usingSelfEmitter) Emitter { return value }),
		"token":        decodeXML(func(value *TokenType) Emitter { return *value }),
		"using":        decodeXML(func(value *usingEmitter) Emitter { return value }),
		"usingbygroup": decodeXML(func(value *usingByGroup) Emitter { return value }),
	}
	mutatorDecoders = map[string]xmlDecoder[Mutator]{
		"include":  decodeXML(func(value *includeMutator) Mutator { return value }),
		"combined": decodeXML(func(value *combinedMutator) Mutator { return value }),
		"mutators": decodeXML(func(value *multiMutator) Mutator { return value }),
		"push":     decodeXML(func(value *pushMutator) Mutator { return value }),
		"pop":      decodeXML(func(value *popMutator) Mutator { return value }),
	}
)

type xmlDecoder[T any] func(*xml.Decoder, xml.StartElement) (T, error)

func decodeXML[T, U any](convert func(*T) U) xmlDecoder[U] {
	return func(d *xml.Decoder, start xml.StartElement) (U, error) {
		value := new(T)
		if err := d.DecodeElement(value, &start); err != nil {
			var zero U
			return zero, err
		}
		return convert(value), nil
	}
}

// fastUnmarshalConfig unmarshals only the Config from an XML lexer definition.
func fastUnmarshalConfig(from fs.FS, path string) (*Config, error) {
	r, err := from.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	dec := xml.NewDecoder(r)
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil, fmt.Errorf("could not find <config> element")
			}
			return nil, err
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local != "config" {
				break
			}

			var config Config
			err = dec.DecodeElement(&config, &se)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", path, err)
			}
			return &config, nil
		}
	}
}

// MustNewXMLLexer constructs a new RegexLexer from an XML file or panics.
func MustNewXMLLexer(from fs.FS, path string) *RegexLexer {
	lex, err := NewXMLLexer(from, path)
	if err != nil {
		panic(err)
	}
	return lex
}

// NewXMLLexer creates a new RegexLexer from an XML lexer definition.
func NewXMLLexer(from fs.FS, path string) (*RegexLexer, error) {
	config, err := fastUnmarshalConfig(from, path)
	if err != nil {
		return nil, err
	}
	return NewXMLLexerFromConfig(config, from, path)
}

// NewXMLLexerFromConfig creates a RegexLexer whose Config is already known,
// loading its rules lazily from the XML definition at path in from.
//
// The analyser regexes in config are compiled lazily on the first call to
// AnalyseText; an invalid analyser regex results in a zero score at analyse
// time rather than an error here.
func NewXMLLexerFromConfig(config *Config, from fs.FS, path string) (*RegexLexer, error) {
	if err := config.validateGlobs(); err != nil {
		return nil, err
	}

	var analyserFn func(string) float32

	if config.Analyse != nil {
		type regexAnalyse struct {
			re    *regexp2.Regexp
			score float32
		}

		compileAnalysers := sync.OnceValue(func() []regexAnalyse {
			regexAnalysers := make([]regexAnalyse, 0, len(config.Analyse.Regexes))

			regexFlags := regexp2.None
			if config.CaseInsensitive {
				regexFlags = regexp2.IgnoreCase
			}
			for _, ra := range config.Analyse.Regexes {
				re, err := regexp2.Compile(ra.Pattern, regexFlags)
				if err != nil {
					return nil
				}

				regexAnalysers = append(regexAnalysers, regexAnalyse{re, ra.Score})
			}
			return regexAnalysers
		})

		analyserFn = func(text string) float32 {
			var score float32

			for _, ra := range compileAnalysers() {
				ok, err := ra.re.MatchString(text)
				if err != nil {
					return 0
				}

				if ok && config.Analyse.First {
					return min(ra.score, 1.0)
				}

				if ok {
					score += ra.score
				}
			}

			return min(score, 1.0)
		}
	}

	return &RegexLexer{
		config:   config,
		analyser: analyserFn,
		fetchRulesFunc: func() (Rules, error) {
			var lexer struct {
				Config
				Rules Rules `xml:"rules"`
			}
			// Try to open .xml fallback to .xml.gz
			fr, err := from.Open(path)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					path += ".gz"
					fr, err = from.Open(path)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			}
			defer fr.Close()
			var r io.Reader = fr
			if strings.HasSuffix(path, ".gz") {
				r, err = gzip.NewReader(r)
				if err != nil {
					return nil, fmt.Errorf("%s: %w", path, err)
				}
			}
			err = xml.NewDecoder(r).Decode(&lexer)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", path, err)
			}
			return lexer.Rules, nil
		},
	}, nil
}

func unmarshalMutator(d *xml.Decoder, start xml.StartElement) (Mutator, error) {
	kind := start.Name.Local
	decode, ok := mutatorDecoders[kind]
	if !ok {
		return nil, fmt.Errorf("unknown mutator %q: %w", kind, errUnknownXMLRuleElement)
	}
	return decode(d, start)
}

func unmarshalEmitter(d *xml.Decoder, start xml.StartElement) (Emitter, error) {
	kind := start.Name.Local
	decode, ok := emitterDecoders[kind]
	if !ok {
		return nil, fmt.Errorf("unknown emitter %q: %w", kind, errUnknownXMLRuleElement)
	}
	return decode(d, start)
}

func (r *Rule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "pattern" {
			r.Pattern = attr.Value
			break
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch token := token.(type) {
		case xml.StartElement:
			mutator, err := unmarshalMutator(d, token)
			if err != nil && !errors.Is(err, errUnknownXMLRuleElement) {
				return err
			} else if err == nil {
				if r.Mutator != nil {
					return fmt.Errorf("duplicate mutator")
				}
				r.Mutator = mutator
				continue
			}
			emitter, err := unmarshalEmitter(d, token)
			if err != nil && !errors.Is(err, errUnknownXMLRuleElement) { // nolint: gocritic
				return err
			} else if err == nil {
				if r.Type != nil {
					return fmt.Errorf("duplicate emitter")
				}
				r.Type = emitter
				continue
			} else {
				return err
			}

		case xml.EndElement:
			return nil
		}
	}
}

type xmlRuleState struct {
	Name  string `xml:"name,attr"`
	Rules []Rule `xml:"rule"`
}

type xmlRules struct {
	States []xmlRuleState `xml:"state"`
}

func (r *Rules) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	xr := xmlRules{}
	if err := d.DecodeElement(&xr, &start); err != nil {
		return err
	}
	if *r == nil {
		*r = Rules{}
	}
	for _, state := range xr.States {
		(*r)[state.Name] = state.Rules
	}
	return nil
}

type xmlTokenType struct {
	Type string `xml:"type,attr"`
}

func (t *TokenType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	el := xmlTokenType{}
	if err := d.DecodeElement(&el, &start); err != nil {
		return err
	}
	tt, err := TokenTypeString(el.Type)
	if err != nil {
		return err
	}
	*t = tt
	return nil
}

func (b *Emitters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch token := token.(type) {
		case xml.StartElement:
			emitter, err := unmarshalEmitter(d, token)
			if err != nil {
				return err
			}
			*b = append(*b, emitter)

		case xml.EndElement:
			return nil
		}
	}
}
