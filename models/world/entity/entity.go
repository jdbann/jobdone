package entity

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
)

type Entity interface {
	Update(tea.Msg) (Entity, tea.Cmd)
	Render(baseStyle lipgloss.Style) string
	Position() (x, y int)
}

type Builder func(logger *zap.Logger) Entity

type Entities []Entity

func (es Entities) At(x, y int) Entity {
	for _, e := range es {
		eX, eY := e.Position()

		if x == eX && y == eY {
			return e
		}
	}

	return nil
}
