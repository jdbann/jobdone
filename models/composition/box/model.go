package box

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/utils/logger"
)

var _ tea.Model = Box{}

type Box struct {
	height, width int
	style         lipgloss.Style

	logger *zap.Logger
}

type Params struct {
	Style lipgloss.Style

	Logger *zap.Logger
}

func New(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Box")

	return Box{
		style: params.Style,

		logger: logger,
	}
}

func (m Box) Init() tea.Cmd {
	return nil
}

func (m Box) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)

		m.height, m.width = msg.Height, msg.Width
	}

	return m, nil
}

func (m Box) View() string {
	return m.style.Copy().Height(m.height).Width(m.width).Render("")
}
