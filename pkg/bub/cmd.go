package bub

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func Wait(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return nil
	}
}
