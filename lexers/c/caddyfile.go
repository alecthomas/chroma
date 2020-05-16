package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CaddyfileCommon are the rules common to both of the lexer variants
var CaddyfileCommon = Rules{
	"site_block_common": {
		// Import keyword
		{`(import)(\s+)([^\s]+)`, ByGroups(Keyword, Text, NameVariableMagic), nil},
		// Matcher definition
		{`@[^\s]+\s+`, NameDecorator, Push("matcher")},
		// These are special, they can nest more directives
		{`handle|route|handle_path|not`, Keyword, Push("nested_directive")},
		// Any other directive
		{`[^\s#]+`, Keyword, Push("directive")},
		Include("base"),
	},
	"matcher": {
		{`\{`, Punctuation, Push("block")},
		// Not can be one-liner
		{`not`, Keyword, Push("deep_not_matcher")},
		// Any other same-line matcher
		{`[^\s#]+`, Keyword, Push("arguments")},
		// Terminators
		{`\n`, Text, Pop(1)},
		{`\}`, Punctuation, Pop(1)},
		Include("base"),
	},
	"block": {
		{`\}`, Punctuation, Pop(2)},
		// Not can be one-liner
		{`not`, Keyword, Push("not_matcher")},
		// Any other subdirective
		{`[^\s#]+`, Keyword, Push("subdirective")},
		Include("base"),
	},
	"nested_block": {
		{`\}`, Punctuation, Pop(2)},
		// Matcher definition
		{`@[^\s]+\s+`, NameDecorator, Push("matcher")},
		// Any other directive
		{`[^\s#]+`, Keyword, Push("nested_directive")},
		Include("base"),
	},
	"not_matcher": {
		{`\}`, Punctuation, Pop(2)},
		{`\{(?=\s)`, Punctuation, Push("block")},
		{`[^\s#]+`, Keyword, Push("arguments")},
	},
	"deep_not_matcher": {
		{`\}`, Punctuation, Pop(2)},
		{`\{(?=\s)`, Punctuation, Push("block")},
		{`[^\s#]+`, Keyword, Push("deep_subdirective")},
	},
	"directive": {
		{`\{(?=\s)`, Punctuation, Push("block")},
		Include("matcher_token"),
		Include("comments_pop"),
		{`\n`, Text, Pop(1)},
		Include("base"),
	},
	"nested_directive": {
		{`\{(?=\s)`, Punctuation, Push("nested_block")},
		Include("matcher_token"),
		Include("comments_pop"),
		{`\n`, Text, Pop(1)},
		Include("base"),
	},
	"subdirective": {
		{`\{(?=\s)`, Punctuation, Push("block")},
		Include("comments_pop"),
		{`\n`, Text, Pop(1)},
		Include("base"),
	},
	"arguments": {
		{`\{(?=\s)`, Punctuation, Push("block")},
		Include("comments_pop"),
		{`\n`, Text, Pop(2)},
		Include("base"),
	},
	"deep_subdirective": {
		{`\{(?=\s)`, Punctuation, Push("block")},
		Include("comments_pop"),
		{`\n`, Text, Pop(3)},
		Include("base"),
	},
	"matcher_token": {
		{`@[^\s]+`, NameDecorator, Push("arguments")},         // Named matcher
		{`/[^\s]+`, NameDecorator, Push("arguments")},         // Path matcher
		{`\*`, NameDecorator, Push("arguments")},              // Wildcard path matcher
		{`\[\<matcher\>\]`, NameDecorator, Push("arguments")}, // Matcher token stub for docs
	},
	"comments": {
		{`^#.*\n`, CommentSingle, nil},   // Comment at start of line
		{`\s+#.*\n`, CommentSingle, nil}, // Comment preceded by whitespace
	},
	"comments_pop": {
		{`^#.*\n`, CommentSingle, Pop(1)},   // Comment at start of line
		{`\s+#.*\n`, CommentSingle, Pop(1)}, // Comment preceded by whitespace
	},
	"base": {
		Include("comments"),
		{`on|off`, NameConstant, nil},
		{`(https?://)?([a-z0-9.-]+)(:)([0-9]+)`, ByGroups(Name, Name, Punctuation, LiteralNumberInteger), nil},
		{`[a-z-]+/[a-z-+]+`, LiteralString, nil},
		{`[0-9]+[km]?\b`, LiteralNumberInteger, nil},
		{`\{[\w+.-]+\}`, NameAttribute, nil}, // Placeholder
		{`[^\s#{}$]+`, LiteralString, nil},
		{`/[^\s#]*`, Name, nil},
		{`\s+`, Text, nil},
	},
}

// CaddyfileRules are the merged rules for the main Caddyfile lexer
var CaddyfileRules = (func(a Rules, b Rules) Rules {
	for k, v := range b {
		a[k] = v
	}
	return a
})(
	Rules{
		"root": {
			Include("comments"),
			// Global options block
			{`^\s*(\{)\s*$`, ByGroups(Punctuation), Push("globals")},
			// Snippets
			{`(\([^\s#]+\))(\s*)(\{)`, ByGroups(NameVariableAnonymous, Text, Punctuation), Push("snippet")},
			// Site label
			{`[^#{(\s,]+`, NameLabel, Push("label")},
			// Site label with placeholder
			{`\{[\w+.-]+\}`, NameAttribute, Push("label")},
		},
		"globals": {
			{`\}`, Punctuation, Pop(1)},
			{`[^\s#]+`, KeywordType, Push("directive")},
			Include("base"),
		},
		"snippet": {
			{`\}`, Punctuation, Pop(1)},
			// Matcher definition
			{`@[^\s]+\s+`, NameDecorator, Push("matcher")},
			// Any directive
			{`[^\s#]+`, KeywordType, Push("directive")},
			Include("base"),
		},
		"label": {
			{`,`, Text, nil},
			{` `, Text, nil},
			{`[^#{(\s,]+`, NameLabel, nil},
			// Comment after non-block label (hack because comments end in \n)
			{`#.*\n`, CommentSingle, Push("site_block")},
			// Note: if \n, we'll never pop out of the site_block, it's valid
			{`\{(?=\s)|\n`, Punctuation, Push("site_block")},
		},
		"site_block": {
			{`\}`, Punctuation, Pop(2)},
			Include("site_block_common"),
		},
	},
	CaddyfileCommon,
)

// CaddyfileDirectiveRules are the merged rules for the secondary lexer
var CaddyfileDirectiveRules = (func(a Rules, b Rules) Rules {
	for k, v := range b {
		a[k] = v
	}
	return a
})(
	Rules{
		// Same as "site_block" in Caddyfile
		"root": {
			Include("site_block_common"),
		},
	},
	CaddyfileCommon,
)

// Caddyfile lexer.
var Caddyfile = internal.Register(MustNewLexer(
	&Config{
		Name:      "Caddyfile",
		Aliases:   []string{"caddyfile", "caddy"},
		Filenames: []string{"Caddyfile*"},
		MimeTypes: []string{},
	},
	CaddyfileRules,
))

// Caddyfile directive-only lexer.
var CaddyfileDirectives = internal.Register(MustNewLexer(
	&Config{
		Name:      "Caddyfile Directives",
		Aliases:   []string{"caddyfile-directives", "caddyfile-d", "caddy-d"},
		Filenames: []string{},
		MimeTypes: []string{},
	},
	CaddyfileDirectiveRules,
))
