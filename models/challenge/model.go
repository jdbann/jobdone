package challenge

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

var _ tea.Model = Challenge{}

type Challenge struct {
	number      int
	title       string
	description string

	logger *zap.Logger
}

type Params struct {
	Logger *zap.Logger
}

func New(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	params.Logger = params.Logger.Named("Challenge")

	return Challenge{
		logger: params.Logger,
	}
}

func (m Challenge) Init() tea.Cmd {
	return nil
}

func (m Challenge) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ChangedMsg:
		m.logger.Debug(
			"Received challenge changed message",
			zap.Object("tea.Msg", msg),
		)
		m.number, m.title, m.description = msg.Number, msg.Title, msg.Description
	}

	return m, nil
}

func (m Challenge) View() string {
	if m.number == 0 {
		return "No active challenge."
	}

	return fmt.Sprintf("Challenge #%d: %s\n\n%s", m.number, m.title, m.description)
}
