package stack_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/composition/stack"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestInit(t *testing.T) {
	m := stack.New(stack.Params{})
	cmd := m.Init()
	teast.AssertCmdsEqual(t, nil, cmd)
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name               string
		slotModelsOptions  [][]teast.FakeModelOption
		msg                tea.Msg
		wantCmd            tea.Cmd
		wantView           string
		wantSlotModelsMsgs [][]tea.Msg
	}{
		{
			name:               "nil messages are skipped",
			slotModelsOptions:  [][]teast.FakeModelOption{{}, {}},
			msg:                nil,
			wantCmd:            nil,
			wantView:           "",
			wantSlotModelsMsgs: [][]tea.Msg{{}, {}},
		},
		{
			name:               "window resize adjusts view sizing",
			slotModelsOptions:  [][]teast.FakeModelOption{},
			msg:                tea.WindowSizeMsg{Height: 2, Width: 2},
			wantCmd:            nil,
			wantView:           "  \n  ",
			wantSlotModelsMsgs: [][]tea.Msg{},
		},
		{
			name:              "window resize passes distributed sizes to slot models",
			slotModelsOptions: [][]teast.FakeModelOption{{}, {}, {}},
			msg:               tea.WindowSizeMsg{Height: 4, Width: 2},
			wantCmd:           nil,
			wantView:          "  \n  \n  \n  ",
			wantSlotModelsMsgs: [][]tea.Msg{
				{tea.WindowSizeMsg{Height: 1, Width: 2}},
				{tea.WindowSizeMsg{Height: 2, Width: 2}},
				{tea.WindowSizeMsg{Height: 1, Width: 2}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var slotModels []teast.FakeModel
			var slots []stack.Slot
			for _, options := range tt.slotModelsOptions {
				model := teast.NewFakeModel(options...)
				slotModels = append(slotModels, model)
				slots = append(slots, stack.Slot{
					Model: model,
				})
			}
			m := stack.New(stack.Params{
				Slots: slots,
			})
			m, cmd := m.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, m.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
			for i := range slotModels {
				teast.AssertMsgsEqual(t, tt.wantSlotModelsMsgs[i], slotModels[i].Msgs())
			}
		})
	}
}
