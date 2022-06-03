package models_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
)

func TestApp_Init(t *testing.T) {
	app := models.NewApp(models.AppParams{
		Healthcheck: fakeModel(initReturns(fakeCmd("Healthcheck init"))),
		Splash:      fakeModel(initReturns(fakeCmd("Splash init"))),
	})

	cmd := app.Init()

	assertCmdsEqual(t, tea.Batch(fakeCmd("Healthcheck init"), fakeCmd("Splash init")), cmd)
}

func TestApp_Update(t *testing.T) {
	tests := []struct {
		name                string
		healthcheckOptions  []fakeModelOption
		splashOptions       []fakeModelOption
		msg                 tea.Msg
		wantCmd             tea.Cmd
		wantView            string
		wantHealthcheckMsgs []tea.Msg
		wantSplashMsgs      []tea.Msg
	}{
		{
			name:                "nil messages are skipped",
			healthcheckOptions:  []fakeModelOption{viewReturns("Fake Healthcheck")},
			splashOptions:       []fakeModelOption{viewReturns("Fake Splash")},
			msg:                 nil,
			wantCmd:             nil,
			wantView:            "Fake Splash",
			wantHealthcheckMsgs: []tea.Msg{},
			wantSplashMsgs:      []tea.Msg{},
		},
		{
			name:               "ctrl+c quits immediately",
			healthcheckOptions: []fakeModelOption{viewReturns("Fake Healthcheck")},
			splashOptions:      []fakeModelOption{viewReturns("Fake Splash")},
			msg: tea.KeyMsg{
				Type: tea.KeyCtrlC,
			},
			wantCmd:             tea.Quit,
			wantView:            "Fake Splash",
			wantHealthcheckMsgs: []tea.Msg{},
			wantSplashMsgs:      []tea.Msg{},
		},
		{
			name:               "q quits immediately",
			healthcheckOptions: []fakeModelOption{viewReturns("Fake Healthcheck")},
			splashOptions:      []fakeModelOption{viewReturns("Fake Splash")},
			msg: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{'q'},
			},
			wantCmd:             tea.Quit,
			wantView:            "Fake Splash",
			wantHealthcheckMsgs: []tea.Msg{},
			wantSplashMsgs:      []tea.Msg{},
		},
		{
			name:                "splash complete switches to healthcheck",
			healthcheckOptions:  []fakeModelOption{viewReturns("Fake Healthcheck")},
			splashOptions:       []fakeModelOption{viewReturns("Fake Splash")},
			msg:                 models.SplashCompleteMsg{},
			wantCmd:             nil,
			wantView:            "Fake Healthcheck",
			wantHealthcheckMsgs: []tea.Msg{models.SplashCompleteMsg{}},
			wantSplashMsgs:      []tea.Msg{},
		},
		{
			name:                "messages are passed to nested models and commands returned",
			healthcheckOptions:  []fakeModelOption{updateReturns(fakeCmd("Message from Healthcheck")), viewReturns("Fake Healthcheck")},
			splashOptions:       []fakeModelOption{updateReturns(fakeCmd("Message from Splash")), viewReturns("Fake Splash")},
			msg:                 fakeMsg{},
			wantCmd:             tea.Batch(fakeCmd("Message from Healthcheck"), fakeCmd("Message from Splash")),
			wantView:            "Fake Splash",
			wantHealthcheckMsgs: []tea.Msg{fakeMsg{}},
			wantSplashMsgs:      []tea.Msg{fakeMsg{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeHealthcheck := fakeModel(tt.healthcheckOptions...)
			fakeSplash := fakeModel(tt.splashOptions...)
			app := models.NewApp(models.AppParams{
				Healthcheck: fakeHealthcheck,
				Splash:      fakeSplash,
			})
			updatedApp, cmd := app.Update(tt.msg)
			view := stripAnsi(updatedApp.View())

			assertViewsEqual(t, tt.wantView, view)
			assertCmdsEqual(t, tt.wantCmd, cmd)
			assertMsgsEqual(t, tt.wantHealthcheckMsgs, fakeHealthcheck.msgs)
			assertMsgsEqual(t, tt.wantSplashMsgs, fakeSplash.msgs)
		})
	}
}
