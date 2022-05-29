package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"jobdone.emailaddress.horse/utils/colors"
)

const title = `         _/    _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/
        _/  _/    _/  _/    _/      _/    _/  _/    _/  _/_/    _/  _/
       _/  _/    _/  _/_/_/        _/    _/  _/    _/  _/  _/  _/  _/_/_/
_/    _/  _/    _/  _/    _/      _/    _/  _/    _/  _/    _/_/  _/
 _/_/      _/_/    _/_/_/        _/_/_/      _/_/    _/      _/  _/_/_/_/`

const subtitle = "You write the app, we'll do the hard work."

var _ tea.Model = Splash{}

type Splash struct {
	height, width int
}

func NewSplash() Splash {
	return Splash{}
}

func (s Splash) Init() tea.Cmd {
	return nil
}

func (s Splash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.height, s.width = msg.Height, msg.Width
	}

	return s, nil
}

func (s Splash) View() string {
	styledTitle := lipgloss.NewStyle().
		Foreground(colors.Indigo12).
		Background(colors.Indigo2).
		Padding(2, 0, 2, 0).
		Render(title)

	titleBar := lipgloss.PlaceHorizontal(
		s.width,
		lipgloss.Center,
		styledTitle,
		lipgloss.WithWhitespaceBackground(colors.Indigo2),
	)

	styledSubtitle := lipgloss.NewStyle().
		Foreground(colors.Indigo11).
		Background(colors.Indigo2).
		Padding(0, 0, 2, 0).
		Render(subtitle)

	subtitleBar := lipgloss.PlaceHorizontal(
		s.width,
		lipgloss.Center,
		styledSubtitle,
		lipgloss.WithWhitespaceBackground(colors.Indigo2),
	)

	return lipgloss.PlaceVertical(
		s.height,
		lipgloss.Center,
		titleBar+"\n"+subtitleBar,
		lipgloss.WithWhitespaceBackground(colors.Indigo1),
	)
}
