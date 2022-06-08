package box

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = Box{}

type Box struct {
	height, width int
	style         lipgloss.Style
}

type Params struct {
	Style lipgloss.Style
}

func New(params Params) tea.Model {
	return Box{
		style: params.Style,
	}
}

func (m Box) Init() tea.Cmd {
	return nil
}

func (m Box) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
	}

	return m, nil
}

func (m Box) View() string {
	return m.style.Copy().Height(m.height).Width(m.width).Render("")
}
