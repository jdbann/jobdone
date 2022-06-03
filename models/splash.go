package models

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jobdone.emailaddress.horse/utils/colors"
)

const (
	splashTitle = "" +
		"         _/    _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/\n" +
		"        _/  _/    _/  _/    _/      _/    _/  _/    _/  _/_/    _/  _/       \n" +
		"       _/  _/    _/  _/_/_/        _/    _/  _/    _/  _/  _/  _/  _/_/_/    \n" +
		"_/    _/  _/    _/  _/    _/      _/    _/  _/    _/  _/    _/_/  _/         \n" +
		" _/_/      _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/    "

	splashSubtitle = "You write the app, we'll do the hard work."

	splashDuration = time.Second * 2
)

var _ tea.Model = Splash{}

type Splash struct {
	logger          *zap.Logger
	height, width   int
	title, subtitle string
	duration        time.Duration
}

type SplashParams struct {
	Logger          *zap.Logger
	Title, Subtitle string
	Duration        time.Duration
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

	if params.Duration == 0 {
		params.Duration = splashDuration
	}

	return Splash{
		logger:   logger,
		title:    params.Title,
		subtitle: params.Subtitle,
		duration: params.Duration,
	}
}

func (s Splash) Init() tea.Cmd {
	s.logger.Debug("Initialised")
	return DismissSplashCmd(s.duration)
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

func DismissSplashCmd(duration time.Duration) func() tea.Msg {
	return func() tea.Msg {
		time.Sleep(duration)
		return SplashCompleteMsg{}
	}
}

type SplashCompleteMsg struct{}

func (msg SplashCompleteMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "SplashCompleteMsg")
	return nil
}
