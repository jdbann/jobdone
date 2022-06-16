package objective

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap/zapcore"
)

type Objective struct {
	Description string
	verifier    verifier
}

type Objectives []Objective

type Params struct {
	Description string
	Verifier    verifier
}

func New(params Params) Objective {
	if params.Verifier == nil {
		params.Verifier = nilVerifier{}
	}

	return Objective{
		Description: params.Description,
		verifier:    params.Verifier,
	}
}

func (o Objective) Update(msg tea.Msg) Objective {
	o.verifier = o.verifier.verify(msg)

	return o
}

func (o Objective) Complete() bool {
	return o.verifier.complete()
}

func (o Objectives) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, objective := range o {
		enc.AppendObject(objective)
	}
	return nil
}

func (o Objective) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("Description", o.Description)
	enc.AddBool("Complete", o.Complete())
	return nil
}
