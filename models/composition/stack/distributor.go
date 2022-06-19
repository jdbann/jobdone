package stack

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type distributor interface {
	availableSize(msg tea.WindowSizeMsg) int
	slotSize(slot Slot) int
	updateSlot(msg tea.WindowSizeMsg, model tea.Model, size int) (tea.Model, tea.Cmd)
	joinViews(views []string) string
}

var _ distributor = HeightDistributor{}

type HeightDistributor struct{}

func (d HeightDistributor) availableSize(msg tea.WindowSizeMsg) int {
	return msg.Height
}

func (d HeightDistributor) slotSize(slot Slot) int {
	return lipgloss.Height(slot.model.View())
}

func (d HeightDistributor) updateSlot(msg tea.WindowSizeMsg, model tea.Model, size int) (tea.Model, tea.Cmd) {
	return model.Update(tea.WindowSizeMsg{Height: size, Width: msg.Width})
}

func (d HeightDistributor) joinViews(views []string) string {
	return lipgloss.JoinVertical(lipgloss.Left, views...)
}

var _ distributor = WidthDistributor{}

type WidthDistributor struct{}

func (d WidthDistributor) availableSize(msg tea.WindowSizeMsg) int {
	return msg.Width
}

func (d WidthDistributor) slotSize(slot Slot) int {
	return lipgloss.Width(slot.model.View())
}

func (d WidthDistributor) updateSlot(msg tea.WindowSizeMsg, model tea.Model, size int) (tea.Model, tea.Cmd) {
	return model.Update(tea.WindowSizeMsg{Height: msg.Height, Width: size})
}

func (d WidthDistributor) joinViews(views []string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}
