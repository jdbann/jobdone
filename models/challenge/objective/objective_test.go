package objective_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/challenge/objective"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestUpdate(t *testing.T) {
	boolVerifier := objective.NewSimpleVerifier(func(msg tea.Msg) bool {
		return msg.(bool)
	})

	tests := []struct {
		name         string
		params       objective.Params
		msgs         tea.Msg
		wantComplete bool
	}{
		{
			name: "nil verifier doesn't complete",
			params: objective.Params{
				Verifier: nil,
			},
			msgs:         teast.FakeMsg{},
			wantComplete: false,
		},
		{
			name: "objective complete if verifier returns true",
			params: objective.Params{
				Verifier: boolVerifier,
			},
			msgs:         true,
			wantComplete: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := objective.New(tt.params)

			o = o.Update(tt.msgs)

			if tt.wantComplete != o.Complete() {
				t.Errorf("expected %t; got %t", tt.wantComplete, o.Complete())
			}
		})
	}
}
