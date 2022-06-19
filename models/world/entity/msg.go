package entity

import "go.uber.org/zap/zapcore"

type TickMsg struct {
	Height, Width int
}

func (msg TickMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "world.TickMsg")
	enc.OpenNamespace("data")
	enc.AddInt("Height", msg.Height)
	enc.AddInt("Width", msg.Width)
	return nil
}
