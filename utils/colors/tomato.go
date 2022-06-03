package colors

import "github.com/charmbracelet/lipgloss"

var (
	Tomato1  = lipgloss.AdaptiveColor{Light: "#fffcfc", Dark: "#1d1412"}
	Tomato2  = lipgloss.AdaptiveColor{Light: "#fff8f7", Dark: "#2a1410"}
	Tomato3  = lipgloss.AdaptiveColor{Light: "#fff0ee", Dark: "#3b1813"}
	Tomato4  = lipgloss.AdaptiveColor{Light: "#ffe6e2", Dark: "#481a14"}
	Tomato5  = lipgloss.AdaptiveColor{Light: "#fdd8d3", Dark: "#541c15"}
	Tomato6  = lipgloss.AdaptiveColor{Light: "#fac7be", Dark: "#652016"}
	Tomato7  = lipgloss.AdaptiveColor{Light: "#f3b0a2", Dark: "#7f2315"}
	Tomato8  = lipgloss.AdaptiveColor{Light: "#ea9280", Dark: "#a42a12"}
	Tomato9  = lipgloss.AdaptiveColor{Light: "#e54d2e", Dark: "#e54d2e"}
	Tomato10 = lipgloss.AdaptiveColor{Light: "#db4324", Dark: "#ec5e41"}
	Tomato11 = lipgloss.AdaptiveColor{Light: "#ca3214", Dark: "#f16a50"}
	Tomato12 = lipgloss.AdaptiveColor{Light: "#341711", Dark: "#feefec"}
)

var Tomato = scale{
	Name: "Tomato",
	Steps: []step{
		{"Tomato1", Tomato1},
		{"Tomato2", Tomato2},
		{"Tomato3", Tomato3},
		{"Tomato4", Tomato4},
		{"Tomato5", Tomato5},
		{"Tomato6", Tomato6},
		{"Tomato7", Tomato7},
		{"Tomato8", Tomato8},
		{"Tomato9", Tomato9},
		{"Tomato10", Tomato10},
		{"Tomato11", Tomato11},
		{"Tomato12", Tomato12},
	},
}
