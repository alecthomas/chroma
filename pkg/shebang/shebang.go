package shebang

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	splitPathRe    = regexp.MustCompile(`[/\\ ]`)
	shebangPattern = `(?i)^%s(\.(exe|cmd|bat|bin))?$`
)

// MatchString check if the given regular expression matches the last part of the
// shebang if one exists.
func MatchString(text string, pattern string) (bool, error) {
	firstLine := strings.ToLower(strings.Split(text, "\n")[0])
	if !strings.HasPrefix(firstLine, "#!") {
		return false, nil
	}

	var parts []string

	for _, line := range splitPathRe.Split(strings.TrimSpace(firstLine[2:]), -1) {
		if line != "" && !strings.HasPrefix(line, "-") {
			parts = append(parts, line)
		}
	}

	if len(parts) == 0 {
		return false, nil
	}

	lastPart := parts[len(parts)-1]

	shebangRe, err := regexp.Compile(fmt.Sprintf(shebangPattern, pattern))
	if err != nil {
		return false, fmt.Errorf("failed to compile shebang regex: %s", err)
	}

	return shebangRe.MatchString(lastPart), nil
}
