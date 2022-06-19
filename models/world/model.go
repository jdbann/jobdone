package world

import (
	"math/rand"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/challenge"
	"jobdone.emailaddress.horse/utils/colors"
	"jobdone.emailaddress.horse/utils/logger"
)

type World struct {
	height, width       int
	mapHeight, mapWidth int
	mapStr              string

	logger *zap.Logger
}

type Params struct {
	MapHeight, MapWidth int

	Logger *zap.Logger
}

func New(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("World")

	return World{
		logger: logger,
	}
}

func (m World) Init() tea.Cmd {
	return nil
}

func (m World) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)

		m.height, m.width = msg.Height, msg.Width
	case challenge.ChangedMsg:
		m.logger.Debug(
			"Received challenge changed message",
			zap.Object("tea.Msg", msg),
		)

		m.mapHeight, m.mapWidth = msg.Challenge.MapHeight, msg.Challenge.MapWidth
		m.mapStr = fillMap(m)
	}
	return m, nil
}

var (
	offlineMessage = lipgloss.NewStyle().
			Background(colors.Tomato1).
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(colors.Tomato6).
			BorderBackground(colors.Tomato1).
			Foreground(colors.Tomato11).
			Padding(1, 4).
			Render("WORLD MAP OFFLINE")

	offlineStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(colors.Tomato6).
			BorderBackground(colors.Tomato1)

	onlineStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(colors.Indigo6).
			BorderBackground(colors.Indigo1)

	mapStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(colors.Green7).
			BorderBackground(colors.Green3).
			Foreground(colors.Green7).
			Background(colors.Green3)
)

func (m World) View() string {
	if m.mapHeight == 0 || m.mapWidth == 0 {
		return offlineStyle.Render(lipgloss.Place(
			m.width-2,
			m.height-2,
			lipgloss.Center,
			lipgloss.Center,
			offlineMessage,
			lipgloss.WithWhitespaceBackground(colors.Tomato1),
			lipgloss.WithWhitespaceForeground(colors.Tomato4),
			lipgloss.WithWhitespaceChars("\\"),
		))
	}

	return onlineStyle.Render(lipgloss.Place(
		m.width-2,
		m.height-2,
		lipgloss.Center,
		lipgloss.Center,
		mapStyle.Height(m.mapHeight).Width(m.mapWidth).Render(m.mapStr),
		lipgloss.WithWhitespaceBackground(colors.Indigo1),
	))
}

const mapChars = "..,,''``;:                             "

func fillMap(m World) string {
	var b strings.Builder
	for y := 0; y < m.mapHeight; y++ {
		for x := 0; x < m.mapWidth; x++ {
			b.WriteByte(mapChars[rand.Intn(len(mapChars))])
		}
		if y != m.height-1 {
			b.WriteRune('\n')
		}
	}
	return b.String()
}
