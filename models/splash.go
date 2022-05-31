package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/utils/colors"
)

const splashTitle = "" +
	"         _/    _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/\n" +
	"        _/  _/    _/  _/    _/      _/    _/  _/    _/  _/_/    _/  _/       \n" +
	"       _/  _/    _/  _/_/_/        _/    _/  _/    _/  _/  _/  _/  _/_/_/    \n" +
	"_/    _/  _/    _/  _/    _/      _/    _/  _/    _/  _/    _/_/  _/         \n" +
	" _/_/      _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/    "

const splashSubtitle = "You write the app, we'll do the hard work."

var _ tea.Model = Splash{}

type Splash struct {
	logger          *zap.Logger
	height, width   int
	title, subtitle string
}

type SplashParams struct {
	Logger          *zap.Logger
	Title, Subtitle string
}

func NewSplash(params SplashParams) Splash {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Splash")

	if params.Title == "" {
		params.Title = splashTitle
	}

	if params.Subtitle == "" {
		params.Subtitle = splashSubtitle
	}

	return Splash{
		logger:   logger,
		title:    params.Title,
		subtitle: params.Subtitle,
	}
}

func (s Splash) Init() tea.Cmd {
	return nil
}

func (s Splash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", windowSizeMsg(msg)),
		)
		s.height, s.width = msg.Height, msg.Width
	}

	return s, nil
}

func (s Splash) View() string {
	styledTitle := lipgloss.NewStyle().
		Foreground(colors.Indigo12).
		Background(colors.Indigo2).
		Padding(2, 0, 2, 0).
		Render(s.title)

	titleBar := lipgloss.PlaceHorizontal(
		s.width,
		lipgloss.Center,
		styledTitle,
		lipgloss.WithWhitespaceBackground(colors.Indigo2),
	)

	styledSubtitle := lipgloss.NewStyle().
		Foreground(colors.Indigo11).
		Background(colors.Indigo2).
		Padding(0, 0, 2, 0).
		Render(s.subtitle)

	subtitleBar := lipgloss.PlaceHorizontal(
		s.width,
		lipgloss.Center,
		styledSubtitle,
		lipgloss.WithWhitespaceBackground(colors.Indigo2),
	)

	return lipgloss.PlaceVertical(
		s.height,
		lipgloss.Center,
		titleBar+"\n"+subtitleBar,
		lipgloss.WithWhitespaceBackground(colors.Indigo1),
	)
}
