package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	filesArgs = kingpin.Arg("file", "Files to use to exercise lexers.").Required().ExistingFiles()
)

func main() {
	kingpin.CommandLine.Help = "Exercise linters against a list of files."
	kingpin.Parse()

	for _, file := range *filesArgs {
		lexer := lexers.Match(file)
		if lexer == nil {
			fmt.Printf("warning: could not find lexer for %q\n", file)
			continue
		}
		fmt.Printf("%s: ", file)
		os.Stdout.Sync()
		text, err := ioutil.ReadFile(file)
		kingpin.FatalIfError(err, "")
		it, err := lexer.Tokenise(nil, string(text))
		kingpin.FatalIfError(err, "%s failed to tokenise %q", lexer.Config().Name, file)
		err = formatters.NoOp.Format(ioutil.Discard, styles.SwapOff, it)
		kingpin.FatalIfError(err, "%s failed to format %q", lexer.Config().Name, file)
		fmt.Printf("ok\n")
	}
}
