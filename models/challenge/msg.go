package challenge

import "go.uber.org/zap/zapcore"

type ChangedMsg struct {
	Number      int
	Title       string
	Description string
}

func (m ChangedMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "challenge.ChangedMsg")
	enc.OpenNamespace("data")
	enc.AddInt("Number", m.Number)
	enc.AddString("Title", m.Title)
	enc.AddString("Description", m.Description)
	return nil
}
