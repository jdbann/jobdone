package stack

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = Stack{}

type Stack struct {
	slots []Slot
}

type Slot struct {
	Model tea.Model
}

type Params struct {
	Slots []Slot
}

func New(params Params) tea.Model {
	return Stack{
		slots: params.Slots,
	}
}

func (m Stack) Init() tea.Cmd {
	return nil
}

func (m Stack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var slotCmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		remainingSlots, remainingHeight := float64(len(m.slots)), float64(msg.Height)
		for i, slot := range m.slots {
			slotHeight := int(math.Round(remainingHeight / remainingSlots))
			m.slots[i].Model, cmd = slot.Model.Update(tea.WindowSizeMsg{Height: slotHeight, Width: msg.Width})
			remainingSlots--
			remainingHeight -= float64(slotHeight)
			slotCmds = append(slotCmds, cmd)
		}

		return m, tea.Batch(slotCmds...)
	}

	return m, nil
}

func (m Stack) View() string {
	var views []string
	for _, slot := range m.slots {
		views = append(views, slot.Model.View())
	}
	return lipgloss.JoinVertical(lipgloss.Left, views...)
}
