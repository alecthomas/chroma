package p

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/shebang"
)

var (
	perl6AnalyserDecl      = regexp.MustCompile(`(?:my|our|has)\s+(?:['\w:-]+\s+)?[$@%&(]`)
	perl6AnalyserDeclScope = regexp.MustCompile(`^\s*(?:(?P<scope>my|our)\s+)?(?:module|class|role|enum|grammar)`)
	perl6AnalyserOperator  = regexp.MustCompile(`#.*`)
	perl6AnalyserShell     = regexp.MustCompile(`^\s*$`)
	perl6AnalyserV6        = regexp.MustCompile(`^\s*(?:use\s+)?v6(?:\.\d(?:\.\d)?)?;`)
	perl6BeginPodRe        = regexp.MustCompile(`^=\w+`)
	perl6EndPodRe          = regexp.MustCompile(`^=(?:end|cut)`)
)

// Perl6 lexer.
var Perl6 = internal.Register(MustNewLexer(
	&Config{
		Name:    "Perl6",
		Aliases: []string{"perl6", "pl6", "raku"},
		Filenames: []string{"*.pl", "*.pm", "*.nqp", "*.p6", "*.6pl", "*.p6l", "*.pl6",
			"*.6pm", "*.p6m", "*.pm6", "*.t", "*.raku", "*.rakumod", "*.rakutest", "*.rakudoc"},
		MimeTypes: []string{"text/x-perl6", "application/x-perl6"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if matched, _ := shebang.MatchString(text, "perl6|rakudo|niecza|pugs"); matched {
		return 1.0
	}

	var (
		result      float32
		hasPerlDecl bool
	)

	// Check for my/our/has declarations.
	if perl6AnalyserDecl.MatchString(text) {
		result = 0.8
		hasPerlDecl = true
	}

	// XXX handle block comments.
	lines := perl6StripPod(text)

	for _, line := range lines {
		line = perl6AnalyserOperator.ReplaceAllLiteralString(line, "")

		if perl6AnalyserShell.MatchString(line) {
			continue
		}

		// Match v6; use v6; use v6.0; use v6.0.0.
		if perl6AnalyserV6.MatchString(line) {
			return 1.0
		}

		// Match class, module, role, enum, grammar declarations.
		classDecl := perl6AnalyserDeclScope.FindStringSubmatch(line)
		if len(classDecl) > 0 {
			if hasPerlDecl || perl6GetSubgroups(classDecl)["scope"] != "" {
				return 1.0
			}

			result = 0.05
			continue
		}
		break
	}

	if strings.Contains(text, ":=") {
		// Same logic as Perl lexer.
		result /= 2
	}

	return result
}))

func perl6StripPod(text string) []string {
	var (
		inPod         bool
		strippedLines []string
	)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if perl6EndPodRe.MatchString(line) {
			inPod = false
			continue
		}

		if perl6BeginPodRe.MatchString(line) {
			inPod = true
			continue
		}

		if !inPod {
			strippedLines = append(strippedLines, line)
		}
	}

	return strippedLines
}

func perl6GetSubgroups(match []string) map[string]string {
	groups := make(map[string]string)

	for i, name := range perl6AnalyserDeclScope.SubexpNames() {
		if i > 0 && i < len(match) {
			groups[name] = match[i]
		}
	}

	return groups
}
