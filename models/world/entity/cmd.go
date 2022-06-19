package entity

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const tps = 2

func TickCmd(width, height int) tea.Cmd {
	return tea.Tick(time.Second/time.Duration(tps), func(t time.Time) tea.Msg {
		return TickMsg{
			Height: height,
			Width:  width,
		}
	})
}
