package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime/pprof"

	"gopkg.in/alecthomas/kingpin.v3-unstable"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var (
	profileFlag = kingpin.Flag("profile", "Enable profiling to file.").PlaceHolder("FILE").String()
	listFlag    = kingpin.Flag("list", "List lexers, styles and formatters.").Bool()

	tokensFlag    = kingpin.Flag("tokens", "Dump raw tokens.").Bool()
	lexerFlag     = kingpin.Flag("lexer", "Lexer to use when formatting.").Default("autodetect").Short('l').String()
	styleFlag     = kingpin.Flag("style", "Style to use for formatting.").Short('s').Default("swapoff").String()
	formatterFlag = kingpin.Flag("formatter", "Formatter to use.").Default("terminal").Short('f').String()
	filesArgs     = kingpin.Arg("files", "Files to highlight.").ExistingFiles()
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
		lex("", os.Stdin, writer)
	} else {
		for _, filename := range *filesArgs {
			r, err := os.Open(filename)
			kingpin.FatalIfError(err, "")
			lex(filename, r, writer)
			r.Close()
		}
	}
}

func listAll() {
	fmt.Printf("lexers:")
	for _, l := range lexers.Registry.Lexers {
		fmt.Printf(" %s", l.Config().Name)
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

func lex(path string, r io.Reader, writer func(*chroma.Token)) {
	contents, err := ioutil.ReadAll(r)
	kingpin.FatalIfError(err, "")
	kingpin.FatalIfError(err, "")
	lexer := chroma.Coalesce(selexer(path))
	err = lexer.Tokenise(nil, string(contents), writer)
	kingpin.FatalIfError(err, "")
}

func selexer(path string) chroma.Lexer {
	if *lexerFlag != "autodetect" {
		return lexers.Get(*lexerFlag)
	}
	return lexers.Match(path)[0]
}

func getWriter(w io.Writer) func(*chroma.Token) {
	if *tokensFlag {
		return func(token *chroma.Token) {
			fmt.Printf("%#v\n", token)
		}
	}
	style := styles.Get(*styleFlag)
	formatter := formatters.Get(*formatterFlag)
	// formatter := formatters.TTY8
	writer, err := formatter.Format(w, style)
	kingpin.FatalIfError(err, "")
	return writer
}
