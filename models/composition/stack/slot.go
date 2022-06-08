package stack

import tea "github.com/charmbracelet/bubbletea"

type Slot struct {
	model     tea.Model
	fixedSize bool
}

func FixedSlot(m tea.Model) Slot {
	return Slot{
		model:     m,
		fixedSize: true,
	}
}

func FlexiSlot(m tea.Model) Slot {
	return Slot{
		model:     m,
		fixedSize: false,
	}
}
