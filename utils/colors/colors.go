package colors

import "github.com/charmbracelet/lipgloss"

type Scale struct {
	Name  string
	Steps []step
}

func (s Scale) Step(n int) lipgloss.TerminalColor {
	return s.Steps[n-1].Color
}

type step struct {
	Name  string
	Color lipgloss.TerminalColor
}

var Scales = []Scale{
	Tomato,
	Indigo,
	Green,
}
