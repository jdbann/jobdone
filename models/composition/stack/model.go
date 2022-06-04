package stack

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = Stack{}

type Stack struct {
	height, width int
}

type Params struct{}

func New(params Params) tea.Model {
	return Stack{}
}

func (m Stack) Init() tea.Cmd {
	return nil
}

func (m Stack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
	}

	return m, nil
}

func (m Stack) View() string {
	return lipgloss.NewStyle().Height(m.height).Width(m.width).Render("")
}
