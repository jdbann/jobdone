package splash

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/utils/colors"
	"jobdone.emailaddress.horse/utils/logger"
)

const (
	title = "" +
		"         _/    _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/\n" +
		"        _/  _/    _/  _/    _/      _/    _/  _/    _/  _/_/    _/  _/       \n" +
		"       _/  _/    _/  _/_/_/        _/    _/  _/    _/  _/  _/  _/  _/_/_/    \n" +
		"_/    _/  _/    _/  _/    _/      _/    _/  _/    _/  _/    _/_/  _/         \n" +
		" _/_/      _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/    "

	subtitle = "Let's get to work..."

	duration = time.Millisecond * 2500

	fps = 60
)

var (
	spring = harmonica.NewSpring(harmonica.FPS(fps), 5.0, 1.0)
)

var _ tea.Model = Splash{}

type Splash struct {
	logger               *zap.Logger
	height, width        int
	title, subtitle      string
	springPos, springVel float64
	duration             time.Duration
}

type Params struct {
	Logger          *zap.Logger
	Title, Subtitle string
	Duration        time.Duration
}

func New(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Splash")

	if params.Title == "" {
		params.Title = title
	}

	if params.Subtitle == "" {
		params.Subtitle = subtitle
	}

	if params.Duration == 0 {
		params.Duration = duration
	}

	return Splash{
		logger:    logger,
		title:     params.Title,
		subtitle:  params.Subtitle,
		springPos: 0,
		springVel: 3,
		duration:  params.Duration,
	}
}

func (s Splash) Init() tea.Cmd {
	s.logger.Debug("Initialised")
	return tea.Batch(DismissCmd(s.duration), AnimateCmd())
}

func (s Splash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)
		s.height, s.width = msg.Height, msg.Width

	case TickMsg:
		s.logger.Debug(
			"Received splash tick message",
			zap.Object("tea.Msg", msg),
		)
		s.springPos, s.springVel = spring.Update(s.springPos, s.springVel, 0.5)

		return s, AnimateCmd()
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
		lipgloss.Position(s.springPos),
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
		lipgloss.Position(1.0-s.springPos),
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
