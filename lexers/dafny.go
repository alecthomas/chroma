package lexers

import (
    . "github.com/alecthomas/chroma/v2" // nolint
)

var Dafny = Register(MustNewLexer(
    &Config{
        Name:            "Dafny",
        Aliases:         []string{"dafny"},
        Filenames:       []string{"*.dfy", "*.dafny"},
        MimeTypes:       []string{""},
        NotMultiline:    true,
        CaseInsensitive: false,
    },
    dafnyRules,
))

// All parenthesized numbers refer to the sections of the Dafny
// reference manual:
// https://dafny-lang.github.io/dafny/DafnyRef/DafnyRef.html
func dafnyRules() Rules {
    return Rules {
        "root": {
            // Tokens and whitespace (2.2)
			{`\n`, Text, nil},
			{`\s+`, Text, nil},
			{`\\\n`, Text, nil},
            {`//(.*?)\n`, CommentSingle, nil},
            {`/\*`, CommentMultiline, Push("multiline-comments")},

            // Reserved words (2.5.1)
            { Words(``, `\b`,
                `abstract`, `allocated`, `as`, `assert`, `assume`,
                `break`, `by`,
                `calc`, `case`,`class`, `codatatype`,
                `colemma`, `const`, `constructor`, `copredicate`,
                `datatype`, `decreases`,
                `else`, `ensures`, `exists`, `export`, `extends`,
                `false`, `forall`, `fresh`, `function`, `ghost`,
                `if`, `imap`, `import`, `in`, `include`, `inductive`,
                `int`, `invariant`, `is`, `iset`, `iterator`,
                `label`, `lemma`, `map`, `match`, `method`,
                `modifies`, `modify`, `module`, `multiset`,
                `nameonly`, `nat`, `new`, `newtype`, `null`,
                `object`, `object?`, `old`, `opened`, `ORDINAL`,
                `predicate`, `print`, `provides`,
                `reads`, `real`, `refines`, `requires`, `return`,
                `returns`, `reveal`, `reveals`,
                `seq`, `set`, `static`, `string`,
                `then`, `this`, `trait`, `true`, `twostate`, `type`,
                `unchanged`, `var`, `while`, `witness`,
                `yield`, `yields`,
            ), KeywordConstant, nil},

            // Character (2.5.5)
			{`'[^\n']'`, LiteralStringSingle, nil},
            // Strings (2.5.6)
            {`"`, LiteralStringDouble, Push("string")},

            // Numerics
            {`\d+i`, LiteralNumber, nil},
            {`0[xX][0-9a-fA-F_]+`, LiteralNumberHex, nil},

            // Constants
            {`(true|false)\b`, KeywordConstant, nil},

            // Operators
            {`(<===>|==>|<==|&&|\|\||==|!=|!)`, Operator, nil}, // Boolean operators (7.1)
            {`(<|<=|>=|>)`, Operator, nil}, // Numeric operators (7.2)
            {`(<<|>>|\+|-|\*|/|\||&|\||\^|%)`, Operator, nil}, // Bitvector operators (7.3)
            {`:=`, Operator, nil}, // Assignment

			{`[|^<>=!()\[\]{}.,;:]`, Punctuation, nil},
			{`[^\W\d]\w*`, NameOther, nil},
        },

        "multiline-comments": {
            {`/\*`, CommentMultiline, Push("multiline-comments")},
            {`\*/`, CommentMultiline, Pop(1)},
            {`[^/*]+`, CommentMultiline, nil},
            {`[/*]`, CommentMultiline, nil},
        },

        "string": {
            {`[^"]+`, LiteralStringDouble, nil},
            {`""`, LiteralStringDouble, nil},
            {`"`, LiteralStringDouble, Pop(1)},
        },
    }
}

