package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

var _ tea.Model = App{}

type App struct {
	logger *zap.Logger
	splash tea.Model
}

type AppParams struct {
	Logger *zap.Logger
	Splash tea.Model
}

func NewApp(params AppParams) App {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("App")

	if params.Splash == nil {
		params.Splash = NewSplash(SplashParams{
			Logger: logger,
		})
	}

	return App{
		logger: logger,
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
			a.logger.Debug(
				"Received quit message",
				zap.Object("tea.Msg", keyMsg(msg)),
			)
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
