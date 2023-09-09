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
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	colorable "github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
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
		XML        string           `hidden:"" help:"Generate XML lexer definitions." type:"existingdir" placeholder:"DIR"`

		Lexer string `group:"select" help:"Lexer to use when formatting or path to an XML file to load." default:"autodetect" short:"l"`
		Style string `group:"select" help:"Style to use for formatting or path to an XML file to load." default:"swapoff" short:"s"`

		Formatter string `group:"format" help:"Formatter to use." default:"terminal" short:"f" enum:"${formatters}"`
		JSON      bool   `group:"format" help:"Convenience flag to use JSON formatter."`
		HTML      bool   `group:"format" help:"Convenience flag to use HTML formatter."`
		SVG       bool   `group:"format" help:"Convenience flag to use SVG formatter."`

		HTMLPrefix                string `group:"html" help:"HTML CSS class prefix." placeholder:"PREFIX"`
		HTMLStyles                bool   `group:"html" help:"Output HTML CSS styles."`
		HTMLAllStyles             bool   `group:"html" help:"Output all HTML CSS styles, including redundant ones."`
		HTMLOnly                  bool   `group:"html" help:"Output HTML fragment."`
		HTMLInlineStyles          bool   `group:"html" help:"Output HTML with inline styles (no classes)."`
		HTMLTabWidth              int    `group:"html" help:"Set the HTML tab width." default:"8"`
		HTMLLines                 bool   `group:"html" help:"Include line numbers in output."`
		HTMLLinesTable            bool   `group:"html" help:"Split line numbers and code in a HTML table"`
		HTMLLinesStyle            string `group:"html" help:"Style for line numbers."`
		HTMLHighlight             string `group:"html" help:"Highlight these lines." placeholder:"N[:M][,...]"`
		HTMLHighlightStyle        string `group:"html" help:"Style used for highlighting lines."`
		HTMLBaseLine              int    `group:"html" help:"Base line number." default:"1"`
		HTMLPreventSurroundingPre bool   `group:"html" help:"Prevent the surrounding pre tag."`
		HTMLLinkableLines         bool   `group:"html" help:"Make the line numbers linkable and be a link to themselves."`

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
func prepareLenient(r io.Reader, filename string) (string, chroma.Lexer, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil, err
	}

	contents := string(data)
	lexer, err := selexer(filename, contents)
	if err != nil {
		return "", nil, err
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}

	return contents, lexer, nil
}

// prepareSpecific prepares contents and lexer for input, exiting if there is no specific lexer available for it.
// Input is consumed only up to peekSize for lexer selection. With fullSize -1, consume r until EOF.
func prepareSpecific(ctx *kong.Context, r io.Reader, filename string, peekSize, fullSize int) (string, chroma.Lexer) {
	data := make([]byte, peekSize)
	n, err := io.ReadFull(r, data)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		ctx.FatalIfErrorf(err)
	}

	lexer, err := selexer(filename, string(data[:n]))
	if err != nil {
		ctx.FatalIfErrorf(err)
	}
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
		ndata = make([]byte, n+len(rest))
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
	}, kong.Groups{
		"format": "Output format:",
		"select": "Select lexer and style:",
		"html":   "HTML formatter options:",
	})
	if cli.XML != "" {
		err := dumpXMLLexerDefinitions(cli.XML)
		ctx.FatalIfErrorf(err)
		return
	}
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
	selectedStyle, err := selectStyle()
	ctx.FatalIfErrorf(err)
	builder := selectedStyle.Builder()
	if cli.HTMLHighlightStyle != "" {
		builder.Add(chroma.LineHighlight, cli.HTMLHighlightStyle)
	}
	if cli.HTMLLinesStyle != "" {
		builder.Add(chroma.LineNumbers, cli.HTMLLinesStyle)
	}
	style, err := builder.Build()
	ctx.FatalIfErrorf(err)

	if cli.Formatter == "html" {
		configureHTMLFormatter(ctx)
	}

	// Dump styles.
	if cli.HTMLStyles {
		formatter := formatters.Get("html").(*html.Formatter)
		err = formatter.WriteCSS(w, style)
		ctx.FatalIfErrorf(err)
		return
	}

	if len(cli.Files) == 0 {
		var contents string
		var lexer chroma.Lexer
		if cli.Fail {
			contents, lexer = prepareSpecific(ctx, os.Stdin, cli.Filename, 1024, -1)
		} else {
			contents, lexer, err = prepareLenient(os.Stdin, cli.Filename)
			ctx.FatalIfErrorf(err)
		}
		format(ctx, w, style, lex(ctx, lexer, contents))
	} else {
		for _, filename := range cli.Files {
			file, err := os.Open(filename)
			ctx.FatalIfErrorf(err)

			if cli.Check {
				contents, lexer, err := prepareLenient(file, filename)
				ctx.FatalIfErrorf(err)
				check(filename, lex(ctx, lexer, contents))
			} else {
				var contents string
				var lexer chroma.Lexer
				if cli.Fail {
					fi, err := file.Stat()
					ctx.FatalIfErrorf(err)
					contents, lexer = prepareSpecific(ctx, file, filename, 1024, int(fi.Size()))
				} else {
					contents, lexer, err = prepareLenient(file, filename)
					ctx.FatalIfErrorf(err)
				}
				format(ctx, w, style, lex(ctx, lexer, contents))
			}

			err = file.Close()
			ctx.FatalIfErrorf(err)
		}
	}
}

func selectStyle() (*chroma.Style, error) {
	style, ok := styles.Registry[cli.Style]
	if ok {
		return style, nil
	}
	r, err := os.Open(cli.Style)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return chroma.NewXMLStyle(r)
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
		html.WithLinkableLineNumbers(cli.HTMLLinkableLines, "L"),
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
	sort.Sort(lexers.GlobalLexerRegistry.Lexers)
	for _, l := range lexers.GlobalLexerRegistry.Lexers {
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

func selexer(path, contents string) (lexer chroma.Lexer, err error) {
	if cli.Lexer == "autodetect" {
		if path != "" {
			lexer := lexers.Match(path)
			if lexer != nil {
				return lexer, nil
			}
		}
		return lexers.Analyse(contents), nil
	}

	if lexer := lexers.Get(cli.Lexer); lexer != nil {
		return lexer, nil
	}
	lexerPath, err := filepath.Abs(cli.Lexer)
	return chroma.NewXMLLexer(os.DirFS("/"), lexerPath)
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

var nameCleanRe = regexp.MustCompile(`[^A-Za-z0-9_#+-]`)

func dumpXMLLexerDefinitions(dir string) error {
	for _, name := range lexers.Names(false) {
		lex := lexers.Get(name)
		if rlex, ok := lex.(*chroma.RegexLexer); ok {
			data, err := chroma.Marshal(rlex)
			if err != nil {
				if errors.Is(err, chroma.ErrNotSerialisable) {
					fmt.Fprintf(os.Stderr, "warning: %q: %s\n", name, err)
					continue
				}
				return err
			}
			name := strings.ToLower(nameCleanRe.ReplaceAllString(lex.Config().Name, "_"))
			filename := filepath.Join(dir, name) + ".xml"
			// fmt.Println(name)
			_, err = os.Stat(filename)
			if err == nil {
				fmt.Fprintf(os.Stderr, "warning: %s already exists\n", filename)
				continue
			}
			err = os.WriteFile(filename, data, 0600)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
