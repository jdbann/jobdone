package models_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func assertViewsEqual(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("\nExpected: %q\nGot:      %q", a, b)
	}
}

func assertCmdsEqual(t *testing.T, a, b tea.Cmd) {
	if !cmdsEqual(a, b) {
		t.Errorf("\nExpected: %#v\nGot:      %#v", a, b)
	}
}

func assertMsgEqual(t *testing.T, a, b tea.Msg) {
	if a != b {
		t.Errorf("\nExpected: %#v\nGot:      %#v", a, b)
	}
}

func assertMsgsEqual(t *testing.T, a, b []tea.Msg) {
	if len(a) != len(b) {
		t.Errorf("Expected %d msgs, got %d msgs", len(a), len(b))
		return
	}

	for i := range a {
		assertMsgEqual(t, a[i], b[i])
	}
}
