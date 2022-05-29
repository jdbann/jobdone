package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
)

func main() {
	p := tea.NewProgram(newApp())
	if err := p.Start(); err != nil {
		fmt.Printf("OH NO! There has been an error: %v", err)
		os.Exit(1)
	}
}

var _ tea.Model = app{}

type app struct {
	splash models.Splash
}

func newApp() app {
	return app{
		splash: models.NewSplash(),
	}
}

func (a app) Init() tea.Cmd {
	return nil
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	}

	return a, nil
}

func (a app) View() string {
	return a.splash.View()
}
