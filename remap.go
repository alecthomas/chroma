package chroma

type remappingLexer struct {
	lexer  Lexer
	mapper func(*Token) []*Token
}

// RemappingLexer remaps a token to a set of, potentially empty, tokens.
func RemappingLexer(lexer Lexer, mapper func(*Token) []*Token) Lexer {
	return &remappingLexer{lexer, mapper}
}

func (r *remappingLexer) Config() *Config {
	return r.lexer.Config()
}

func (r *remappingLexer) Tokenise(options *TokeniseOptions, text string) (Iterator, error) {
	it, err := r.lexer.Tokenise(options, text)
	if err != nil {
		return nil, err
	}
	buffer := []*Token{}
	return func() *Token {
		for {
			if len(buffer) > 0 {
				t := buffer[0]
				buffer = buffer[1:]
				return t
			}
			t := it()
			if t == nil {
				return t
			}
			buffer = r.mapper(t)
		}
	}, nil
}

type TypeMapping struct {
	From TokenType
	To   TokenType
}
type TypeRemappingMap map[TypeMapping][]string

// TypeRemappingLexer remaps types of tokens coming from a parent Lexer.
//
// eg. Map "defvaralias" tokens of type NameVariable to NameFunction:
//
// 		mapping := TypeRemappingMap{
// 			{NameVariable, NameFunction}: {"defvaralias"},
// 		}
// 		lexer = TypeRemappingLexer(lexer, mapping)
func TypeRemappingLexer(lexer Lexer, mapping TypeRemappingMap) Lexer {
	// Lookup table for fast remapping.
	lut := map[TokenType]map[string]TokenType{}
	for rt, kl := range mapping {
		km, ok := lut[rt.From]
		if !ok {
			km = map[string]TokenType{}
			lut[rt.From] = km
		}
		for _, k := range kl {
			km[k] = rt.To
		}

	}
	return RemappingLexer(lexer, func(t *Token) []*Token {
		if k, ok := lut[t.Type]; ok {
			if tt, ok := k[t.Value]; ok {
				t = t.Clone()
				t.Type = tt
			}
		}
		return []*Token{t}
	})
}
