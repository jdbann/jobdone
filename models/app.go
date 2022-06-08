package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/composition/box"
	"jobdone.emailaddress.horse/models/composition/stack"
	"jobdone.emailaddress.horse/utils/colors"
)

var _ tea.Model = App{}

type App struct {
	showSplash bool

	healthcheck tea.Model
	splash      tea.Model

	logger *zap.Logger
}

type AppParams struct {
	Healthcheck tea.Model
	Splash      tea.Model

	Logger *zap.Logger
}

func NewApp(params AppParams) App {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("App")

	if params.Healthcheck == nil {
		params.Healthcheck = stack.New(stack.Params{
			Slots: []stack.Slot{
				stack.FlexiSlot(box.New(box.Params{
					Style: lipgloss.NewStyle().Background(colors.Indigo1),
				})),
				stack.FixedSlot(NewHealthcheck(HealthcheckParams{
					Logger: logger,
				})),
			},
		})
	}

	if params.Splash == nil {
		params.Splash = NewSplash(SplashParams{
			Logger: logger,
		})
	}

	return App{
		showSplash: true,

		healthcheck: params.Healthcheck,
		splash:      params.Splash,

		logger: logger,
	}
}

func (a App) Init() tea.Cmd {
	a.logger.Debug("Initialised")
	return tea.Batch(
		a.healthcheck.Init(),
		a.splash.Init(),
	)
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

	case SplashCompleteMsg:
		a.showSplash = false
		a.logger.Debug(
			"Received splash complete message",
			zap.Object("tea.Msg", msg),
		)
	}

	// Pass message to nested models
	var cmd tea.Cmd
	a.healthcheck, cmd = a.healthcheck.Update(msg)

	if a.showSplash {
		var splashCmd tea.Cmd
		a.splash, splashCmd = a.splash.Update(msg)
		cmd = tea.Batch(cmd, splashCmd)
	}

	return a, cmd
}

func (a App) View() string {
	if a.showSplash {
		return a.splash.View()
	}

	return a.healthcheck.View()
}
