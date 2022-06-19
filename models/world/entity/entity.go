package entity

import (
	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
)

type Entity interface {
	tea.Model
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
