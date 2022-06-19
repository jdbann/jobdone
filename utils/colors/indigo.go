package colors

import "github.com/charmbracelet/lipgloss"

var (
	Indigo1  = lipgloss.AdaptiveColor{Light: "#fdfdfe", Dark: "#131620"}
	Indigo2  = lipgloss.AdaptiveColor{Light: "#f8faff", Dark: "#15192d"}
	Indigo3  = lipgloss.AdaptiveColor{Light: "#f0f4ff", Dark: "#192140"}
	Indigo4  = lipgloss.AdaptiveColor{Light: "#e6edfe", Dark: "#1c274f"}
	Indigo5  = lipgloss.AdaptiveColor{Light: "#d9e2fc", Dark: "#1f2c5c"}
	Indigo6  = lipgloss.AdaptiveColor{Light: "#c6d4f9", Dark: "#22346e"}
	Indigo7  = lipgloss.AdaptiveColor{Light: "#aec0f5", Dark: "#273e89"}
	Indigo8  = lipgloss.AdaptiveColor{Light: "#8da4ef", Dark: "#2f4eb2"}
	Indigo9  = lipgloss.AdaptiveColor{Light: "#3e63dd", Dark: "#3e63dd"}
	Indigo10 = lipgloss.AdaptiveColor{Light: "#3a5ccc", Dark: "#5373e7"}
	Indigo11 = lipgloss.AdaptiveColor{Light: "#3451b2", Dark: "#849dff"}
	Indigo12 = lipgloss.AdaptiveColor{Light: "#101d46", Dark: "#eef1fd"}
)

var Indigo = Scale{
	Name: "Indigo",
	Steps: []step{
		{"Indigo1", Indigo1},
		{"Indigo2", Indigo2},
		{"Indigo3", Indigo3},
		{"Indigo4", Indigo4},
		{"Indigo5", Indigo5},
		{"Indigo6", Indigo6},
		{"Indigo7", Indigo7},
		{"Indigo8", Indigo8},
		{"Indigo9", Indigo9},
		{"Indigo10", Indigo10},
		{"Indigo11", Indigo11},
		{"Indigo12", Indigo12},
	},
}
