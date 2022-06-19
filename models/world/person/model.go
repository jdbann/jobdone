package person

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"jobdone.emailaddress.horse/models/world/entity"
	"jobdone.emailaddress.horse/pkg/bub"
	"jobdone.emailaddress.horse/pkg/client"
)

type Person struct {
	x, y     int
	style    lipgloss.Style
	localID  string
	remoteID string
	client   registerPerformer

	logger *zap.Logger
}

type Params struct {
	X, Y   int
	Client registerPerformer

	Logger *zap.Logger
}

func New(params Params) entity.Entity {
	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Person")

	if params.Client == nil {
		params.Client = client.New(client.Params{})
	}

	return Person{
		x:       params.X,
		y:       params.Y,
		style:   personStyle(nextColor()),
		localID: generateLocalID(),
		client:  params.Client,

		logger: logger,
	}
}

func Builder(params Params) entity.Builder {
	return func(logger *zap.Logger) entity.Entity {
		params.Logger = logger
		return New(params)
	}
}

func (m Person) Init() tea.Cmd {
	return RegisterCmd(m.client, m.localID)
}

func (m Person) Update(msg tea.Msg) (entity.Entity, tea.Cmd) {
	switch msg := msg.(type) {
	case entity.TickMsg:
		m.logger.Debug(
			"Received world tick message",
			zap.Object("tea.Msg", msg),
		)
		return randomStep(m, msg.Width, msg.Height), nil

	case RegisterFailedMsg:
		if msg.LocalID != m.localID {
			return m, nil
		}

		m.logger.Debug(
			"Received register failed message",
			zap.Object("tea.Msg", msg),
		)
		return m, tea.Sequentially(bub.Wait(time.Second*2), RegisterCmd(m.client, m.localID))

	case RegisterSucceededMsg:
		if msg.LocalID != m.localID {
			return m, nil
		}

		m.logger.Debug(
			"Received register succeeded message",
			zap.Object("tea.Msg", msg),
		)
		m.remoteID = msg.RemoteID
	}

	return m, nil
}

func (m Person) Render(baseStyle lipgloss.Style) string {
	return m.style.Inherit(baseStyle).Render("O")
}

func (m Person) Position() (x int, y int) {
	return m.x, m.y
}

func generateLocalID() string {
	b := make([]byte, 12)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}
