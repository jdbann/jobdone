package challenge

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/pkg/glam"
	"jobdone.emailaddress.horse/utils/logger"
)

var _ tea.Model = Challenge{}

type renderer = interface {
	Render(string) (string, error)
	SetWordWrap(int) error
}

type Challenge struct {
	challenge Definition

	renderer renderer

	logger *zap.Logger
}

type Params struct {
	Challenge Definition

	Renderer renderer

	Logger *zap.Logger
}

func New(params Params) tea.Model {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	params.Logger = params.Logger.Named("Challenge")

	if params.Renderer == nil {
		var err error
		params.Renderer, err = glam.NewRenderer()
		if err != nil {
			panic(err)
		}
	}

	return Challenge{
		challenge: params.Challenge,

		renderer: params.Renderer,

		logger: params.Logger,
	}
}

func (m Challenge) Init() tea.Cmd {
	return SwitchCmd(Challenge1)
}

func (m Challenge) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)
		if msg.Width <= glam.MaxWidth {
			m.renderer.SetWordWrap(msg.Width)
		}

	case ChangedMsg:
		m.logger.Debug(
			"Received challenge changed message",
			zap.Object("tea.Msg", msg),
		)
		m.challenge = msg.Challenge
		return m, nil
	}

	for i, objective := range m.challenge.Objectives {
		m.challenge.Objectives[i] = objective.Update(msg)
	}

	return m, nil
}

func (m Challenge) View() string {
	if m.challenge.Number == 0 {
		return "No active challenge."
	}

	var content strings.Builder
	err := mdTemplate.ExecuteTemplate(&content, "challenge", m.challenge)
	if err != nil {
		panic(err)
	}

	styledContent, err := m.renderer.Render(content.String())
	if err != nil {
		panic(err)
	}

	for _, objective := range m.challenge.Objectives {
		styledContent = styledContent + "\n" + objective.View()
	}

	return styledContent
}
