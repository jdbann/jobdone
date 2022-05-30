package models_test

import (
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
)

var ansiCodes = regexp.MustCompile(`\x1b\[[\d;]m`)

// stripAnsi removes ANSI codes from provided string. Helpful for comparing
// returned values from View() calls on tea.Model instances.
func stripAnsi(s string) string {
	return ansiCodes.ReplaceAllString(s, "")
}

// cmdsEqual checks whether the provided tea.Cmd functions are both nil or both
// return the same value.
func cmdsEqual(a, b tea.Cmd) bool {
	if a == nil && b == nil {
		return true
	}

	if (a != nil && b != nil) && (a() != b()) {
		return false
	} else {
		return true
	}
}
