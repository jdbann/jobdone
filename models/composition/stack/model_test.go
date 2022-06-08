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
		name     string
		params   stack.Params
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name: "nil messages are skipped",
			params: stack.Params{
				Slots: []stack.Slot{stack.FlexiSlot(teast.NewFakeModel(t))},
			},
			msg:      nil,
			wantCmd:  nil,
			wantView: "",
		},
		{
			name: "window resize passes distributed sizes to slot models",
			params: stack.Params{
				Slots: []stack.Slot{
					stack.FlexiSlot(teast.NewFakeModel(t, teast.ExpectMessages(tea.WindowSizeMsg{Height: 1, Width: 2}))),
					stack.FlexiSlot(teast.NewFakeModel(t, teast.ExpectMessages(tea.WindowSizeMsg{Height: 2, Width: 2}))),
					stack.FlexiSlot(teast.NewFakeModel(t, teast.ExpectMessages(tea.WindowSizeMsg{Height: 1, Width: 2}))),
				},
			},
			msg:      tea.WindowSizeMsg{Height: 4, Width: 2},
			wantCmd:  nil,
			wantView: "\n\n",
		},
		{
			name: "window resize respects fixed size slots",
			params: stack.Params{
				Slots: []stack.Slot{
					stack.FlexiSlot(teast.NewFakeModel(t, teast.ExpectMessages(tea.WindowSizeMsg{Height: 3, Width: 2}))),
					stack.FixedSlot(teast.NewFakeModel(t, teast.ExpectMessages(tea.WindowSizeMsg{Height: 1, Width: 2}), teast.ViewReturns("one fixed line"))),
					stack.FlexiSlot(teast.NewFakeModel(t, teast.ExpectMessages(tea.WindowSizeMsg{Height: 3, Width: 2}))),
				},
			},
			msg:      tea.WindowSizeMsg{Height: 7, Width: 2},
			wantCmd:  nil,
			wantView: "              \none fixed line\n              ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := stack.New(tt.params)
			m, cmd := m.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, m.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}

func TestView(t *testing.T) {
	m := stack.New(stack.Params{
		Slots: []stack.Slot{
			stack.FlexiSlot(teast.NewFakeModel(t, teast.ViewReturns("Slot 1 View"))),
			stack.FlexiSlot(teast.NewFakeModel(t, teast.ViewReturns("Slot 2 View"))),
		},
	})
	wantView := "" +
		"Slot 1 View\n" +
		"Slot 2 View"
	teast.AssertViewsEqual(t, wantView, m.View())
}
