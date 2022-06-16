package objective

import (
	tea "github.com/charmbracelet/bubbletea"
)

type verifier interface {
	verify(msg tea.Msg) verifier
	complete() bool
}

type nilVerifier struct{}

func (v nilVerifier) verify(_ tea.Msg) verifier {
	return v
}

func (v nilVerifier) complete() bool {
	return false
}

type simpleVerifier struct {
	checkFn  func(tea.Msg) bool
	verified bool
}

// NewSimpleVerifier builds a basic verifier which calls checkFn with any
// tea.Msg received and stores the result.
func NewSimpleVerifier(checkFn func(tea.Msg) bool) verifier {
	return simpleVerifier{checkFn: checkFn}
}

func (v simpleVerifier) verify(msg tea.Msg) verifier {
	v.verified = v.verified || v.checkFn(msg)

	return v
}

func (v simpleVerifier) complete() bool {
	return v.verified
}
