package challenge

import "go.uber.org/zap/zapcore"

type ChangedMsg struct {
	Challenge Definition
}

func (m ChangedMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "challenge.ChangedMsg")
	enc.OpenNamespace("data")
	enc.AddObject("Challenge", m.Challenge)
	return nil
}
