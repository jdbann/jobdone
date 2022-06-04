package teast

import tea "github.com/charmbracelet/bubbletea"

// FakeCmd builds a fake tea.Cmd for testing that commands are returned from
// nested models.
func FakeCmd(msg interface{}) func() tea.Msg {
	return func() tea.Msg { return msg }
}
