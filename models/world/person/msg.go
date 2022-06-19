package person

import "go.uber.org/zap/zapcore"

type RegisterFailedMsg struct {
	LocalID string
	Err     error
}

func (msg RegisterFailedMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "person.RegisterFailedMsg")
	enc.OpenNamespace("data")
	enc.AddString("LocalID", msg.LocalID)
	enc.AddString("Err", msg.Err.Error())
	return nil
}

type RegisterSucceededMsg struct {
	LocalID  string
	RemoteID string
}

func (msg RegisterSucceededMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "person.RegisterSucceededMsg")
	enc.OpenNamespace("data")
	enc.AddString("LocalID", msg.LocalID)
	enc.AddString("RemoteID", msg.RemoteID)
	return nil
}
