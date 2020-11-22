package doctype

import (
	"fmt"
	"regexp"
	"strings"
)

var doctypeLookupRe = regexp.MustCompile(`(?ms)(<\?.*?\?>)?\s*<!DOCTYPE\s+([a-zA-Z_][a-zA-Z0-9]*(?: \s+[a-zA-Z_][a-zA-Z0-9]*\s+"[^"]*")?)[^>]*>`)

// MatchString check if the doctype matches a regular expression (if present).
func MatchString(text string, pattern string) (bool, error) {
	// Note that this method only checks the first part of a DOCTYPE.
	// eg: 'html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"'
	m := doctypeLookupRe.FindStringSubmatch(text)

	if len(m) == 0 {
		return false, nil
	}

	doctypeRe, err := regexp.Compile(fmt.Sprintf("(?i)%s", pattern))
	if err != nil {
		return false, fmt.Errorf("failed to compile doctype regex: %s", err)
	}

	return doctypeRe.MatchString(strings.TrimSpace(m[2])), nil
}
