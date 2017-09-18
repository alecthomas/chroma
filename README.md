# Chroma - A general purpose syntax highlighter for Go [![](https://godoc.org/github.com/alecthomas/chroma?status.svg)](http://godoc.org/github.com/alecthomas/chroma) [![Build Status](https://travis-ci.org/alecthomas/chroma.png)](https://travis-ci.org/alecthomas/chroma) [![Gitter chat](https://badges.gitter.im/alecthomas.png)](https://gitter.im/alecthomas/Lobby)

Chroma takes source code and other structured text and converts it into syntax
highlighted HTML, ANSI-coloured text, etc.

Chroma is based heavily on [Pygments](http://pygments.org/), and includes
translaters for Pygments lexers and styles.

## Table of Contents

<!-- MarkdownTOC -->

- [Supported languages](#supported-languages)
- [Using the library](#using-the-library)
  - [Quick start](#quick-start)
  - [Identifying the language](#identifying-the-language)
  - [Formatting the output](#formatting-the-output)
  - [The HTML formatter](#the-html-formatter)
- [More detail](#more-detail)
  - [Lexers](#lexers)
  - [Formatters](#formatters)
  - [Styles](#styles)
- [Command-line interface](#command-line-interface)
- [What's missing compared to Pygments?](#whats-missing-compared-to-pygments)

<!-- /MarkdownTOC -->

## Supported languages

ABNF, ActionScript, ActionScript 3, Ada, Angular2, ANTLR, ApacheConf, APL,
AppleScript, Awk, Bash, Batchfile, BlitzBasic, BNF, Brainfuck, C, Cap'n
Proto, Ceylon, CFEngine3, ChaiScript, Python, Cheetah, Common Lisp, Clojure,
COBOL, CoffeeScript, cfstatement, Coq, C++, Crystal, CSS, Cython, Dart, Diff,
Django/Jinja, Docker, DTD, EBNF, Elixir, Elm, EmacsLisp, Erlang, Factor, Fish,
Forth, Fortran, FSharp, GAS, Genshi Text, GLSL, Gnuplot, Go, Groovy,
Handlebars, Haskell, Haxe, JavaScript, HTML, Hy, Idris, INI, Io, Java, JSON,
Julia, Kotlin, Lighttpd configuration file, LLVM, Lua, Base Makefile, Mako,
markdown, Perl, Mason, Mathematica, Modula-2, Myghty, MySQL, NASM, Newspeak,
Nginx configuration file, Nimrod, OCaml, Octave, PacmanConf, PHP, Pig,
PkgConfig, PL/pgSQL, PostgreSQL SQL dialect, PostScript, POVRay, PowerShell,
Prolog, Protocol Buffer, Puppet, Python 3, QBasic, Racket, Ragel, reg, Rexx,
Ruby, Rust, Sass, Scala, Scheme, Scilab, Smalltalk, Smarty, Snobol, SPARQL,
SQL, SquidConf, Swift, TASM, Tcl, Tcsh, Termcap, Terminfo, Terraform, Thrift,
Transact-SQL, Turtle, Twig, TypeScript, TypoScript, verilog, vhdl, VimL, XML,
Xorg

_I will attempt to keep this section up to date, but an authoritative list can be
displayed with `chroma --list`._

## Using the library

Chroma, like Pygments, has the concepts of
[lexers](https://github.com/alecthomas/chroma/tree/master/lexers),
[formatters](https://github.com/alecthomas/chroma/tree/master/formatters) and
[styles](https://github.com/alecthomas/chroma/tree/master/styles).

Lexers convert source text into a stream of tokens, styles specify how token
types are mapped to colours, and formatters convert tokens and styles into
formatted output.

A package exists for each of these, with a global `Registry` variable
containing all of the registered implementations. There are also helper
functions for using the registry in each package, such as looking up lexers by
name or matching filenames, etc.

In all cases, if a lexer, formatter or style can not be determined, `nil` will
be returned. In this situation you may want to default to the `Fallback`
value in each respective package, which provides sane defaults.

### Quick start

A convenience function exists that can be used to simply format some source
text, without any effort:

```go
err := quick.Highlight(os.Stdout, someSourceCode, "go", "html", "monokai")
```

### Identifying the language

To highlight code, you'll first have to identify what language the code is
written in. There are three primary ways to do that:

1. Detect the language from its filename.

    ```go
    lexer := lexers.Match("foo.go")
    ```

3. Explicitly specify the language by its Chroma syntax ID (a full list is available from `lexers.Names()`).

    ```go
    lexer := lexers.Get("go")
    ```

3. Detect the language from its content.

    ```go
    lexer := lexers.Analyse("package main\n\nfunc main()\n{\n}\n")
    ```

In all cases, `nil` will be returned if the langauge can not be identified.

```go
if lexer == nil {
  lexer = lexers.Fallback
}
```

At this point, it should be noted that some lexers can be extremely chatty. To
mitigate this, you can use the coalescing lexer to coalesce runs of identical
token types into a single token:

```go
lexer = chroma.Coalesce(lexer)
```

### Formatting the output

Once a language is identified you will need to pick a formatter and a style (theme).

```go
style := styles.Get("swapoff")
if style == nil {
  style = styles.Fallback
}
formatter := formatters.Get("html")
if formatter == nil {
  formatter = formatters.Fallback
}
```

Then obtain a formatting function from the formatter:

```go
writer, err := formatter.Format(w, style)
```

And finally, lex the source code and write the output:

```go
contents, err := ioutil.ReadAll(r)
err := lexer.Tokenise(nil, string(contents), writer)
```

### The HTML formatter

By default the `html` registered formatter generates standalone HTML with
embedded CSS. More flexibility is available through the `lexers/html` package.

Firstly, the output generated by the formatter can be customised with the
following constructor options:

- `Standalone()` - generate standalone HTML with embedded CSS.
- `WithClasses()` - use classes rather than inlined style attributes.
- `ClassPrefix(prefix)` - prefix each generated CSS class.

If `WithClasses()` is used, the corresponding CSS can be obtained from the formatter with:

```go
formatter := html.New(html.WithClasses())
err := formatter.WriteCSS(w, style)
```

## More detail

### Lexers

See the [Pygments documentation](http://pygments.org/docs/lexerdevelopment/)
for details on implementing lexers. Most concepts apply directly to Chroma,
but see existing lexer implementations for real examples.

In many cases lexers can be automatically converted directly from Pygments by
using the included Python 3 script `pygments2chroma.py`. I use something like
the following:

```
python3 ~/Projects/chroma/_tools/pygments2chroma.py \
  pygments.lexers.jvm.KotlinLexer \
  > ~/Projects/chroma/lexers/kotlin.go \
  && gofmt -s -w ~/Projects/chroma/lexers/*.go
```

See notes in [pygments-lexers.go](https://github.com/alecthomas/chroma/blob/master/pygments-lexers.txt)
for a list of lexers, and notes on some of the issues importing them.

### Formatters

Chroma supports HTML output, as well as terminal output in 8 colour, 256 colour, and true-colour.

A `noop` formatter is included that outputs the token text only, and a `raw`
formatter outputs raw token structs. The latter is useful for debugging lexers.

### Styles

Chroma styles use the [same syntax](http://pygments.org/docs/styles/) as Pygments.

All Pygments styles have been converted to Chroma using the `_tools/style.py` script.

## Command-line interface

A command-line interface to Chroma is included. It can be installed with:

```
go get -u github.com/alecthomas/chroma/cmd/chroma
```

## What's missing compared to Pygments?

- Quite a few lexers, for various reasons (pull-requests welcome):
    - Pygments lexers for complex languages often include custom code to
      handle certain aspects, such as Perl6's ability to nest code inside
      regular expressions. These require time and effort to convert.
    - I mostly only converted languages I had heard of, to reduce the porting cost.
- Some more esoteric features of Pygments are omitted for simplicity.
- Though the Chroma API supports content detection, very few languages support them.
  I have plans to implement a statistical analyser at some point, but not enough time.
