package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap/zapcore"
)

// keyMsg is a wrapper for tea.KeyMsg to satisfy the zapcore.ObjectMarshaler
// interface so it can be logged with `zap.Object("tea.Msg", keyMsg(msg))`.
type keyMsg tea.KeyMsg

func (m keyMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "tea.KeyMsg")
	enc.OpenNamespace("data")
	enc.AddString("keyType", m.Type.String())
	enc.AddString("runes", string(m.Runes))
	enc.AddBool("alt", m.Alt)
	return nil
}

// windowSizeMsg is a wrapper for tea.WindowSizeMsg to satisfy the
// zapcore.ObjectMarshaler interface so it can be logged with
// `zap.Object("tea.Msg", windowSizeMsg(msg))`.
type windowSizeMsg tea.WindowSizeMsg

func (m windowSizeMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "tea.WindowSizeMsg")
	enc.OpenNamespace("data")
	enc.AddInt("width", m.Width)
	enc.AddInt("height", m.Height)
	return nil
}
