package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var (
	// Populated by goreleaser.
	version = "?"
	commit  = "?"
	date    = "?"

	description = `
Chroma is a general purpose syntax highlighting library and corresponding
command, for Go.
`

	cli struct {
		Version    kong.VersionFlag `help:"Show version."`
		Profile    string           `hidden:"" help:"Enable profiling to file."`
		List       bool             `help:"List lexers, styles and formatters."`
		Unbuffered bool             `help:"Do not buffer output."`
		Trace      bool             `help:"Trace lexer states as they are traversed."`
		Check      bool             `help:"Do not format, check for tokenisation errors instead."`
		Filename   string           `help:"Filename to use for selecting a lexer when reading from stdin."`
		Fail       bool             `help:"Exit silently with status 1 if no specific lexer was found."`

		Lexer     string `help:"Lexer to use when formatting." default:"autodetect" short:"l" enum:"${lexers}"`
		Style     string `help:"Style to use for formatting." default:"swapoff" short:"s" enum:"${styles}"`
		Formatter string `help:"Formatter to use." default:"terminal" short:"f" enum:"${formatters}"`

		JSON bool `help:"Output JSON representation of tokens."`

		HTML                      bool   `help:"Enable HTML mode (equivalent to '--formatter html')."`
		HTMLPrefix                string `help:"HTML CSS class prefix." placeholder:"PREFIX"`
		HTMLStyles                bool   `help:"Output HTML CSS styles."`
		HTMLAllStyles             bool   `help:"Output all HTML CSS styles, including redundant ones."`
		HTMLOnly                  bool   `help:"Output HTML fragment."`
		HTMLInlineStyles          bool   `help:"Output HTML with inline styles (no classes)."`
		HTMLTabWidth              int    `help:"Set the HTML tab width." default:"8"`
		HTMLLines                 bool   `help:"Include line numbers in output."`
		HTMLLinesTable            bool   `help:"Split line numbers and code in a HTML table"`
		HTMLLinesStyle            string `help:"Style for line numbers."`
		HTMLHighlight             string `help:"Highlight these lines." placeholder:"N[:M][,...]"`
		HTMLHighlightStyle        string `help:"Style used for highlighting lines."`
		HTMLBaseLine              int    `help:"Base line number." default:"1"`
		HTMLPreventSurroundingPre bool   `help:"Prevent the surrounding pre tag."`

		SVG bool `help:"Output SVG representation of tokens."`

		Files []string `arg:"" optional:"" help:"Files to highlight." type:"existingfile"`
	}
)

type flushableWriter interface {
	io.Writer
	Flush() error
}

type nopFlushableWriter struct{ io.Writer }

func (n *nopFlushableWriter) Flush() error { return nil }

// prepareLenient prepares contents and lexer for input, using fallback lexer if no specific one is available for it.
func prepareLenient(ctx *kong.Context, r io.Reader, filename string) (string, chroma.Lexer) {
	data, err := ioutil.ReadAll(r)
	ctx.FatalIfErrorf(err)

	contents := string(data)
	lexer := selexer(filename, contents)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	return contents, lexer
}

// prepareSpecific prepares contents and lexer for input, exiting if there is no specific lexer available for it.
// Input is consumed only up to peekSize for lexer selection. With fullSize -1, consume r until EOF.
func prepareSpecific(ctx *kong.Context, r io.Reader, filename string, peekSize, fullSize int) (string, chroma.Lexer) {
	data := make([]byte, peekSize)
	n, err := io.ReadFull(r, data)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		ctx.FatalIfErrorf(err)
	}

	lexer := selexer(filename, string(data[:n]))
	if lexer == nil {
		ctx.Exit(1)
	}

	if n < peekSize {
		return string(data[:n]), lexer
	}
	var ndata []byte
	if fullSize == -1 {
		rest, err := io.ReadAll(r)
		ctx.FatalIfErrorf(err)
		ndata = make([]byte, n + len(rest))
		copy(ndata, data[:n])
		copy(ndata[n:], rest)
	} else {
		ndata = make([]byte, fullSize)
		copy(ndata, data[:n])
		_, err = io.ReadFull(r, ndata[n:])
		ctx.FatalIfErrorf(err)
	}

	return string(ndata), lexer
}

func main() {
	ctx := kong.Parse(&cli, kong.Description(description), kong.Vars{
		"version":    fmt.Sprintf("%s-%s-%s", version, commit, date),
		"lexers":     "autodetect," + strings.Join(lexers.Names(true), ","),
		"styles":     strings.Join(styles.Names(), ","),
		"formatters": strings.Join(formatters.Names(), ","),
	})
	if cli.List {
		listAll()
		return
	}
	if cli.Profile != "" {
		f, err := os.Create(cli.Profile)
		ctx.FatalIfErrorf(err)
		err = pprof.StartCPUProfile(f)
		ctx.FatalIfErrorf(err)
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		go func() {
			<-signals
			pprof.StopCPUProfile()
			os.Exit(128 + 3)
		}()
		defer pprof.StopCPUProfile()
	}
	if path.Base(os.Args[0]) == ".lessfilter" {
		// https://manpages.debian.org/lesspipe#USER_DEFINED_FILTERS
		cli.Fail = true
	}

	var out io.Writer = os.Stdout
	if runtime.GOOS == "windows" && isatty.IsTerminal(os.Stdout.Fd()) {
		out = colorable.NewColorableStdout()
	}
	var w flushableWriter
	if cli.Unbuffered {
		w = &nopFlushableWriter{out}
	} else {
		w = bufio.NewWriterSize(out, 16384)
	}
	defer w.Flush() // nolint: errcheck

	switch {
	case cli.JSON:
		cli.Formatter = "json"
	case cli.HTML:
		cli.Formatter = "html"
	case cli.SVG:
		cli.Formatter = "svg"
	}

	// Retrieve user-specified style, clone it, and add some overrides.
	builder := styles.Get(cli.Style).Builder()
	if cli.HTMLHighlightStyle != "" {
		builder.Add(chroma.LineHighlight, cli.HTMLHighlightStyle)
	}
	if cli.HTMLLinesStyle != "" {
		builder.Add(chroma.LineNumbers, cli.HTMLLinesStyle)
	}
	style, err := builder.Build()
	ctx.FatalIfErrorf(err)

	// Dump styles.
	if cli.HTMLStyles {
		options := []html.Option{html.WithClasses(true)}
		if cli.HTMLAllStyles {
			options = append(options, html.WithAllClasses(true))
		}
		formatter := html.New(options...)
		err = formatter.WriteCSS(w, style)
		ctx.FatalIfErrorf(err)
		return
	}

	if cli.Formatter == "html" {
		configureHTMLFormatter(ctx)
	}
	if len(cli.Files) == 0 {
		var contents string
		var lexer chroma.Lexer
		if cli.Fail {
			contents, lexer = prepareSpecific(ctx, os.Stdin, cli.Filename, 1024, -1)
		} else {
			contents, lexer = prepareLenient(ctx, os.Stdin, cli.Filename)
		}
		format(ctx, w, style, lex(ctx, lexer, contents))
	} else {
		for _, filename := range cli.Files {
			file, err := os.Open(filename)
			ctx.FatalIfErrorf(err)

			if cli.Check {
				contents, lexer := prepareLenient(ctx, file, filename)
				check(filename, lex(ctx, lexer, contents))
			} else {
				var contents string
				var lexer chroma.Lexer
				if cli.Fail {
					fi, err := file.Stat()
					ctx.FatalIfErrorf(err)
					contents, lexer = prepareSpecific(ctx, file, filename, 1024, int(fi.Size()))
				} else {
					contents, lexer = prepareLenient(ctx, file, filename)
				}
				format(ctx, w, style, lex(ctx, lexer, contents))
			}

			err = file.Close()
			ctx.FatalIfErrorf(err)
		}
	}
}

func configureHTMLFormatter(ctx *kong.Context) {
	options := []html.Option{
		html.TabWidth(cli.HTMLTabWidth),
		html.BaseLineNumber(cli.HTMLBaseLine),
		html.ClassPrefix(cli.HTMLPrefix),
		html.WithAllClasses(cli.HTMLAllStyles),
		html.WithClasses(!cli.HTMLInlineStyles),
		html.Standalone(!cli.HTMLOnly),
		html.WithLineNumbers(cli.HTMLLines),
		html.LineNumbersInTable(cli.HTMLLinesTable),
		html.PreventSurroundingPre(cli.HTMLPreventSurroundingPre),
	}
	if len(cli.HTMLHighlight) > 0 {
		ranges := [][2]int{}
		for _, span := range strings.Split(cli.HTMLHighlight, ",") {
			parts := strings.Split(span, ":")
			if len(parts) > 2 {
				ctx.Fatalf("range should be N[:M], not %q", span)
			}
			start, err := strconv.ParseInt(parts[0], 10, 64)
			ctx.FatalIfErrorf(err, "min value of range should be integer not %q", parts[0])
			end := start
			if len(parts) == 2 {
				end, err = strconv.ParseInt(parts[1], 10, 64)
				ctx.FatalIfErrorf(err, "max value of range should be integer not %q", parts[1])
			}
			ranges = append(ranges, [2]int{int(start), int(end)})
		}
		options = append(options, html.HighlightLines(ranges))
	}
	formatters.Register("html", html.New(options...))
}

func listAll() {
	fmt.Println("lexers:")
	sort.Sort(lexers.Registry.Lexers)
	for _, l := range lexers.Registry.Lexers {
		config := l.Config()
		fmt.Printf("  %s\n", config.Name)
		filenames := []string{}
		filenames = append(filenames, config.Filenames...)
		filenames = append(filenames, config.AliasFilenames...)
		if len(config.Aliases) > 0 {
			fmt.Printf("    aliases: %s\n", strings.Join(config.Aliases, " "))
		}
		if len(filenames) > 0 {
			fmt.Printf("    filenames: %s\n", strings.Join(filenames, " "))
		}
		if len(config.MimeTypes) > 0 {
			fmt.Printf("    mimetypes: %s\n", strings.Join(config.MimeTypes, " "))
		}
	}
	fmt.Println()
	fmt.Printf("styles:")
	for _, name := range styles.Names() {
		fmt.Printf(" %s", name)
	}
	fmt.Println()
	fmt.Printf("formatters:")
	for _, name := range formatters.Names() {
		fmt.Printf(" %s", name)
	}
	fmt.Println()
}

func lex(ctx *kong.Context, lexer chroma.Lexer, contents string) chroma.Iterator {
	if rel, ok := lexer.(*chroma.RegexLexer); ok {
		rel.Trace(cli.Trace)
	}
	lexer = chroma.Coalesce(lexer)
	it, err := lexer.Tokenise(nil, contents)
	ctx.FatalIfErrorf(err)
	return it
}

func selexer(path, contents string) (lexer chroma.Lexer) {
	if cli.Lexer != "autodetect" {
		return lexers.Get(cli.Lexer)
	}
	if path != "" {
		lexer := lexers.Match(path)
		if lexer != nil {
			return lexer
		}
	}
	return lexers.Analyse(contents)
}

func format(ctx *kong.Context, w io.Writer, style *chroma.Style, it chroma.Iterator) {
	formatter := formatters.Get(cli.Formatter)
	err := formatter.Format(w, style, it)
	ctx.FatalIfErrorf(err)
}

func check(filename string, it chroma.Iterator) {
	line, col := 1, 0
	for token := it(); token != chroma.EOF; token = it() {
		if token.Type == chroma.Error {
			fmt.Printf("%s:%d:%d %q\n", filename, line, col, token.String())
		}
		for _, c := range token.String() {
			col++
			if c == '\n' {
				line, col = line+1, 0
			}
		}
	}
}
