package y

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var YANG = internal.Register(MustNewLexer(
	&Config{
		Name:      "YANG",
		Aliases:   []string{"yang"},
		Filenames: []string{"*.yang"},
		MimeTypes: []string{"application/yang"},
	},
	Rules{
		"root": {
			{`\s+`, Whitespace, nil},
			{`[\{\}\;]+`, Punctuation, nil},
			{`(?<![\-\w])(and|or|not|\+|\.)(?![\-\w])`, Operator, nil},

			{`"(?:\\"|[^"])*?"`, StringDouble, nil},
			{`'(?:\\'|[^'])*?'`, StringSingle, nil},

			{`/\*`, CommentMultiline, Push("comments")},
			{`//.*?$`, CommentSingle, nil},

			//match BNF stmt for `node-identifier` with [ prefix ":"]
			{`(?:^|(?<=[\s{};]))([\w.-]+)(:)([\w.-]+)(?=[\s{};])`, ByGroups(KeywordNamespace, Punctuation, Text), nil},

			//match BNF stmt `date-arg-str`
			{`([0-9]{4}\-[0-9]{2}\-[0-9]{2})(?=[\s\{\}\;])`, LiteralDate, nil},
			{`([0-9]+\.[0-9]+)(?=[\s\{\}\;])`, NumberFloat, nil},
			{`([0-9]+)(?=[\s\{\}\;])`, NumberInteger, nil},

			//TOP_STMTS_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `module`, `submodule`), Keyword, nil},
			//MODULE_HEADER_STMT_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `yang-version`, `namespace`, `prefix`, `belongs-to`), Keyword, nil},
			//META_STMT_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `organization`, `contact`, `description`, `reference`, `revision`), Keyword, nil},
			//LINKAGE_STMTS_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `import`, `include`, `revision-date`), Keyword, nil},
			//BODY_STMT_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `extension`, `feature`, `identity`, `typedef`, `grouping`, `augment`, `rpc`, `notification`, `deviation`, `action`, `argument`, `if-feature`, `input`, `output`), Keyword, nil},
			//DATA_DEF_STMT_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `container`, `leaf-list`, `leaf`, `list`, `choice`, `anydata`, `anyxml`, `uses`, `case`, `config`, `deviate`, `must`, `when`, `presence`, `refine`), Keyword, nil},
			//TYPE_STMT_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `type`, `units`, `default`, `status`, `bit`, `enum`, `error-app-tag`, `error-message`, `fraction-digits`, `length`, `min-elements`, `max-elements`, `modifier`, `ordered-by`, `path`, `pattern`, `position`, `range`, `require-instance`, `value`, `yin-element`, `base`), Keyword, nil},
			//LIST_STMT_KEYWORDS
			{Words(``, `(?=[^\w\-\:])`, `key`, `mandatory`, `unique`), Keyword, nil},

			//CONSTANTS_KEYWORDS - RFC7950 other keywords
			{Words(``, `(?=[^\w\-\:])`, `true`, `false`, `current`, `obsolete`, `deprecated`, `add`, `delete`, `replace`, `not-supported`, `invert-match`, `max`, `min`, `unbounded`, `user`), NameClass, nil},

			//RFC7950 Built-In Types
			{Words(``, `(?=[^\w\-\:])`, `binary`, `bits`, `boolean`, `decimal64`, `empty`, `enumeration`, `int8`, `int16`, `int32`, `int64`, `string`, `uint8`, `uint16`, `uint32`, `uint64`, `union`, `leafref`, `identityref`, `instance-identifier`), NameClass, nil},


			{`[^;{}\s\'\"]+`, Text ,nil},
		},
		"comments": {
			{`[^*/]`, CommentMultiline, nil},
			{`/\*`, CommentMultiline, Push("comment")},
			{`\*/`, CommentMultiline, Pop(1)},
			{`[*/]`, CommentMultiline, nil},
		},
	},
))
