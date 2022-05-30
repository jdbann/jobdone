package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
)

func main() {
	p := tea.NewProgram(newApp(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("OH NO! There has been an error: %v", err)
		os.Exit(1)
	}
}

var _ tea.Model = app{}

type app struct {
	splash tea.Model
}

func newApp() app {
	return app{
		splash: models.NewSplash(models.SplashParams{}),
	}
}

func (a app) Init() tea.Cmd {
	return nil
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle app messages
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	}

	// Pass message to nested models
	var cmd tea.Cmd
	a.splash, cmd = a.splash.Update(msg)

	return a, cmd
}

func (a app) View() string {
	return a.splash.View()
}
