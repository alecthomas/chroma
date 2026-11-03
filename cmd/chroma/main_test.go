package main

import (
	"strings"
	"testing"

	"github.com/alecthomas/kong"

	"github.com/alecthomas/chroma/v3/formatters"
	"github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/styles"
)

// --html-styles dumps CSS from the html formatter even when the selected
// formatter is something else, so the html options must still apply.
func TestHTMLStylesAppliesHTMLOptions(t *testing.T) {
	saved := cli
	t.Cleanup(func() {
		cli = saved
		formatters.Register("html", html.New(html.Standalone(true), html.WithClasses(true)))
	})

	cli.Formatter = "terminal"
	cli.HTMLStyles = true
	cli.HTMLPrefix = "pfx-"
	cli.HTMLTabWidth = 4

	if !usesHTMLFormatter() {
		t.Fatal("expected html formatter to be in use")
	}
	configureHTMLFormatter(&kong.Context{})

	w := &strings.Builder{}
	if err := formatters.Get("html").(*html.Formatter).WriteCSS(w, styles.Fallback); err != nil {
		t.Fatal(err)
	}
	out := w.String()
	if !strings.Contains(out, ".pfx-chroma") {
		t.Errorf("--html-prefix not applied:\n%s", out)
	}
	if !strings.Contains(out, "tab-size: 4") {
		t.Errorf("--html-tab-width not applied:\n%s", out)
	}
}
