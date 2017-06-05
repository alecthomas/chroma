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
)

var (
	profileFlag = kingpin.Flag("profile", "Enable profiling to file.").PlaceHolder("FILE").String()
	tokensFlag  = kingpin.Flag("tokens", "Dump raw tokens.").Bool()
	lexerFlag   = kingpin.Flag("lexer", "Lexer to use when formatting (default is to autodetect).").Short('l').String()
	filesArgs   = kingpin.Arg("files", "Files to highlight.").ExistingFiles()
)

func main() {
	kingpin.Parse()
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
		lexer := lexers.Registry.Get(*lexerFlag)
		contents, err := ioutil.ReadAll(os.Stdin)
		kingpin.FatalIfError(err, "")
		err = lexer.Tokenise(nil, string(contents), writer)
		kingpin.FatalIfError(err, "")
	} else {
		for _, filename := range *filesArgs {
			lexers := lexers.Registry.Match(filename)
			lexer := lexers[0]
			lexer = chroma.Coalesce(lexer)
			contents, err := ioutil.ReadFile(filename)
			kingpin.FatalIfError(err, "")
			err = lexer.Tokenise(nil, string(contents), writer)
			kingpin.FatalIfError(err, "")
		}
	}
}

func getWriter(w io.Writer) func(*chroma.Token) {
	if *tokensFlag {
		return func(token *chroma.Token) {
			fmt.Println(token)
		}
	}
	formatter := formatters.Console(formatters.DefaultConsoleTheme)
	writer, err := formatter.Format(w)
	kingpin.FatalIfError(err, "")
	return writer
}
