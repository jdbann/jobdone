package teast

import tea "github.com/charmbracelet/bubbletea"

type FakeModel interface {
	tea.Model
	Msgs() []tea.Msg
}

// NewFakeModel builds a simple tea.Model for injecting into tests. Prepare it
// with specific behaviour by passing various FakeModelOptions
func NewFakeModel(opts ...FakeModelOption) *fakeModel {
	f := &fakeModel{}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

type fakeModel struct {
	initCmd   tea.Cmd
	updateCmd tea.Cmd
	view      string

	msgs []tea.Msg
}

type FakeModelOption = func(*fakeModel)

func InitReturns(cmd tea.Cmd) FakeModelOption {
	return func(f *fakeModel) {
		f.initCmd = cmd
	}
}

func UpdateReturns(cmd tea.Cmd) FakeModelOption {
	return func(f *fakeModel) {
		f.updateCmd = cmd
	}
}

func ViewReturns(view string) FakeModelOption {
	return func(f *fakeModel) {
		f.view = view
	}
}

func (f *fakeModel) Init() tea.Cmd {
	return f.initCmd
}

func (f *fakeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	f.msgs = append(f.msgs, msg)
	return f, f.updateCmd
}

func (f *fakeModel) View() string {
	return f.view
}

// Msgs returns the messages which have been passed to the fake model's Update
// method.
func (f *fakeModel) Msgs() []tea.Msg {
	return f.msgs
}
