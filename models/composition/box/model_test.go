package box_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/composition/box"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		params  func(*testing.T) box.Params
		wantCmd tea.Cmd
	}{
		{
			name:    "no model returns nil",
			params:  func(_ *testing.T) box.Params { return box.Params{} },
			wantCmd: nil,
		},
		{
			name: "model returns model's init cmd",
			params: func(t *testing.T) box.Params {
				return box.Params{
					Model: teast.NewFakeModel(t, teast.InitReturns(teast.FakeCmd("content"))),
				}
			},
			wantCmd: teast.FakeCmd("content"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := box.New(tt.params(t))
			cmd := m.Init()
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		params   func(*testing.T) box.Params
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name:     "nil messages are skipped",
			params:   func(_ *testing.T) box.Params { return box.Params{} },
			msg:      nil,
			wantCmd:  nil,
			wantView: "",
		},
		{
			name:     "window resize changes view size",
			params:   func(_ *testing.T) box.Params { return box.Params{} },
			msg:      tea.WindowSizeMsg{Height: 2, Width: 2},
			wantCmd:  nil,
			wantView: "  \n  ",
		},
		{
			name: "messages are passed to content model and commands returned",
			params: func(t *testing.T) box.Params {
				return box.Params{
					Model: teast.NewFakeModel(t, teast.ExpectMessages(teast.FakeMsg{}), teast.UpdateReturns(teast.FakeCmd("content")), teast.ViewReturns("Content")),
				}
			},
			msg:      teast.FakeMsg{},
			wantCmd:  teast.FakeCmd("content"),
			wantView: "Content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := box.New(tt.params(t))
			m, cmd := m.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, m.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}
