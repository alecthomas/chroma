# Lexer tests

This directory contains input source and expected output lexer tokens.

Input filenames for lexers are in the form `<name>.actual`. Expected output filenames are in the form `<name>.expected`.

Each input filename is parsed by the corresponding lexer and checked against the expected JSON-encoded token list.


To add/update tests do the following:

1. `export LEXER=csharp`
1. Create/edit a file `lexers/testdata/${LEXER}.actual` (eg. `csharp.actual`).
2. Run `go run ./cmd/chroma/main.go --lexer ${LEXER} --json lexers/testdata/${LEXER}.actual > lexers/testdata/${LEXER}.expected`.
3. Run `go test -v -run TestLexers ./lexers`.


eg.

```bash
$ export LEXER=csharp
$ go run ./cmd/chroma/main.go --lexer ${LEXER} --json lexers/testdata/${LEXER}.actual > lexers/testdata/${LEXER}.expected
$ go test -v -run TestLexers ./lexers
=== RUN   TestLexers
=== RUN   TestLexers/C#
=== RUN   TestLexers/CSS
--- PASS: TestLexers (0.01s)
    --- PASS: TestLexers/C# (0.00s)
    --- PASS: TestLexers/CSS (0.00s)
PASS
ok    github.com/alecthomas/chroma/lexers 0.032s
```

