package main

import (
	"fmt"
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
	formatter := formatters.Console(formatters.DefaultConsoleTheme)
	for _, filename := range *filesArgs {
		lexers := lexers.Registry.Match(filename)
		lexer := lexers[0]
		lexer = chroma.Coalesce(lexer)
		contents, err := ioutil.ReadFile(filename)
		kingpin.FatalIfError(err, "")
		tokens, err := lexer.Tokenise(string(contents))
		kingpin.FatalIfError(err, "")
		if *tokensFlag {
			for _, token := range tokens {
				fmt.Println(token)
			}
		} else {
			formatter.Format(os.Stdout, tokens)
		}
	}
}
