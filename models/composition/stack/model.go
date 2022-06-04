package stack

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ tea.Model = Stack{}

type Stack struct {
	height, width int
	slots         []Slot
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
		m.height, m.width = msg.Height, msg.Width

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
	return lipgloss.NewStyle().Height(m.height).Width(m.width).Render("")
}
