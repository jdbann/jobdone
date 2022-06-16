package challenge

import tea "github.com/charmbracelet/bubbletea"

func SwitchCmd(c Definition) tea.Cmd {
	return func() tea.Msg {
		return ChangedMsg{Challenge: c}
	}
}
