package challenge

import "go.uber.org/zap/zapcore"

type Definition struct {
	Number      int
	Title       string
	Description string
}

var (
	Challenge1 = Definition{
		Number:      1,
		Title:       "Induction Day",
		Description: "Welcome to your first day! It's a simple start. We just want to get a healthcheck endpoint working so we can make sure you're hooked up to the system correctly.",
	}
)

func (d Definition) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("Number", d.Number)
	enc.AddString("Title", d.Title)
	enc.AddString("Description", d.Description)
	return nil
}
