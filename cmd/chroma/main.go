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
	profileFlag = kingpin.Flag("profile", "Enable profiling to file.").String()
	tokensFlag  = kingpin.Flag("tokens", "Dump raw tokens.").Bool()
	filesArgs   = kingpin.Arg("files", "Files to highlight.").Required().ExistingFiles()
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
	for _, filename := range *filesArgs {
		lexers := lexers.Registry.Match(filename)
		lexer := lexers[0]
		lexer = chroma.Coalesce(lexer)
		contents, err := ioutil.ReadFile(filename)
		kingpin.FatalIfError(err, "")
		err = lexer.Tokenise(string(contents), writer)
		kingpin.FatalIfError(err, "")
	}
}

func getWriter(w io.Writer) func(chroma.Token) {
	if *tokensFlag {
		return func(token chroma.Token) {
			fmt.Println(token)
		}
	} else {
		formatter := formatters.Console(formatters.DefaultConsoleTheme)
		writer, err := formatter.Format(w)
		kingpin.FatalIfError(err, "")
		return writer
	}
}
