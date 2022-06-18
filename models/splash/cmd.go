package splash

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func AnimateCmd() tea.Cmd {
	return tea.Tick(time.Second/time.Duration(fps), func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}

func DismissCmd(duration time.Duration) func() tea.Msg {
	return func() tea.Msg {
		time.Sleep(duration)
		return CompleteMsg{}
	}
}
