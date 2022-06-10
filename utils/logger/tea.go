package logger

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap/zapcore"
)

// keyMsg is a wrapper for tea.KeyMsg to satisfy the zapcore.ObjectMarshaler
// interface so it can be logged with `zap.Object("tea.Msg", keyMsg(msg))`.
type KeyMsg tea.KeyMsg

func (m KeyMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "tea.KeyMsg")
	enc.OpenNamespace("data")
	enc.AddString("keyType", m.Type.String())
	enc.AddString("runes", string(m.Runes))
	enc.AddBool("alt", m.Alt)
	return nil
}

// WindowSizeMsg is a wrapper for tea.WindowSizeMsg to satisfy the
// zapcore.ObjectMarshaler interface so it can be logged with
// `zap.Object("tea.Msg", WindowSizeMsg(msg))`.
type WindowSizeMsg tea.WindowSizeMsg

func (m WindowSizeMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "tea.WindowSizeMsg")
	enc.OpenNamespace("data")
	enc.AddInt("width", m.Width)
	enc.AddInt("height", m.Height)
	return nil
}
