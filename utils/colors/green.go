package colors

import "github.com/charmbracelet/lipgloss"

var (
	Green1  = lipgloss.AdaptiveColor{Light: "#fbfefc", Dark: "#0d1912"}
	Green2  = lipgloss.AdaptiveColor{Light: "#f2fcf5", Dark: "#0c1f17"}
	Green3  = lipgloss.AdaptiveColor{Light: "#e9f9ee", Dark: "#0f291e"}
	Green4  = lipgloss.AdaptiveColor{Light: "#ddf3e4", Dark: "#113123"}
	Green5  = lipgloss.AdaptiveColor{Light: "#ccebd7", Dark: "#133929"}
	Green6  = lipgloss.AdaptiveColor{Light: "#b4dfc4", Dark: "#164430"}
	Green7  = lipgloss.AdaptiveColor{Light: "#92ceac", Dark: "#1b543a"}
	Green8  = lipgloss.AdaptiveColor{Light: "#5bb98c", Dark: "#236e4a"}
	Green9  = lipgloss.AdaptiveColor{Light: "#30a46c", Dark: "#30a46c"}
	Green10 = lipgloss.AdaptiveColor{Light: "#299764", Dark: "#3cb179"}
	Green11 = lipgloss.AdaptiveColor{Light: "#18794e", Dark: "#4cc38a"}
	Green12 = lipgloss.AdaptiveColor{Light: "#153226", Dark: "#e5fbeb"}
)

var Green = Scale{
	Name: "Green",
	Steps: []step{
		{"Green1", Green1},
		{"Green2", Green2},
		{"Green3", Green3},
		{"Green4", Green4},
		{"Green5", Green5},
		{"Green6", Green6},
		{"Green7", Green7},
		{"Green8", Green8},
		{"Green9", Green9},
		{"Green10", Green10},
		{"Green11", Green11},
		{"Green12", Green12},
	},
}
