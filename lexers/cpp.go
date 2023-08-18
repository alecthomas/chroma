package lexers

import (
	"regexp"

	. "github.com/alecthomas/chroma/v2" // nolint
)

var (
	cppAnalyserIncludeRe   = regexp.MustCompile(`#include <[a-z_]+>`)
	cppAnalyserNamespaceRe = regexp.MustCompile(`using namespace `)
)

var CPP = Register(MustNewXMLLexer(
	embedded,
	"embedded/c++.xml",
).SetConfig(
	&Config{
		Name:      "C++",
		Aliases:   []string{"cpp", "c++"},
		Filenames: []string{"*.cpp", "*.hpp", "*.c++", "*.h++", "*.cc", "*.hh", "*.cxx", "*.hxx", "*.C", "*.H", "*.cp", "*.CPP", "*.cppm", "*.ixx", "*.tpp"},
		MimeTypes: []string{"text/x-c++hdr", "text/x-c++src"},
		Priority:  0.1,
		EnsureNL:  true,
	},
)).SetAnalyser(func(text string) float32 {
	if cppAnalyserIncludeRe.MatchString(text) {
		return 0.2
	}

	if cppAnalyserNamespaceRe.MatchString(text) {
		return 0.4
	}

	return 0
})
