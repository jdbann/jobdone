package teast

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// NewFakeModel builds a simple tea.Model for injecting into tests. Prepare it
// with specific behaviour by passing various FakeModelOptions
func NewFakeModel(t *testing.T, opts ...FakeModelOption) tea.Model {
	f := &fakeModel{}

	for _, opt := range opts {
		opt(f)
	}

	t.Cleanup(func() {
		AssertMsgsEqual(t, f.expectedMsgs, f.msgs)
	})

	return f
}

type fakeModel struct {
	initCmd      tea.Cmd
	updateCmd    tea.Cmd
	view         string
	expectedMsgs []tea.Msg

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

func ExpectMessages(msgs ...tea.Msg) FakeModelOption {
	return func(f *fakeModel) {
		f.expectedMsgs = msgs
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
