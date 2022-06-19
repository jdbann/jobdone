package person

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/world/entity"
)

type Person struct {
	x, y  int
	style lipgloss.Style

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
		x:     params.X,
		y:     params.Y,
		style: personStyle(nextColor()),

		logger: logger,
	}
}

func Builder(params Params) entity.Builder {
	return func(logger *zap.Logger) entity.Entity {
		params.Logger = logger
		return New(params)
	}
}

func (m Person) Update(msg tea.Msg) (entity.Entity, tea.Cmd) {
	switch msg := msg.(type) {
	case entity.TickMsg:
		m.logger.Debug(
			"Received world tick message",
			zap.Object("tea.Msg", msg),
		)

		m.x = constrain(0, msg.Width-1, m.x+randStep())
		m.y = constrain(0, msg.Height-1, m.y+randStep())
	}

	return m, nil
}

func (m Person) Render(baseStyle lipgloss.Style) string {
	return m.style.Inherit(baseStyle).Render("O")
}

func (m Person) Position() (x int, y int) {
	return m.x, m.y
}

var steps = []int{-1, 0, 0, 0, 1}

func randStep() int {
	return steps[rand.Intn(5)]
}

func constrain(min, max, val int) int {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}
