package app_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/app"
	"jobdone.emailaddress.horse/models/splash"
	"jobdone.emailaddress.horse/pkg/teast"
)

func TestApp_Init(t *testing.T) {
	model := app.New(app.Params{
		Challenge: teast.NewFakeModel(t, teast.InitReturns(teast.FakeCmd("Challenge init"))),
		Splash:    teast.NewFakeModel(t, teast.InitReturns(teast.FakeCmd("Splash init"))),
	})

	cmd := model.Init()

	teast.AssertCmdsEqual(t, tea.Batch(teast.FakeCmd("Challenge init"), teast.FakeCmd("Splash init")), cmd)
}

func TestApp_Update(t *testing.T) {
	tests := []struct {
		name             string
		challengeOptions []teast.FakeModelOption
		splashOptions    []teast.FakeModelOption
		msg              tea.Msg
		wantCmd          tea.Cmd
		wantView         string
	}{
		{
			name:             "nil messages are skipped",
			challengeOptions: []teast.FakeModelOption{teast.ViewReturns("Fake Challenge")},
			splashOptions:    []teast.FakeModelOption{teast.ViewReturns("Fake Splash")},
			msg:              nil,
			wantCmd:          nil,
			wantView:         "Fake Splash",
		},
		{
			name:             "ctrl+c quits immediately",
			challengeOptions: []teast.FakeModelOption{teast.ViewReturns("Fake Challenge")},
			splashOptions:    []teast.FakeModelOption{teast.ViewReturns("Fake Splash")},
			msg: tea.KeyMsg{
				Type: tea.KeyCtrlC,
			},
			wantCmd:  tea.Quit,
			wantView: "Fake Splash",
		},
		{
			name:             "q quits immediately",
			challengeOptions: []teast.FakeModelOption{teast.ViewReturns("Fake Challenge")},
			splashOptions:    []teast.FakeModelOption{teast.ViewReturns("Fake Splash")},
			msg: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune{'q'},
			},
			wantCmd:  tea.Quit,
			wantView: "Fake Splash",
		},
		{
			name: "splash complete switches to challenge",
			challengeOptions: []teast.FakeModelOption{
				teast.ViewReturns("Fake Challenge"),
				teast.ExpectMessages(splash.CompleteMsg{}),
			},
			splashOptions: []teast.FakeModelOption{teast.ViewReturns("Fake Splash")},
			msg:           splash.CompleteMsg{},
			wantCmd:       nil,
			wantView:      "Fake Challenge",
		},
		{
			name: "messages are passed to nested models and commands returned",
			challengeOptions: []teast.FakeModelOption{
				teast.UpdateReturns(teast.FakeCmd("Message from Challenge")),
				teast.ViewReturns("Fake Challenge"),
				teast.ExpectMessages(teast.FakeMsg{}),
			},
			splashOptions: []teast.FakeModelOption{
				teast.UpdateReturns(teast.FakeCmd("Message from Splash")),
				teast.ViewReturns("Fake Splash"),
				teast.ExpectMessages(teast.FakeMsg{}),
			},
			msg:      teast.FakeMsg{},
			wantCmd:  tea.Batch(teast.FakeCmd("Message from Challenge"), teast.FakeCmd("Message from Splash")),
			wantView: "Fake Splash",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeChallenge := teast.NewFakeModel(t, tt.challengeOptions...)
			fakeSplash := teast.NewFakeModel(t, tt.splashOptions...)
			model := app.New(app.Params{
				Challenge: fakeChallenge,
				Splash:    fakeSplash,
			})
			updatedApp, cmd := model.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, updatedApp.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}
