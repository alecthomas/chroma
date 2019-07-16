package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/alecthomas/kong"
	"github.com/gorilla/mux"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var (
	templateFiles = rice.MustFindBox("templates")
	staticFiles   = rice.MustFindBox("static")

	htmlTemplate = template.Must(template.New("html").Parse(templateFiles.MustString("index.html.tmpl")))
)

var cli struct {
	Bind string `help:"HTTP bind address." default:"127.0.0.1:8080"`
}

type context struct {
	Background       template.CSS
	SelectedLanguage string
	Languages        []string
	SelectedStyle    string
	Styles           []string
	Text             string
	HTML             template.HTML
	Error            string
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := contextFromRequest(r)
	style := styles.Get(ctx.SelectedStyle)
	if style == nil {
		style = styles.Fallback
	}
	ctx.Background = template.CSS(html.StyleEntryToCSS(style.Get(chroma.Background)))

	language := lexers.Get(ctx.SelectedLanguage)
	if language == nil {
		language = lexers.Analyse(ctx.Text)
		if language != nil {
			ctx.SelectedLanguage = language.Config().Name
		}
	}
	if language == nil {
		language = lexers.Fallback
	}

	tokens, err := language.Tokenise(nil, ctx.Text)
	if err != nil {
		ctx.Error = err.Error()
	} else {
		buf := &strings.Builder{}
		formatter := html.New()
		err = formatter.Format(buf, style, tokens)
		if err != nil {
			ctx.Error = err.Error()
		} else {
			ctx.HTML = template.HTML(buf.String()) // nolint: gosec
		}
	}

	err = htmlTemplate.Execute(w, &ctx)
	if err != nil {
		panic(err)
	}
}

func contextFromRequest(r *http.Request) context {
	err := r.ParseForm()
	ctx := context{
		SelectedLanguage: r.Form.Get("language"),
		SelectedStyle:    r.Form.Get("style"),
		Text:             r.Form.Get("text"),
	}
	if err != nil {
		ctx.Error = err.Error()
		return ctx
	}
	if ctx.SelectedStyle == "" {
		ctx.SelectedStyle = "monokailight"
	}
	for _, lexer := range lexers.Registry.Lexers {
		ctx.Languages = append(ctx.Languages, lexer.Config().Name)
	}
	sort.Strings(ctx.Languages)
	for _, style := range styles.Registry {
		ctx.Styles = append(ctx.Styles, style.Name)
	}
	sort.Strings(ctx.Styles)
	return ctx
}

func main() {
	ctx := kong.Parse(&cli)
	log.Println("Starting")

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(handler))
	router.Handle("/static/{file:.*}", http.StripPrefix("/static/", http.FileServer(staticFiles.HTTPBox())))

	err := http.ListenAndServe(cli.Bind, router)
	ctx.FatalIfErrorf(err)
}
