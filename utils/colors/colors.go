package colors

import "github.com/charmbracelet/lipgloss"

type scale struct {
	Name  string
	Steps []step
}

type step struct {
	Name  string
	Color lipgloss.TerminalColor
}

var Scales = []scale{
	Tomato,
	Indigo,
	Green,
}
