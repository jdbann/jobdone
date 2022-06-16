package challenge_test

import (
	"net/http"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/challenge"
	"jobdone.emailaddress.horse/models/healthcheck"
)

func TestChallenge1(t *testing.T) {
	tests := []struct {
		name         string
		msgs         []tea.Msg
		wantComplete bool
	}{
		{
			name:         "starts incomplete",
			msgs:         nil,
			wantComplete: false,
		},
		{
			name: "successful healthcheck completes",
			msgs: []tea.Msg{
				healthcheck.ResponseMsg{StatusCode: http.StatusOK},
			},
			wantComplete: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := challenge.Challenge1

			for _, msg := range tt.msgs {
				c = c.Update(msg)
			}

			if tt.wantComplete != c.Complete() {
				t.Errorf("expected %t; got %t", tt.wantComplete, c.Complete())
			}
		})
	}
}
