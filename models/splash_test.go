package models_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
)

func TestSplash_Init(t *testing.T) {
	splash := models.NewSplash(models.SplashParams{})

	cmd := splash.Init()

	if cmd != nil {
		t.Errorf("Expected %#v, got %#v", nil, cmd)
	}
}

func TestSplash_Update(t *testing.T) {
	tests := []struct {
		name     string
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name:    "nil",
			msg:     nil,
			wantCmd: nil,
			wantView: "" +
				"     \n" +
				"     \n" +
				"Title\n" +
				"     \n" +
				"     \n" +
				"Subtitle\n" +
				"        \n" +
				"        ",
		},
		{
			name:    "tea.WindowSizeMsg",
			msg:     tea.WindowSizeMsg{Height: 10, Width: 10},
			wantCmd: nil,
			wantView: "" +
				"          \n" +
				"          \n" +
				"          \n" +
				"  Title   \n" +
				"          \n" +
				"          \n" +
				" Subtitle \n" +
				"          \n" +
				"          \n" +
				"          ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splash := models.NewSplash(models.SplashParams{Title: "Title", Subtitle: "Subtitle"})
			updatedSplash, cmd := splash.Update(tt.msg)
			view := stripAnsi(updatedSplash.View())

			if view != tt.wantView {
				t.Errorf("\nExpected: %q\nGot:      %q", tt.wantView, view)
			}

			if cmd == nil && tt.wantCmd == nil {
				return
			}

			if (cmd != nil && tt.wantCmd != nil) && (tt.wantCmd() != cmd()) {
				t.Errorf("\nExpected: %#v\nGot:      %#v", tt.wantCmd, cmd)
			} else {
				t.Errorf("\nExpected: %#v\nGot:      %#v", tt.wantCmd, cmd)
			}
		})
	}
}
