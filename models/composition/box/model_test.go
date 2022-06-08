package box_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/composition/box"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestInit(t *testing.T) {
	m := box.New(box.Params{})
	cmd := m.Init()
	teast.AssertCmdsEqual(t, nil, cmd)
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name:     "nil messages are skipped",
			msg:      nil,
			wantCmd:  nil,
			wantView: "",
		},
		{
			name:     "window resize changes view size",
			msg:      tea.WindowSizeMsg{Height: 2, Width: 2},
			wantCmd:  nil,
			wantView: "  \n  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := box.New(box.Params{})
			m, cmd := m.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, m.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}
