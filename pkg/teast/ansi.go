package teast

import (
	"regexp"
	"strings"
)

var ansiCodes = regexp.MustCompile(`\x1b\[[\d;]+m`)

// StripAnsiCodes removes ANSI codes from provided string. Helpful for comparing
// returned values from View() calls on tea.Model instances.
func StripAnsiCodes(s string) string {
	return ansiCodes.ReplaceAllString(s, "")
}

// SquashWhitespace removes leading and trailing whitespace and condenses runs
// of whitespace to a single space character.
func SquashWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
