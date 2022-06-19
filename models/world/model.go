package world

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/challenge"
	"jobdone.emailaddress.horse/models/world/entity"
	"jobdone.emailaddress.horse/utils/colors"
	"jobdone.emailaddress.horse/utils/logger"
)

type mapRenderer interface {
	Render(entity.Entities) string
}

type World struct {
	height, width       int
	mapHeight, mapWidth int
	mapRenderer         mapRenderer
	entities            entity.Entities

	logger *zap.Logger
}

type Params struct {
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
	var cmds []tea.Cmd

	for _, entity := range m.entities {
		cmds = append(cmds, entity.Init())
	}

	return tea.Batch(cmds...)
}

func (m World) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)
		m.height, m.width = msg.Height, msg.Width
		return m, nil

	case challenge.ChangedMsg:
		m.logger.Debug(
			"Received challenge changed message",
			zap.Object("tea.Msg", msg),
		)
		m.mapHeight, m.mapWidth = msg.Challenge.MapHeight, msg.Challenge.MapWidth
		m.mapRenderer = newMap(msg.Challenge.MapWidth, msg.Challenge.MapHeight)
		m.entities = make(entity.Entities, len(msg.Challenge.Entities))
		for i, builder := range msg.Challenge.Entities {
			m.entities[i] = builder(m.logger)
			cmds = append(cmds, m.entities[i].Init())
		}
		cmds = append(cmds, entity.TickCmd(m.mapWidth, m.mapHeight))
		return m, tea.Batch(cmds...)

	case entity.TickMsg:
		m.logger.Debug(
			"Received world tick message",
			zap.Object("tea.Msg", msg),
		)
		cmds = append(cmds, entity.TickCmd(m.mapWidth, m.mapHeight))
	}

	var cmd tea.Cmd
	for i, entity := range m.entities {
		m.entities[i], cmd = entity.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
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
			BorderBackground(colors.Green3)
)

func (m World) View() string {
	if m.mapRenderer == nil {
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
		mapStyle.Render(m.mapRenderer.Render(m.entities)),
		lipgloss.WithWhitespaceBackground(colors.Indigo1),
	))
}
