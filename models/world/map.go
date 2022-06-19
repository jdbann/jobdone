package world

import (
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"jobdone.emailaddress.horse/models/world/entity"
	"jobdone.emailaddress.horse/utils/colors"
)

const mapChars = "..,,''``;:                             "

var (
	grassStyle = lipgloss.NewStyle().
		Foreground(colors.Green7).
		Background(colors.Green3)
)

type ansiMap [][]string

func newMap(width, height int) mapRenderer {
	if width == 0 || height == 0 {
		return nil
	}

	var m = make([][]string, height)

	for y := 0; y < height; y++ {
		m[y] = make([]string, width)
		for x := 0; x < width; x++ {
			m[y][x] = grassStyle.Render(randomMapChar())
		}
	}

	return ansiMap(m)
}

func (m ansiMap) Render(entities entity.Entities) string {
	var b strings.Builder
	for y, mapRow := range m {
		for x, mapTile := range mapRow {
			entity := entities.At(x, y)

			if entity != nil {
				b.WriteString(entity.View())
				continue
			}

			b.WriteString(mapTile)
		}

		if y != len(m)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func randomMapChar() string {
	return string(mapChars[rand.Intn(len(mapChars))])
}
