package challenge_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/challenge"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestInit(t *testing.T) {
	m := challenge.New(challenge.Params{})
	teast.AssertCmdsEqual(t, nil, m.Init())
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
			wantView: "No active challenge.",
		},
		{
			name: "challenge changed message sets view",
			msg: challenge.ChangedMsg{
				Number:      999,
				Title:       "Test Challenge",
				Description: "A test challenge.",
			},
			wantCmd: nil,
			wantView: `Challenge #999: Test Challenge

A test challenge.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := challenge.New(challenge.Params{})
			m, cmd := m.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, m.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}
