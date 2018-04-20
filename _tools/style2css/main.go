package main

import (
	"os"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	nameArg = kingpin.Arg("name", "Name of input style.").Required().Enum(styles.Names()...)
)

func main() {
	kingpin.Parse()

	formatter := html.New(html.WithClasses())
	err := formatter.WriteCSS(os.Stdout, styles.Get("solarized-dark"))
	kingpin.FatalIfError(err, "")
}
