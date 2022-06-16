package objective_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/challenge/objective"
)

func TestSimpleVerifier(t *testing.T) {
	boolCheckFn := func(msg tea.Msg) bool {
		return msg.(bool)
	}

	tests := []struct {
		name         string
		msgs         []tea.Msg
		wantComplete bool
	}{
		{
			name:         "complete if checkFn returns true",
			msgs:         []tea.Msg{true},
			wantComplete: true,
		},
		{
			name:         "incomplete if checkFn returns false",
			msgs:         []tea.Msg{false},
			wantComplete: false,
		},
		{
			name:         "complete if checkFn already returned true",
			msgs:         []tea.Msg{true, false},
			wantComplete: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := objective.New(objective.Params{
				Verifier: objective.NewSimpleVerifier(boolCheckFn),
			})

			for _, msg := range tt.msgs {
				m = m.Update(msg)
			}

			if tt.wantComplete != m.Complete() {
				t.Errorf("expected %t; got %t", tt.wantComplete, m.Complete())
			}
		})
	}
}
