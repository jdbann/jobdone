package challenge

import (
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap/zapcore"
	"jobdone.emailaddress.horse/models/challenge/objective"
	"jobdone.emailaddress.horse/models/healthcheck"
)

type Definition struct {
	Number      int
	Title       string
	Description string
	Objectives  objective.Objectives
}

var (
	Challenge1 = Definition{
		Number:      1,
		Title:       "Induction Day",
		Description: "Welcome to your first day! It's a simple start. We just want to get a healthcheck endpoint working so we can make sure you're hooked up to the system correctly.",
		Objectives: []objective.Objective{
			objective.New(objective.Params{
				Description: "Respond successfully to healthcheck request",
				Verifier: objective.NewSimpleVerifier(func(msg tea.Msg) bool {
					hc, ok := msg.(healthcheck.ResponseMsg)
					if !ok || hc.StatusCode != http.StatusOK {
						return false
					}
					return true
				}),
			}),
		},
	}
)

func (d Definition) Complete() bool {
	for _, objective := range d.Objectives {
		if !objective.Complete() {
			return false
		}
	}

	return true
}

func (d Definition) Update(msg tea.Msg) Definition {
	for i, objective := range d.Objectives {
		d.Objectives[i] = objective.Update(msg)
	}

	return d
}

func (d Definition) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("Number", d.Number)
	enc.AddString("Title", d.Title)
	enc.AddString("Description", d.Description)
	enc.AddArray("Objectives", d.Objectives)
	return nil
}
