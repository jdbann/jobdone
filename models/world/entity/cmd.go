package entity

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const tps = 5

func TickCmd() tea.Cmd {
	return tea.Tick(time.Second/time.Duration(tps), func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}
