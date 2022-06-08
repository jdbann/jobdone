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

type Params struct {
	Slots []Slot
}

func New(params Params) tea.Model {
	return Stack{
		slots: params.Slots,
	}
}

func (m Stack) Init() tea.Cmd {
	var slotCmds []tea.Cmd

	for _, slot := range m.slots {
		cmd := slot.model.Init()
		slotCmds = append(slotCmds, cmd)
	}

	return tea.Batch(slotCmds...)
}

func (m Stack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg == nil {
		return m, nil
	}

	var cmd tea.Cmd
	var slotCmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		remainingSlots, remainingHeight := float64(len(m.slots)), float64(msg.Height)

		for i, slot := range m.slots {
			if !slot.fixedSize {
				continue
			}

			slotHeight := lipgloss.Height(slot.model.View())
			m.slots[i].model, cmd = slot.model.Update(tea.WindowSizeMsg{Height: slotHeight, Width: msg.Width})
			remainingSlots--
			remainingHeight -= float64(slotHeight)
			slotCmds = append(slotCmds, cmd)
		}

		for i, slot := range m.slots {
			if slot.fixedSize {
				continue
			}

			slotHeight := int(math.Round(remainingHeight / remainingSlots))
			m.slots[i].model, cmd = slot.model.Update(tea.WindowSizeMsg{Height: slotHeight, Width: msg.Width})
			remainingSlots--
			remainingHeight -= float64(slotHeight)
			slotCmds = append(slotCmds, cmd)
		}

		return m, tea.Batch(slotCmds...)
	}

	for i, slot := range m.slots {
		m.slots[i].model, cmd = slot.model.Update(msg)
		slotCmds = append(slotCmds, cmd)
	}

	return m, tea.Batch(slotCmds...)
}

func (m Stack) View() string {
	var views []string
	for _, slot := range m.slots {
		views = append(views, slot.model.View())
	}
	return lipgloss.JoinVertical(lipgloss.Left, views...)
}
