package entity

import "go.uber.org/zap/zapcore"

type TickMsg struct{}

func (msg TickMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "world.TickMsg")
	return nil
}
