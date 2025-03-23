package chroma

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type AST struct {
	Config *Config    `parser:"( 'config' '{' @@ '}'"`
	States []stateAST `parser:"@@* )*"`
}

type stateAST struct {
	Name  string `parser:"'state' @Ident '{'"`
	Rules []Rule `parser:"@@* '}'"`
}

type tokenTypeAST struct {
	Type TokenType `parser:"@Ident"`
}

func (t tokenTypeAST) Emit(groups []string, state *LexerState) Iterator {
	return t.Type.Emit(groups, state)
}

var (
	lex = lexer.MustSimple([]lexer.SimpleRule{
		{"Punct", `[][={}(),:;]`},
		{"Whitespace", `\s+`},
		{"Comment", `//.*`},
		{"Boolean", `\b(true|false)\b`},
		{"Ident", `[a-zA-Z-][a-zA-Z0-9-]*`},
		{"String", `"(\\.|[^"])*"`},
		{"Float", `[-+]?\d*\.\d+([eE][-+]?\d+)?`},
		{"Int", `[-+]?\d+`},
		{"Regex", `/(\\.|[^/])+/`},
	})
	parser = participle.MustBuild[AST](
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.Map(func(token lexer.Token) (lexer.Token, error) {
			token.Value = token.Value[1 : len(token.Value)-1]
			return token, nil
		}, "Regex"),
		participle.Elide("Whitespace", "Comment"),
		participle.UseLookahead(1),
		participle.Union[Emitter](&byGroupsEmitter{}, &usingEmitter{}, &usingSelfEmitter{}, &tokenTypeAST{}),
		participle.Union[Mutator](&includeMutator{}, &combinedMutator{}, &pushMutator{}, &popMutator{}),
	)
)
