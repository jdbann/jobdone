package person

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
	"jobdone.emailaddress.horse/utils/colors"
)

func personStyle(scale colors.Scale) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(scale.Step(10))
}

var nextColor = func() func() colors.Scale {
	colorCyle := []colors.Scale{
		colors.Tomato,
		colors.Indigo,
		colors.Green,
	}
	rand.Shuffle(len(colorCyle), func(i, j int) {
		colorCyle[i], colorCyle[j] = colorCyle[j], colorCyle[i]
	})

	cyclePosition := 0

	return func() colors.Scale {
		defer func() {
			cyclePosition = (cyclePosition + 1) % len(colorCyle)
		}()
		return colorCyle[cyclePosition]
	}
}()
