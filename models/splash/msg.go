package splash

import "go.uber.org/zap/zapcore"

type CompleteMsg struct{}

func (msg CompleteMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "splash.CompleteMsg")
	return nil
}

type TickMsg struct{}

func (msg TickMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "splash.TickMsg")
	return nil
}
