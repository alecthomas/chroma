package xml

import (
	"regexp"

	"github.com/alecthomas/chroma/pkg/doctype"
	"github.com/dlclark/regexp2"
)

var (
	tagRe            = regexp2.MustCompile(`(?ism)<(.+?)(\s.*?)?>.*?</.+?>`, regexp2.None)
	xmlDeclarationRe = regexp.MustCompile(`(?i)\s*<\?xml[^>]*\?>`)
)

// MatchString check if a text looks like XML.
func MatchString(text string) bool {
	// Check if a doctype exists or if we have some tags.
	if xmlDeclarationRe.MatchString(text) {
		return true
	}

	if matched, _ := doctype.MatchString(text, ""); matched {
		return true
	}

	if len(text) > 1000 {
		text = text[:1000]
	}

	if matched, _ := tagRe.MatchString(text); matched {
		return true
	}

	return false
}
