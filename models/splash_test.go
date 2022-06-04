package models_test

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestSplash_Init(t *testing.T) {
	splash := models.NewSplash(models.SplashParams{
		Duration: time.Nanosecond * 1,
	})

	cmd := splash.Init()

	teast.AssertCmdsEqual(t, models.DismissSplashCmd(0), cmd)
}

func TestSplash_Update(t *testing.T) {
	tests := []struct {
		name     string
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name:    "nil messages are skipped",
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
			name:    "window resize adjusts view sizing",
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

			teast.AssertViewsEqual(t, tt.wantView, updatedSplash.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}
