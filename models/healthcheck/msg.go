package healthcheck

import "go.uber.org/zap/zapcore"

type ResponseMsg struct {
	StatusCode int
	Err        error
}

func (m ResponseMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "HealthcheckResponseMsg")
	enc.OpenNamespace("data")
	enc.AddInt("statusCode", m.StatusCode)

	if m.Err != nil {
		enc.AddString("err", m.Err.Error())
	} else {
		enc.AddString("err", "")
	}

	return nil
}
