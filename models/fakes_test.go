package models_test

import tea "github.com/charmbracelet/bubbletea"

// fakeMsg is a fake tea.Msg for testing that messages are passed to nested
// models.
type fakeMsg struct{}

// fakeCmd builds a fake tea.Cmd for testing that commands are returned from
// nested models.
func fakeCmd(msg interface{}) func() tea.Msg {
	return func() tea.Msg { return msg }
}

// fakeModel builds a simple tea.Model for injecting into tests.
func fakeModel(opts ...fakeModelOption) *_fakeModel {
	f := &_fakeModel{}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

type _fakeModel struct {
	initCmd   tea.Cmd
	updateCmd tea.Cmd
	view      string

	msgs []tea.Msg
}

type fakeModelOption = func(*_fakeModel)

func initReturns(cmd tea.Cmd) fakeModelOption {
	return func(f *_fakeModel) {
		f.initCmd = cmd
	}
}

func updateReturns(cmd tea.Cmd) fakeModelOption {
	return func(f *_fakeModel) {
		f.updateCmd = cmd
	}
}

func viewReturns(view string) fakeModelOption {
	return func(f *_fakeModel) {
		f.view = view
	}
}

func (f *_fakeModel) Init() tea.Cmd {
	return f.initCmd
}

func (f *_fakeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	f.msgs = append(f.msgs, msg)
	return f, f.updateCmd
}

func (f *_fakeModel) View() string {
	return f.view
}
