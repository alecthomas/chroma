# Lexer tests

This directory contains input source and expected output lexer tokens.

Input filenames for lexers are in the form `<name>.<name>`. Expected output filenames are in the form `<name>.expected`.

Each input filename is parsed by the corresponding lexer and checked against the expected JSON-encoded token list.


To add/update tests do the following:

1. `export LEXER=css`
1. Create/edit a file `lexers/testdata/${LEXER}.${LEXER}` (eg. `css.css`).
2. Run `go run ./cmd/chroma/main.go --lexer ${LEXER} --json lexers/testdata/${LEXER}.${LEXER} > lexers/testdata/${LEXER}.expected`.
3. Run `go test -v ./lexers`.


eg.

```bash
$ export LEXER=css
$ go run ./cmd/chroma/main.go --lexer ${LEXER} --json lexers/testdata/${LEXER}.${LEXER} > lexers/testdata/${LEXER}.expected
$ cat lexers/testdata/${LEXER}.expected
[
  {"type":"Punctuation","value":":"},
  {"type":"NameDecorator","value":"root"},
  {"type":"Text","value":" "},
  {"type":"Punctuation","value":"{"},
  {"type":"Text","value":"\n  "},
  {"type":"NameVariable","value":"--variable-name"},
  {"type":"Text","value":""},
  {"type":"Punctuation","value":":"},
  {"type":"Text","value":" "},
  {"type":"LiteralNumberHex","value":"#fff"},
  {"type":"Punctuation","value":";"},
  {"type":"Text","value":"\n"},
  {"type":"Punctuation","value":"}"},
  {"type":"Text","value":"\n"}
]
$ go test -v ./lexers
=== RUN   TestDiffLexerWithoutTralingNewLine
--- PASS: TestDiffLexerWithoutTralingNewLine (0.00s)
=== RUN   TestLexers
=== RUN   TestLexers/CSS
--- PASS: TestLexers (0.00s)
    --- PASS: TestLexers/CSS (0.00s)
=== RUN   TestCompileAllRegexes
--- PASS: TestCompileAllRegexes (0.61s)
=== RUN   TestGet
=== RUN   TestGet/ByName
=== RUN   TestGet/ByAlias
=== RUN   TestGet/ViaFilename
--- PASS: TestGet (0.00s)
    --- PASS: TestGet/ByName (0.00s)
    --- PASS: TestGet/ByAlias (0.00s)
    --- PASS: TestGet/ViaFilename (0.00s)
PASS
ok    github.com/alecthomas/chroma/lexers 0.649s
```

