package models

import tea "github.com/charmbracelet/bubbletea"

var _ tea.Model = Splash{}

type Splash struct {
}

func NewSplash() Splash {
	return Splash{}
}

func (s Splash) Init() tea.Cmd {
	return nil
}

func (s Splash) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s Splash) View() string {
	return "JOB DONE!"
}
