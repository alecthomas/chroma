package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/alecthomas/kong"
	konghcl "github.com/alecthomas/kong-hcl"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
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

type context struct {
	Background       template.CSS
	SelectedLanguage string
	Languages        []string
	SelectedStyle    string
	Styles           []string
	CSRFField        template.HTML
	Version          string
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r)
	err := htmlTemplate.Execute(w, &ctx)
	if err != nil {
		panic(err)
	}
}

type renderRequest struct {
	Language string `json:"language"`
	Style    string `json:"style"`
	Text     string `json:"text"`
	Classes  bool   `json:"classes"`
}

type renderResponse struct {
	Error      string `json:"error,omitempty"`
	HTML       string `json:"html,omitempty"`
	Language   string `json:"language,omitempty"`
	Background string `json:"background,omitempty"`
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
	req := &renderRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	var rep *renderResponse
	if err != nil {
		rep = &renderResponse{Error: err.Error()}
	} else {
		rep, err = render(req)
		if err != nil {
			rep = &renderResponse{Error: err.Error()}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(rep)
}

func render(req *renderRequest) (*renderResponse, error) {
	language := lexers.Get(req.Language)
	if language == nil {
		language = lexers.Analyse(req.Text)
		if language != nil {
			req.Language = language.Config().Name
		}
	}
	if language == nil {
		language = lexers.Fallback
	}

	tokens, err := chroma.Coalesce(language).Tokenise(nil, req.Text)
	if err != nil {
		return nil, err
	}

	style := styles.Get(req.Style)
	if style == nil {
		style = styles.Fallback
	}

	buf := &strings.Builder{}
	options := []html.Option{}
	if req.Classes {
		options = append(options, html.WithClasses(true), html.Standalone(true))
	}
	formatter := html.New(options...)
	err = formatter.Format(buf, style, tokens)
	if err != nil {
		return nil, err
	}
	lang := language.Config().Name
	if language == lexers.Fallback {
		lang = ""
	}
	return &renderResponse{
		Language:   lang,
		HTML:       buf.String(),
		Background: html.StyleEntryToCSS(style.Get(chroma.Background)),
	}, nil
}

func newContext(r *http.Request) context {
	ctx := context{
		SelectedStyle: "monokailight",
		CSRFField:     csrf.TemplateField(r),
		Version:       fmt.Sprintf("%d", staticFiles.Time().Unix()),
	}
	style := styles.Get(ctx.SelectedStyle)
	if style == nil {
		style = styles.Fallback
	}
	ctx.Background = template.CSS(html.StyleEntryToCSS(style.Get(chroma.Background)))
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
	var cli struct {
		Config  kong.ConfigFlag `help:"Load configuration." placeholder:"FILE"`
		Bind    string          `help:"HTTP bind address." default:"127.0.0.1:8080"`
		CSRFKey string          `help:"CSRF key." default:""`
	}
	ctx := kong.Parse(&cli, kong.Configuration(konghcl.Loader))

	log.Println("Starting")

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(index)).Methods("GET")
	router.Handle("/api/render", http.HandlerFunc(renderHandler)).Methods("POST")
	router.Handle("/static/{file:.*}", http.StripPrefix("/static/", http.FileServer(staticFiles.HTTPBox()))).Methods("GET")

	options := []csrf.Option{}
	if cli.CSRFKey == "" {
		options = append(options, csrf.Secure(false))
	}

	root := handlers.CORS()(csrf.Protect([]byte(cli.CSRFKey), options...)(router))

	err := http.ListenAndServe(cli.Bind, root)
	ctx.FatalIfErrorf(err)
}
