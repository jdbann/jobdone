package cmd

import (
	"flag"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models/app"
	"jobdone.emailaddress.horse/utils/logger"
)

type startConfig struct {
	debugMode bool
}

func Start(args []string) error {
	config, err := parseStartFlags(args)
	if err != nil {
		return err
	}

	logger, err := logger.New(logger.Params{DebugMode: config.debugMode})
	if err != nil {
		return err
	}
	defer logger.Sync()

	logger.Debug("Launching...")

	app := app.New(app.Params{
		Logger: logger,
	})

	p := tea.NewProgram(app, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return err
	}

	logger.Debug("Exiting...")

	return nil
}

func parseStartFlags(args []string) (startConfig, error) {
	config := startConfig{}

	startFlags := flag.NewFlagSet("", flag.ExitOnError)
	startFlags.BoolVar(&config.debugMode, "debug", false, "log debug messages to "+logger.DefaultOutputPath)

	err := startFlags.Parse(args)
	if err != nil {
		return startConfig{}, err
	}

	return config, nil
}
