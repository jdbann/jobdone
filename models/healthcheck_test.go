package models_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
	"jobdone.emailaddress.horse/pkg/teast"
)

var fakeHealthcheckError = errors.New("fake error")

type fakeHealthcheckPerformer struct {
	statusCode int
	err        error
}

func (f fakeHealthcheckPerformer) Healthcheck() (*http.Response, error) {
	return &http.Response{StatusCode: f.statusCode}, f.err
}

func TestHealthcheck_Init(t *testing.T) {
	client := fakeHealthcheckPerformer{}

	healthcheck := models.NewHealthcheck(models.HealthcheckParams{
		Client: client,
	})

	cmd := healthcheck.Init()

	teast.AssertCmdsEqual(t, func() tea.Msg { return models.CheckHealthCmd(client)(time.Now()) }, cmd)
}

func TestHealthcheck_Update(t *testing.T) {
	tests := []struct {
		name     string
		client   fakeHealthcheckPerformer
		msg      tea.Msg
		wantCmd  tea.Cmd
		wantView string
	}{
		{
			name:     "nil messages are skipped",
			client:   fakeHealthcheckPerformer{},
			msg:      nil,
			wantCmd:  nil,
			wantView: " ◌  Error with server connection ",
		},
		{
			name:     "successful response message sets as healthy",
			msg:      models.HealthcheckResponseMsg{StatusCode: http.StatusOK},
			wantCmd:  tea.Tick(0, models.CheckHealthCmd(fakeHealthcheckPerformer{})),
			wantView: " ●  Server connection healthy ",
		},
		{
			name:     "successful response message sets as healthy",
			msg:      models.HealthcheckResponseMsg{Err: errors.New("fake error")},
			wantCmd:  tea.Tick(0, models.CheckHealthCmd(fakeHealthcheckPerformer{})),
			wantView: " ◌  Error with server connection ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			healthcheck := models.NewHealthcheck(models.HealthcheckParams{
				CheckFrequency: time.Nanosecond * 1,
				Client:         tt.client,
			})
			updatedHealthcheck, cmd := healthcheck.Update(tt.msg)

			teast.AssertViewsEqual(t, tt.wantView, updatedHealthcheck.View())
			teast.AssertCmdsEqual(t, tt.wantCmd, cmd)
		})
	}
}

func TestCheckHealthCmd(t *testing.T) {
	tests := []struct {
		name    string
		client  fakeHealthcheckPerformer
		wantMsg tea.Msg
	}{
		{
			name:    "successful healthcheck",
			client:  fakeHealthcheckPerformer{statusCode: http.StatusOK},
			wantMsg: models.HealthcheckResponseMsg{StatusCode: http.StatusOK},
		},
		{
			name:    "failed healthcheck",
			client:  fakeHealthcheckPerformer{err: fakeHealthcheckError},
			wantMsg: models.HealthcheckResponseMsg{Err: fakeHealthcheckError},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := models.CheckHealthCmd(tt.client)(time.Now())

			teast.AssertMsgEqual(t, tt.wantMsg, msg)
		})
	}
}
