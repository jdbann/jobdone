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
		name              string
		slotModelsOptions [][]teast.FakeModelOption
		msg               tea.Msg
		wantCmd           tea.Cmd
		wantView          string
	}{
		{
			name:              "nil messages are skipped",
			slotModelsOptions: [][]teast.FakeModelOption{{}},
			msg:               nil,
			wantCmd:           nil,
			wantView:          "",
		},
		{
			name: "window resize passes distributed sizes to slot models",
			slotModelsOptions: [][]teast.FakeModelOption{{
				teast.ExpectMessages(tea.WindowSizeMsg{Height: 1, Width: 2}),
			}, {
				teast.ExpectMessages(tea.WindowSizeMsg{Height: 2, Width: 2}),
			}, {
				teast.ExpectMessages(tea.WindowSizeMsg{Height: 1, Width: 2}),
			}},
			msg:      tea.WindowSizeMsg{Height: 4, Width: 2},
			wantCmd:  nil,
			wantView: "\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var slotModels []tea.Model
			var slots []stack.Slot
			for _, options := range tt.slotModelsOptions {
				model := teast.NewFakeModel(t, options...)
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
		})
	}
}

func TestView(t *testing.T) {
	m := stack.New(stack.Params{
		Slots: []stack.Slot{
			{Model: teast.NewFakeModel(t, teast.ViewReturns("Slot 1 View"))},
			{Model: teast.NewFakeModel(t, teast.ViewReturns("Slot 2 View"))},
		},
	})
	wantView := "" +
		"Slot 1 View\n" +
		"Slot 2 View"
	teast.AssertViewsEqual(t, wantView, m.View())
}
