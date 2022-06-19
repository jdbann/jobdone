package person

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/world/entity"
)

type Person struct {
	x, y int

	logger *zap.Logger
}

type Params struct {
	X, Y int

	Logger *zap.Logger
}

func New(params Params) entity.Entity {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Person")

	return Person{
		x: params.X,
		y: params.Y,

		logger: logger,
	}
}

func Builder(params Params) entity.Builder {
	return func(logger *zap.Logger) entity.Entity {
		params.Logger = logger
		return New(params)
	}
}

func (m Person) Init() tea.Cmd {
	return nil
}

func (m Person) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case entity.TickMsg:
		m.logger.Debug(
			"Received world tick message",
			zap.Object("tea.Msg", msg),
		)
	}

	return m, nil
}

func (m Person) View() string {
	return "O"
}

func (m Person) Position() (x int, y int) {
	return m.x, m.y
}
