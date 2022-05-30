package models_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
)

func TestApp_Init(t *testing.T) {
	app := models.NewApp(models.AppParams{})

	cmd := app.Init()

	assertCmdsEqual(t, nil, cmd)
}

func TestApp_Update(t *testing.T) {
	tests := []struct {
		name           string
		splash         tea.Model
		msg            tea.Msg
		wantCmd        tea.Cmd
		wantView       string
		wantSplashMsgs []tea.Msg
	}{
		{
			name:           "nil messages are skipped",
			splash:         fakeModel(),
			msg:            nil,
			wantCmd:        nil,
			wantView:       "Fake Model",
			wantSplashMsgs: []tea.Msg{},
		},
		{
			name:   "ctrl+c quits immediately",
			splash: fakeModel(),
			msg: tea.KeyMsg{
				Type: tea.KeyCtrlC,
			},
			wantCmd:        tea.Quit,
			wantView:       "Fake Model",
			wantSplashMsgs: []tea.Msg{},
		},
		{
			name:   "q quits immediately",
			splash: fakeModel(),
			msg: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{'q'},
			},
			wantCmd:        tea.Quit,
			wantView:       "Fake Model",
			wantSplashMsgs: []tea.Msg{},
		},
		{
			name:           "messages are passed to nested Splash model and commands returned",
			splash:         fakeModel(updateReturns(fakeCmd("Message from Splash"))),
			msg:            fakeMsg{},
			wantCmd:        fakeCmd("Message from Splash"),
			wantView:       "Fake Model",
			wantSplashMsgs: []tea.Msg{fakeMsg{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeSplash := &_fakeModel{view: "Fake Model"}
			app := models.NewApp(models.AppParams{Splash: fakeSplash})
			updatedApp, cmd := app.Update(tt.msg)
			view := stripAnsi(updatedApp.View())

			assertViewsEqual(t, tt.wantView, view)
			assertCmdsEqual(t, tt.wantCmd, cmd)
			assertMsgsEqual(t, tt.wantSplashMsgs, fakeSplash.msgs)
		})
	}
}
