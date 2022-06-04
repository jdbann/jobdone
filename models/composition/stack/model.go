package stack

import tea "github.com/charmbracelet/bubbletea"

var _ tea.Model = Stack{}

type Stack struct{}

type Params struct{}

func New(params Params) tea.Model {
	return Stack{}
}

func (m Stack) Init() tea.Cmd {
	return nil
}

func (m Stack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Stack) View() string {
	return ""
}
