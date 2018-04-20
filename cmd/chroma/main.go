package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"gopkg.in/alecthomas/kingpin.v3-unstable"

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

	profileFlag    = kingpin.Flag("profile", "Enable profiling to file.").Hidden().String()
	listFlag       = kingpin.Flag("list", "List lexers, styles and formatters.").Bool()
	unbufferedFlag = kingpin.Flag("unbuffered", "Do not buffer output.").Bool()
	traceFlag      = kingpin.Flag("trace", "Trace lexer states as they are traversed.").Bool()
	checkFlag      = kingpin.Flag("check", "Do not format, check for tokenization errors instead.").Bool()
	filenameFlag   = kingpin.Flag("filename", "Filename to use for selecting a lexer when reading from stdin.").String()

	lexerFlag     = kingpin.Flag("lexer", "Lexer to use when formatting.").PlaceHolder("autodetect").Short('l').Enum(lexers.Names(true)...)
	styleFlag     = kingpin.Flag("style", "Style to use for formatting.").Short('s').Default("swapoff").Enum(styles.Names()...)
	formatterFlag = kingpin.Flag("formatter", "Formatter to use.").Default("terminal").Short('f').Enum(formatters.Names()...)

	jsonFlag = kingpin.Flag("json", "Output JSON representation of tokens.").Bool()

	htmlFlag               = kingpin.Flag("html", "Enable HTML mode (equivalent to '--formatter html').").Bool()
	htmlPrefixFlag         = kingpin.Flag("html-prefix", "HTML CSS class prefix.").PlaceHolder("PREFIX").String()
	htmlStylesFlag         = kingpin.Flag("html-styles", "Output HTML CSS styles.").Bool()
	htmlOnlyFlag           = kingpin.Flag("html-only", "Output HTML fragment.").Bool()
	htmlInlineStyleFlag    = kingpin.Flag("html-inline-styles", "Output HTML with inline styles (no classes).").Bool()
	htmlTabWidthFlag       = kingpin.Flag("html-tab-width", "Set the HTML tab width.").Default("8").Int()
	htmlLinesFlag          = kingpin.Flag("html-lines", "Include line numbers in output.").Bool()
	htmlLinesTableFlag     = kingpin.Flag("html-lines-table", "Split line numbers and code in a HTML table").Bool()
	htmlLinesStyleFlag     = kingpin.Flag("html-lines-style", "Style for line numbers.").String()
	htmlHighlightFlag      = kingpin.Flag("html-highlight", "Highlight these lines.").PlaceHolder("N[:M][,...]").String()
	htmlHighlightStyleFlag = kingpin.Flag("html-highlight-style", "Style used for highlighting lines.").String()
	htmlBaseLineFlag       = kingpin.Flag("html-base-line", "Base line number.").Default("1").Int()

	filesArgs = kingpin.Arg("files", "Files to highlight.").ExistingFiles()
)

type flushableWriter interface {
	io.Writer
	Flush() error
}

type nopFlushableWriter struct{ io.Writer }

func (n *nopFlushableWriter) Flush() error { return nil }

func main() {
	kingpin.CommandLine.Version(fmt.Sprintf("%s-%s-%s", version, commit, date))
	kingpin.CommandLine.Help = `
Chroma is a general purpose syntax highlighting library and corresponding
command, for Go.
`
	kingpin.Parse()
	if *listFlag {
		listAll()
		return
	}
	if *profileFlag != "" {
		f, err := os.Create(*profileFlag)
		kingpin.FatalIfError(err, "")
		pprof.StartCPUProfile(f)
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		go func() {
			<-signals
			pprof.StopCPUProfile()
			os.Exit(128 + 3)
		}()
		defer pprof.StopCPUProfile()
	}

	var out io.Writer = os.Stdout
	if runtime.GOOS == "windows" && isatty.IsTerminal(os.Stdout.Fd()) {
		out = colorable.NewColorableStdout()
	}
	var w flushableWriter
	if *unbufferedFlag {
		w = &nopFlushableWriter{out}
	} else {
		w = bufio.NewWriterSize(out, 16384)
	}
	defer w.Flush()

	if *jsonFlag {
		*formatterFlag = "json"
	}

	if *htmlFlag {
		*formatterFlag = "html"
	}

	// Retrieve user-specified style, clone it, and add some overrides.
	builder := styles.Get(*styleFlag).Builder()
	if *htmlHighlightStyleFlag != "" {
		builder.Add(chroma.LineHighlight, *htmlHighlightStyleFlag)
	}
	if *htmlLinesStyleFlag != "" {
		builder.Add(chroma.LineNumbers, *htmlLinesStyleFlag)
	}
	style, err := builder.Build()
	kingpin.FatalIfError(err, "")

	// Dump styles.
	if *htmlStylesFlag {
		formatter := html.New(html.WithClasses())
		formatter.WriteCSS(w, style)
		return
	}

	if *formatterFlag == "html" {
		options := []html.Option{
			html.TabWidth(*htmlTabWidthFlag),
			html.BaseLineNumber(*htmlBaseLineFlag),
		}
		if *htmlPrefixFlag != "" {
			options = append(options, html.ClassPrefix(*htmlPrefixFlag))
		}
		if !*htmlInlineStyleFlag {
			options = append(options, html.WithClasses())
		}
		if !*htmlOnlyFlag {
			options = append(options, html.Standalone())
		}
		if *htmlLinesFlag {
			options = append(options, html.WithLineNumbers())
		}
		if *htmlLinesTableFlag {
			options = append(options, html.LineNumbersInTable())
		}
		if len(*htmlHighlightFlag) > 0 {
			ranges := [][2]int{}
			for _, span := range strings.Split(*htmlHighlightFlag, ",") {
				parts := strings.Split(span, ":")
				if len(parts) > 2 {
					kingpin.Fatalf("range should be N[:M], not %q", span)
				}
				start, err := strconv.ParseInt(parts[0], 10, 64)
				kingpin.FatalIfError(err, "min value of range should be integer not %q", parts[0])
				end := start
				if len(parts) == 2 {
					end, err = strconv.ParseInt(parts[1], 10, 64)
					kingpin.FatalIfError(err, "max value of range should be integer not %q", parts[1])
				}
				ranges = append(ranges, [2]int{int(start), int(end)})
			}
			options = append(options, html.HighlightLines(ranges))
		}
		formatters.Register("html", html.New(options...))
	}
	if len(*filesArgs) == 0 {
		contents, err := ioutil.ReadAll(os.Stdin)
		kingpin.FatalIfError(err, "")
		format(w, style, lex(*filenameFlag, string(contents)))
	} else {
		for _, filename := range *filesArgs {
			contents, err := ioutil.ReadFile(filename)
			kingpin.FatalIfError(err, "")
			if *checkFlag {
				check(filename, lex(filename, string(contents)))
			} else {
				format(w, style, lex(filename, string(contents)))
			}
		}
	}
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

func lex(path string, contents string) chroma.Iterator {
	lexer := selexer(path, contents)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	if rel, ok := lexer.(*chroma.RegexLexer); ok {
		rel.Trace(*traceFlag)
	}
	lexer = chroma.Coalesce(lexer)
	it, err := lexer.Tokenise(nil, string(contents))
	kingpin.FatalIfError(err, "")
	return it
}

func selexer(path, contents string) (lexer chroma.Lexer) {
	if *lexerFlag != "" {
		return lexers.Get(*lexerFlag)
	}
	if path != "" {
		lexer := lexers.Match(path)
		if lexer != nil {
			return lexer
		}
	}
	return lexers.Analyse(contents)
}

func format(w io.Writer, style *chroma.Style, it chroma.Iterator) {
	formatter := formatters.Get(*formatterFlag)
	err := formatter.Format(w, style, it)
	kingpin.FatalIfError(err, "")
}

func check(filename string, it chroma.Iterator) {
	line, col := 1, 0
	for token := it(); token != nil; token = it() {
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
