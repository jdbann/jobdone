package challenge_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/challenge"
	"jobdone.emailaddress.horse/models/challenge/objective"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestInit(t *testing.T) {
	m := challenge.New(challenge.Params{})
	teast.AssertCmdsEqual(t, challenge.SwitchCmd(challenge.Challenge1), m.Init())
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		params   challenge.Params
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name:     "nil messages are skipped",
			params:   challenge.Params{},
			msg:      nil,
			wantCmd:  nil,
			wantView: "No active challenge.",
		},
		{
			name:   "challenge changed message sets view",
			params: challenge.Params{},
			msg: challenge.ChangedMsg{
				Challenge: challenge.Definition{
					Number:      999,
					Title:       "Test Challenge",
					Description: "A test challenge.",
					Objectives: []objective.Objective{
						objective.New(objective.Params{
							Description: "A test objective",
						}),
					},
				},
			},
			wantCmd: nil,
			wantView: `Challenge #999: Test Challenge
--------
A test challenge.
[ ] A test objective`,
		},
		{
			name: "messages are passed to objectives for verifying completion",
			params: challenge.Params{
				Challenge: challenge.Definition{
					Number:      999,
					Title:       "Test Challenge",
					Description: "A test challenge.",
					Objectives: []objective.Objective{
						objective.New(objective.Params{
							Description: "A test objective",
							Verifier: objective.NewSimpleVerifier(func(_ tea.Msg) bool {
								return true
							}),
						}),
					},
				},
			},
			msg:     teast.FakeMsg{},
			wantCmd: nil,
			wantView: `Challenge #999: Test Challenge
--------
A test challenge.
[✓] A test objective`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := challenge.New(tt.params)
			m, cmd := m.Update(tt.msg)

			teast.AssertViewsContentEqual(t, tt.wantView, m.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}
