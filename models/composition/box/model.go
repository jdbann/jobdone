package box

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/utils/logger"
)

var _ tea.Model = Box{}

type Box struct {
	model         tea.Model
	height, width int
	style         lipgloss.Style

	logger *zap.Logger
}

type Params struct {
	Model tea.Model
	Style lipgloss.Style

	Logger *zap.Logger
}

func New(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Box")

	return Box{
		model: params.Model,
		style: params.Style,

		logger: logger,
	}
}

func (m Box) Init() tea.Cmd {
	if m.model == nil {
		return nil
	}

	return m.model.Init()
}

func (m Box) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)

		m.height, m.width = msg.Height, msg.Width

		if m.model == nil {
			return m, nil
		}

		modelMsg := tea.WindowSizeMsg{Height: m.contentHeight(), Width: m.contentWidth()}
		m.model, cmd = m.model.Update(modelMsg)

		return m, cmd
	}

	if m.model != nil {
		m.model, cmd = m.model.Update(msg)
	}

	return m, cmd
}

func (m Box) View() string {
	var content string
	if m.model != nil {
		content = m.model.View()
	}

	return m.style.Copy().
		Height(m.contentHeight() + m.style.GetVerticalPadding()).
		Width(m.contentWidth() + m.style.GetHorizontalPadding()).
		Render(content)
}

func (m Box) contentHeight() int {
	return m.height - m.style.GetVerticalFrameSize()
}

func (m Box) contentWidth() int {
	return m.width - m.style.GetHorizontalFrameSize()
}
