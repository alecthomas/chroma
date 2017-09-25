package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
	"gopkg.in/alecthomas/kingpin.v3-unstable"

	"github.com/alecthomas/chroma"
)

const (
	outputTemplate = `package styles

import (
	"github.com/alecthomas/chroma"
)

// {{.Name}} style.
var {{.Name}} = Register(chroma.MustNewStyle("{{.Name|Lower}}", chroma.StyleEntries{
{{- range .Rules}}
{{- if .Prelude|TokenType}}
	chroma.{{.Prelude|TokenType}}: "{{.Declarations|TranslateDecls}}",
{{- end}}
{{- end}}
}))
`
)

var (
	typeByClass = map[string]chroma.TokenType{
		".hll": chroma.Background,
	}

	cssNamedColours = map[string]string{
		"black": "#000000", "silver": "#c0c0c0", "gray": "#808080", "white": "#ffffff",
		"maroon": "#800000", "red": "#ff0000", "purple": "#800080", "fuchsia": "#ff00ff",
		"green": "#008000", "lime": "#00ff00", "olive": "#808000", "yellow": "#ffff00",
		"navy": "#000080", "blue": "#0000ff", "teal": "#008080", "aqua": "#00ffff",
		"orange": "#ffa500", "aliceblue": "#f0f8ff", "antiquewhite": "#faebd7", "aquamarine": "#7fffd4",
		"azure": "#f0ffff", "beige": "#f5f5dc", "bisque": "#ffe4c4", "blanchedalmond": "#ffebcd",
		"blueviolet": "#8a2be2", "brown": "#a52a2a", "burlywood": "#deb887", "cadetblue": "#5f9ea0",
		"chartreuse": "#7fff00", "chocolate": "#d2691e", "coral": "#ff7f50", "cornflowerblue": "#6495ed",
		"cornsilk": "#fff8dc", "crimson": "#dc143c", "cyan": "#00ffff", "darkblue": "#00008b",
		"darkcyan": "#008b8b", "darkgoldenrod": "#b8860b", "darkgray": "#a9a9a9", "darkgreen": "#006400",
		"darkgrey": "#a9a9a9", "darkkhaki": "#bdb76b", "darkmagenta": "#8b008b", "darkolivegreen": "#556b2f",
		"darkorange": "#ff8c00", "darkorchid": "#9932cc", "darkred": "#8b0000", "darksalmon": "#e9967a",
		"darkseagreen": "#8fbc8f", "darkslateblue": "#483d8b", "darkslategray": "#2f4f4f", "darkslategrey": "#2f4f4f",
		"darkturquoise": "#00ced1", "darkviolet": "#9400d3", "deeppink": "#ff1493", "deepskyblue": "#00bfff",
		"dimgray": "#696969", "dimgrey": "#696969", "dodgerblue": "#1e90ff", "firebrick": "#b22222",
		"floralwhite": "#fffaf0", "forestgreen": "#228b22", "gainsboro": "#dcdcdc", "ghostwhite": "#f8f8ff",
		"gold": "#ffd700", "goldenrod": "#daa520", "greenyellow": "#adff2f", "grey": "#808080",
		"honeydew": "#f0fff0", "hotpink": "#ff69b4", "indianred": "#cd5c5c", "indigo": "#4b0082",
		"ivory": "#fffff0", "khaki": "#f0e68c", "lavender": "#e6e6fa", "lavenderblush": "#fff0f5",
		"lawngreen": "#7cfc00", "lemonchiffon": "#fffacd", "lightblue": "#add8e6", "lightcoral": "#f08080",
		"lightcyan": "#e0ffff", "lightgoldenrodyellow": "#fafad2", "lightgray": "#d3d3d3", "lightgreen": "#90ee90",
		"lightgrey": "#d3d3d3", "lightpink": "#ffb6c1", "lightsalmon": "#ffa07a", "lightseagreen": "#20b2aa",
		"lightskyblue": "#87cefa", "lightslategray": "#778899", "lightslategrey": "#778899", "lightsteelblue": "#b0c4de",
		"lightyellow": "#ffffe0", "limegreen": "#32cd32", "linen": "#faf0e6", "magenta": "#ff00ff",
		"mediumaquamarine": "#66cdaa", "mediumblue": "#0000cd", "mediumorchid": "#ba55d3", "mediumpurple": "#9370db",
		"mediumseagreen": "#3cb371", "mediumslateblue": "#7b68ee", "mediumspringgreen": "#00fa9a", "mediumturquoise": "#48d1cc",
		"mediumvioletred": "#c71585", "midnightblue": "#191970", "mintcream": "#f5fffa", "mistyrose": "#ffe4e1",
		"moccasin": "#ffe4b5", "navajowhite": "#ffdead", "oldlace": "#fdf5e6", "olivedrab": "#6b8e23",
		"orangered": "#ff4500", "orchid": "#da70d6", "palegoldenrod": "#eee8aa", "palegreen": "#98fb98",
		"paleturquoise": "#afeeee", "palevioletred": "#db7093", "papayawhip": "#ffefd5", "peachpuff": "#ffdab9",
		"peru": "#cd853f", "pink": "#ffc0cb", "plum": "#dda0dd", "powderblue": "#b0e0e6",
		"rosybrown": "#bc8f8f", "royalblue": "#4169e1", "saddlebrown": "#8b4513", "salmon": "#fa8072",
		"sandybrown": "#f4a460", "seagreen": "#2e8b57", "seashell": "#fff5ee", "sienna": "#a0522d",
		"skyblue": "#87ceeb", "slateblue": "#6a5acd", "slategray": "#708090", "slategrey": "#708090",
		"snow": "#fffafa", "springgreen": "#00ff7f", "steelblue": "#4682b4", "tan": "#d2b48c",
		"thistle": "#d8bfd8", "tomato": "#ff6347", "turquoise": "#40e0d0", "violet": "#ee82ee",
		"wheat": "#f5deb3", "whitesmoke": "#f5f5f5", "yellowgreen": "#9acd32", "rebeccapurple": "#663399",
	}

	nameArg = kingpin.Arg("name", "Name of output style.").Required().String()
	fileArg = kingpin.Arg("stylesheets", ".css file to import").Required().ExistingFile()
)

func init() {
	for tt, str := range chroma.StandardTypes {
		typeByClass["."+str] = tt
	}
}

func translateDecls(decls []*css.Declaration) string {
	out := []string{}
	for _, decl := range decls {
		switch decl.Property {
		case "color":
			clr := decl.Value
			if c, ok := cssNamedColours[clr]; ok {
				clr = c
			}
			out = append(out, clr)
		case "background-color":
			out = append(out, "bg:"+decl.Value)
		case "font-style":
			if strings.Contains(decl.Value, "italic") {
				out = append(out, "italic")
			}
		case "font-weight":
			if strings.Contains(decl.Value, "bold") {
				out = append(out, "bold")
			}
		case "text-decoration":
			if strings.Contains(decl.Value, "underline") {
				out = append(out, "underline")
			}
		}
	}
	return strings.Join(out, " ")
}

func main() {
	kingpin.Parse()
	source, err := ioutil.ReadFile(*fileArg)
	kingpin.FatalIfError(err, "")
	css, err := parser.Parse(string(source))
	kingpin.FatalIfError(err, "")

	context := map[string]interface{}{
		"Name":  *nameArg,
		"Rules": css.Rules,
	}
	tmpl := template.Must(template.New("style").Funcs(template.FuncMap{
		"Lower":          strings.ToLower,
		"TranslateDecls": translateDecls,
		"TokenType":      func(s string) chroma.TokenType { return typeByClass[s] },
	}).Parse(outputTemplate))
	err = tmpl.Execute(os.Stdout, context)
	kingpin.FatalIfError(err, "")
}
