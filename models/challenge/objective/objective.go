package objective

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap/zapcore"
	"jobdone.emailaddress.horse/utils/colors"
)

type Objective struct {
	description string
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
		description: params.Description,
		verifier:    params.Verifier,
	}
}

func (o Objective) Update(msg tea.Msg) Objective {
	o.verifier = o.verifier.verify(msg)

	return o
}

var (
	completeCheckbox = lipgloss.NewStyle().Background(colors.Green3).Foreground(colors.Green11).Render("[✓]")
	completeStyle    = lipgloss.NewStyle().Background(colors.Green2).Foreground(colors.Green11)

	incompleteCheckbox = lipgloss.NewStyle().Background(colors.Tomato3).Foreground(colors.Tomato11).Render("[ ]")
	incompleteStyle    = lipgloss.NewStyle().Background(colors.Tomato2).Foreground(colors.Tomato11)
)

func (o Objective) View() string {
	style := incompleteStyle
	checkbox := incompleteCheckbox
	if o.Complete() {
		style = completeStyle
		checkbox = completeCheckbox
	}

	return checkbox + style.Render(" "+o.description)
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
	enc.AddString("Description", o.description)
	enc.AddBool("Complete", o.Complete())
	return nil
}
