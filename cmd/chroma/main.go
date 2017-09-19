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
	profileFlag = kingpin.Flag("profile", "Enable profiling to file.").Hidden().String()
	listFlag    = kingpin.Flag("list", "List lexers, styles and formatters.").Bool()

	lexerFlag     = kingpin.Flag("lexer", "Lexer to use when formatting.").PlaceHolder("autodetect").Short('l').Enum(lexers.Names(true)...)
	styleFlag     = kingpin.Flag("style", "Style to use for formatting.").Short('s').Default("swapoff").Enum(styles.Names()...)
	formatterFlag = kingpin.Flag("formatter", "Formatter to use.").Default("terminal").Short('f').Enum(formatters.Names()...)

	htmlFlag            = kingpin.Flag("html", "Enable HTML mode (equivalent to '--formatter html').").Bool()
	htmlPrefixFlag      = kingpin.Flag("html-prefix", "HTML CSS class prefix.").PlaceHolder("PREFIX").String()
	htmlStylesFlag      = kingpin.Flag("html-styles", "Output HTML CSS styles.").Bool()
	htmlOnlyFlag        = kingpin.Flag("html-only", "Output HTML fragment.").Bool()
	htmlInlineStyleFlag = kingpin.Flag("html-inline-styles", "Output HTML with inline styles (no classes).").Bool()

	filesArgs = kingpin.Arg("files", "Files to highlight.").ExistingFiles()
)

func main() {
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
	w := bufio.NewWriterSize(out, 16384)
	defer w.Flush()
	if *htmlFlag {
		*formatterFlag = "html"
	}
	if *formatterFlag == "html" {
		options := []html.Option{}
		if *htmlPrefixFlag != "" {
			options = append(options, html.ClassPrefix(*htmlPrefixFlag))
		}

		// Dump styles.
		if *htmlStylesFlag {
			formatter := html.New(html.WithClasses())
			formatter.WriteCSS(w, styles.Get(*styleFlag))
			return
		}
		if !*htmlInlineStyleFlag {
			options = append(options, html.WithClasses())
		}
		if !*htmlOnlyFlag {
			options = append(options, html.Standalone())
		}
		formatters.Register("html", html.New(options...))
	}
	writer := getWriter(w)
	if len(*filesArgs) == 0 {
		contents, err := ioutil.ReadAll(os.Stdin)
		kingpin.FatalIfError(err, "")
		lex("", string(contents), writer)
	} else {
		for _, filename := range *filesArgs {
			contents, err := ioutil.ReadFile(filename)
			kingpin.FatalIfError(err, "")
			lex(filename, string(contents), writer)
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
	for name := range styles.Registry {
		fmt.Printf(" %s", name)
	}
	fmt.Println()
	fmt.Printf("formatters:")
	for name := range formatters.Registry {
		fmt.Printf(" %s", name)
	}
	fmt.Println()
}

func lex(path string, contents string, writer func(*chroma.Token)) {
	lexer := selexer(path, contents)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)
	err := lexer.Tokenise(nil, string(contents), writer)
	kingpin.FatalIfError(err, "")
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

func getWriter(w io.Writer) func(*chroma.Token) {
	style := styles.Get(*styleFlag)
	formatter := formatters.Get(*formatterFlag)
	// formatter := formatters.TTY8
	writer, err := formatter.Format(w, style)
	kingpin.FatalIfError(err, "")
	return writer
}
