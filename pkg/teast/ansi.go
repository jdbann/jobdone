package teast

import "regexp"

var ansiCodes = regexp.MustCompile(`\x1b\[[\d;]m`)

// StripAnsiCodes removes ANSI codes from provided string. Helpful for comparing
// returned values from View() calls on tea.Model instances.
func StripAnsiCodes(s string) string {
	return ansiCodes.ReplaceAllString(s, "")
}
