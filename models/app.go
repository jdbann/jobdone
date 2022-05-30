package models

import tea "github.com/charmbracelet/bubbletea"

var _ tea.Model = App{}

type App struct {
	splash tea.Model
}

type AppParams struct {
	Splash tea.Model
}

func NewApp(params AppParams) App {
	if params.Splash == nil {
		params.Splash = NewSplash(SplashParams{})
	}

	return App{
		splash: params.Splash,
	}
}

func (a App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg == nil {
		return a, nil
	}

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

func (a App) View() string {
	return a.splash.View()
}
