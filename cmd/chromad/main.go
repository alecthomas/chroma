package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/alecthomas/kong"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var htmlTemplate = template.Must(template.New("html").Parse(`
<!doctype html>
<html>
<head>
	<title>Chroma Playground</title>
    <!-- other stuff here -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.5/css/bulma.min.css" />
	<style>
		textarea {
			font-family:Consolas,Monaco,Lucida Console,Liberation Mono,DejaVu Sans Mono,Bitstream Vera Sans Mono,Courier New, monospace;
		}
		#output {
			{{.Background}}
		}
		#output pre {
			padding: 0;
		}
	</style>
</head>
<body>
<div class="container">
	{{if .Error}}<div class="notification">{{.Error}}</div>{{end}}

<h1 class="title">Chroma Playground</h1>

<form method="post">
	<div class="columns">
		<div class="column field">
			<label class="label">Language</label>
			<div class="control">
				<div class="select">
					<select name="language">
						<option value="" disabled{{if eq "" $.SelectedLanguage}} selected{{end}}>Language</option>
					{{- range .Languages}}
						<option value="{{.}}"{{if eq . $.SelectedLanguage}} selected{{end}}>{{.}}</option>
					{{end -}}
					</select>
				</div>
			</div>
		</div>

		<div class="column field">
			<label class="label">Style</label>
			<div class="control">
				<div class="select">
					<select name="style">
						<option value="" disabled{{if eq "" $.SelectedStyle}} selected{{end}}>Style</option>
					{{- range .Styles}}
						<option value="{{.}}"{{if eq . $.SelectedStyle}} selected{{end}}>{{.}}</option>
					{{end -}}
					</select>
				</div>
			</div>
		</div>
	</div>

	<div class="field">
		<label class="label">Code</label>
		<div class="control">
			<textarea class="textarea" name="text" rows="25" cols="80">{{.Text}}</textarea>
		</div>
	</div>

	<div class="field">
		<div class="control">
			<button class="button is-link">Submit</button>
		</div>
	</div>

	<hr>

	<label class="label">Output</label>
	<div class="field box" id="output">
		{{.HTML}}
	</div>
</form>
</div>
</body>
</html>
`))

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
	err := http.ListenAndServe(cli.Bind, http.HandlerFunc(handler))
	ctx.FatalIfErrorf(err)
}
