package teast

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func AssertViewsEqual(t *testing.T, a, b string) {
	t.Helper()
	a, b = StripAnsiCodes(a), StripAnsiCodes(b)
	if a != b {
		t.Errorf("\nExpected: %q\nGot:      %q", a, b)
	}
}

func AssertCmdsEqual(t *testing.T, a, b tea.Cmd) {
	t.Helper()
	if !CmdsEqual(a, b) {
		t.Errorf("\nExpected: %#v\nGot:      %#v", a, b)
	}
}

func AssertMsgEqual(t *testing.T, a, b tea.Msg) {
	t.Helper()
	if a != b {
		t.Errorf("\nExpected: %#v\nGot:      %#v", a, b)
	}
}

func AssertMsgsEqual(t *testing.T, a, b []tea.Msg) {
	t.Helper()
	if len(a) != len(b) {
		t.Errorf("Expected %d msgs, got %d msgs", len(a), len(b))
		return
	}

	for i := range a {
		AssertMsgEqual(t, a[i], b[i])
	}
}
