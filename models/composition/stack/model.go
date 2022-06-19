package stack

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/utils/logger"
)

var _ tea.Model = Stack{}

type Stack struct {
	slots       []Slot
	distributor distributor

	logger *zap.Logger
}

type Params struct {
	Slots []Slot

	Logger *zap.Logger
}

func NewVertical(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Stack")

	return Stack{
		slots:       params.Slots,
		distributor: HeightDistributor{},

		logger: logger,
	}
}

func NewHorizontal(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Stack")

	return Stack{
		slots:       params.Slots,
		distributor: WidthDistributor{},

		logger: logger,
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
		m.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)

		remainingSlots, remainingSize := float64(len(m.slots)), float64(m.distributor.availableSize(msg))

		for i, slot := range m.slots {
			if !slot.fixedSize {
				continue
			}

			slotSize := m.distributor.slotSize(slot)
			m.slots[i].model, cmd = m.distributor.updateSlot(msg, slot.model, slotSize)
			remainingSlots--
			remainingSize -= float64(slotSize)
			slotCmds = append(slotCmds, cmd)
		}

		for i, slot := range m.slots {
			if slot.fixedSize {
				continue
			}

			slotSize := int(math.Round(remainingSize / remainingSlots))
			m.slots[i].model, cmd = m.distributor.updateSlot(msg, slot.model, slotSize)
			remainingSlots--
			remainingSize -= float64(slotSize)
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
	return m.distributor.joinViews(views)
}
