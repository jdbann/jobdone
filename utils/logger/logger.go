package logger

import "go.uber.org/zap"

const DefaultOutputPath = "./jobdone.log"

type Params struct {
	DebugMode bool
}

func New(params Params) (*zap.Logger, error) {
	if !params.DebugMode {
		return zap.NewNop(), nil
	}

	cfg := zap.NewProductionConfig()

	cfg.OutputPaths = []string{DefaultOutputPath}
	cfg.ErrorOutputPaths = []string{DefaultOutputPath}

	cfg.Level.SetLevel(zap.DebugLevel)

	return cfg.Build()
}
