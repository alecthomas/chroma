package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strings"

	"gopkg.in/alecthomas/kingpin.v3-unstable"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var (
	profileFlag = kingpin.Flag("profile", "Enable profiling to file.").Hidden().String()
	listFlag    = kingpin.Flag("list", "List lexers, styles and formatters.").Bool()

	lexerFlag     = kingpin.Flag("lexer", "Lexer to use when formatting.").Default("autodetect").Short('l').String()
	styleFlag     = kingpin.Flag("style", "Style to use for formatting.").Short('s').Default("swapoff").String()
	formatterFlag = kingpin.Flag("formatter", "Formatter to use.").Default("terminal").Short('f').String()

	filesArgs = kingpin.Arg("files", "Files to highlight.").ExistingFiles()
)

func main() {
	kingpin.Parse()
	if *listFlag {
		listAll()
		return
	}
	if *profileFlag != "" {
		f, err := os.Create(*profileFlag)
		kingpin.FatalIfError(err, "")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	w := bufio.NewWriterSize(os.Stdout, 16384)
	defer w.Flush()
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
	for _, l := range lexers.Registry.Lexers {
		config := l.Config()
		fmt.Printf("  %s\n", config.Name)
		filenames := []string{}
		filenames = append(filenames, config.Filenames...)
		filenames = append(filenames, config.AliasFilenames...)
		fmt.Printf("    aliases: %s\n", strings.Join(config.Aliases, " "))
		fmt.Printf("    filenames: %s\n", strings.Join(filenames, " "))
		fmt.Printf("    mimetypes: %s\n", strings.Join(config.MimeTypes, " "))
		fmt.Printf("    priority: %d\n", config.Priority)
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
	lexer := chroma.Coalesce(selexer(path))
	err := lexer.Tokenise(nil, string(contents), writer)
	kingpin.FatalIfError(err, "")
}

func selexer(path string) chroma.Lexer {
	if *lexerFlag != "autodetect" {
		return lexers.Get(*lexerFlag)
	}
	return lexers.Match(path)[0]
}

func getWriter(w io.Writer) func(*chroma.Token) {
	style := styles.Get(*styleFlag)
	formatter := formatters.Get(*formatterFlag)
	// formatter := formatters.TTY8
	writer, err := formatter.Format(w, style)
	kingpin.FatalIfError(err, "")
	return writer
}
