package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/challenge"
	"jobdone.emailaddress.horse/models/composition/box"
	"jobdone.emailaddress.horse/models/composition/stack"
	"jobdone.emailaddress.horse/models/healthcheck"
	"jobdone.emailaddress.horse/models/splash"
	"jobdone.emailaddress.horse/utils/colors"
	"jobdone.emailaddress.horse/utils/logger"
)

var _ tea.Model = App{}

type App struct {
	showSplash bool

	challenge tea.Model
	splash    tea.Model

	logger *zap.Logger
}

type AppParams struct {
	Challenge tea.Model
	Splash    tea.Model

	Logger *zap.Logger
}

func NewApp(params AppParams) App {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("App")

	if params.Challenge == nil {
		params.Challenge = stack.New(stack.Params{
			Slots: []stack.Slot{
				stack.FlexiSlot(box.New(box.Params{
					Model: challenge.New(challenge.Params{
						Logger: logger.Named("Stack").Named("Slot"),
					}),
					Style: lipgloss.NewStyle().
						Background(colors.Indigo1).
						Border(lipgloss.NormalBorder(), true).
						BorderForeground(colors.Indigo6).
						BorderBackground(colors.Indigo1).
						Padding(0, 1),
					Logger: logger.Named("Stack"),
				})),
				stack.FixedSlot(healthcheck.New(healthcheck.Params{
					Logger: logger.Named("Stack"),
				})),
			},
			Logger: logger,
		})
	}

	if params.Splash == nil {
		params.Splash = splash.New(splash.Params{
			Logger: logger,
		})
	}

	return App{
		showSplash: true,

		challenge: params.Challenge,
		splash:    params.Splash,

		logger: logger,
	}
}

func (a App) Init() tea.Cmd {
	a.logger.Debug("Initialised")
	return tea.Batch(
		a.challenge.Init(),
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
				zap.Object("tea.Msg", logger.KeyMsg(msg)),
			)
			return a, tea.Quit
		}

	case splash.CompleteMsg:
		a.showSplash = false
		a.logger.Debug(
			"Received splash complete message",
			zap.Object("tea.Msg", msg),
		)
	}

	// Pass message to nested models
	var cmd tea.Cmd
	a.challenge, cmd = a.challenge.Update(msg)

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

	return a.challenge.View()
}
