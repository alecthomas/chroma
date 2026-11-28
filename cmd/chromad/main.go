package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"mime"
	"net/http"
	"sort"
	"strings"

	"github.com/alecthomas/kong"
	konghcl "github.com/alecthomas/kong-hcl"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/alecthomas/chroma/v3"
	"github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/internal/playground"
	"github.com/alecthomas/chroma/v3/lexers"
	"github.com/alecthomas/chroma/v3/styles"
)

var (
	version = "devel"

	//go:embed templates/index.html.tmpl
	indexTemplate string
	//go:embed static
	staticFiles embed.FS

	htmlTemplate = template.Must(template.New("html").
			Funcs(template.FuncMap{
			"JS": func(filename string) template.JS {
				if version == "devel" {
					return template.JS(`import "./static/` + filename + "\";\n")
				}
				content, err := staticFiles.ReadFile("static/" + strings.TrimSuffix(filename, ".js") + ".min.js")
				if err != nil {
					panic(err)
				}
				return template.JS(content)
			},
			"CSS": func(filename string) template.CSS {
				if version == "devel" {
					return template.CSS(`@import url("./static/` + filename + "\");")
				}
				content, err := staticFiles.ReadFile("static/" + strings.TrimSuffix(filename, ".css") + ".min.css")
				if err != nil {
					panic(err)
				}
				return template.CSS(content)
			},
		}).Parse(indexTemplate))
)

type styleInfo struct {
	Name        string `json:"name"`
	Mode        string `json:"mode"`
	Counterpart string `json:"counterpart,omitempty"`
}

type context struct {
	Background       template.CSS
	SelectedLanguage string
	Languages        []string
	SelectedStyle    string
	Styles           []styleInfo
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
	res, err := playground.Highlight(req.Text, req.Language, req.Style, req.Classes)
	if err != nil {
		return nil, err
	}
	return &renderResponse{
		Language:   res.Language,
		HTML:       res.HTML,
		Background: res.Background,
	}, nil
}

func newContext(r *http.Request) context {
	ctx := context{
		SelectedStyle: "monokailight",
		CSRFField:     csrf.TemplateField(r),
		Version:       version,
	}
	style := styles.Get(ctx.SelectedStyle)
	if style == nil {
		style = styles.Fallback
	}
	ctx.Background = template.CSS(html.StyleEntryToCSS(style.Get(chroma.Background)))
	if ctx.SelectedStyle == "" {
		ctx.SelectedStyle = "monokailight"
	}
	for _, lexer := range lexers.All() {
		ctx.Languages = append(ctx.Languages, lexer.Config().Name)
	}
	sort.Strings(ctx.Languages)
	for _, name := range styles.Names() {
		s := styles.Get(name)
		ctx.Styles = append(ctx.Styles, styleInfo{
			Name:        name,
			Mode:        s.Mode().String(),
			Counterpart: s.Counterpart,
		})
	}
	return ctx
}

func main() {
	var cli struct {
		Version kong.VersionFlag `help:"Show version."`
		Config  kong.ConfigFlag  `help:"Load configuration." placeholder:"FILE"`
		Bind    string           `help:"HTTP bind address." default:"127.0.0.1:8080"`
		CSRFKey string           `help:"CSRF key." default:""`
	}
	ctx := kong.Parse(&cli, kong.Configuration(konghcl.Loader), kong.Vars{"version": version})

	log.Printf("Starting chromad %s on http://%s\n", version, cli.Bind)

	mime.AddExtensionType(".js", "application/javascript")

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(index)).Methods("GET")
	router.Handle("/api/render", http.HandlerFunc(renderHandler)).Methods("POST")
	router.Handle("/static/{file:.*}", http.FileServer(http.FS(staticFiles))).Methods("GET")

	options := []csrf.Option{}
	if cli.CSRFKey == "" {
		options = append(options, csrf.Secure(false))
	}

	root := handlers.CORS()(csrf.Protect([]byte(cli.CSRFKey), options...)(router))

	err := http.ListenAndServe(cli.Bind, root)
	ctx.FatalIfErrorf(err)
}
