package models

import (
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jobdone.emailaddress.horse/utils/colors"
	"jobdone.emailaddress.horse/utils/logger"
)

var _ tea.Model = Healthcheck{}

type healthcheckPerformer interface {
	Healthcheck() (*http.Response, error)
}

type HealthcheckClient struct{}

func (c HealthcheckClient) Healthcheck() (*http.Response, error) {
	return http.Get("http://localhost:3000/health")
}

type Healthcheck struct {
	checkFrequency time.Duration
	client         healthcheckPerformer
	healthy        bool
	width          int

	logger *zap.Logger
}

type HealthcheckParams struct {
	CheckFrequency time.Duration
	Client         healthcheckPerformer

	Logger *zap.Logger
}

func NewHealthcheck(params HealthcheckParams) Healthcheck {
	if params.CheckFrequency == 0 {
		params.CheckFrequency = time.Second * 5
	}

	if params.Client == nil {
		params.Client = HealthcheckClient{}
	}

	if params.Logger == nil {
		params.Logger = zap.NewNop()
	}

	logger := params.Logger.Named("Healthcheck")

	return Healthcheck{
		checkFrequency: params.CheckFrequency,
		client:         params.Client,

		logger: logger,
	}
}

func (h Healthcheck) Init() tea.Cmd {
	h.logger.Debug("Initialised")
	return func() tea.Msg { return CheckHealthCmd(h.client)(time.Now()) }
}

func (h Healthcheck) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h.logger.Debug(
			"Received window resize message",
			zap.Object("tea.Msg", logger.WindowSizeMsg(msg)),
		)

		h.width = msg.Width

	case HealthcheckResponseMsg:
		h.logger.Debug(
			"Received healthcheck response message",
			zap.Object("tea.Msg", msg),
		)

		h.healthy = msg.StatusCode == http.StatusOK
		return h, tea.Tick(h.checkFrequency, CheckHealthCmd(h.client))
	}

	return h, nil
}

var (
	msgStyle       = lipgloss.NewStyle().Padding(0, 1)
	indicatorStyle = lipgloss.NewStyle().Padding(0, 1)

	healthyDescription = msgStyle.Copy().Background(colors.Green3).Foreground(colors.Green11).Render("Server connection healthy")
	healthyIndicator   = indicatorStyle.Copy().Background(colors.Green5).Foreground(colors.Green11).Render("●")
	healthyWhitespace  = []lipgloss.WhitespaceOption{lipgloss.WithWhitespaceBackground(colors.Green1)}

	unhealthyDescription = msgStyle.Copy().Background(colors.Tomato3).Foreground(colors.Tomato11).Render("Error with server connection")
	unhealthyIndicator   = indicatorStyle.Copy().Background(colors.Tomato5).Foreground(colors.Tomato11).Render("◌")
	unhealthyWhitespace  = []lipgloss.WhitespaceOption{lipgloss.WithWhitespaceBackground(colors.Tomato1)}
)

func (h Healthcheck) View() string {
	statusMsg := healthyIndicator + healthyDescription
	whitespaceOptions := healthyWhitespace

	if !h.healthy {
		statusMsg = unhealthyIndicator + unhealthyDescription
		whitespaceOptions = unhealthyWhitespace
	}

	return lipgloss.PlaceHorizontal(h.width, lipgloss.Left, statusMsg, whitespaceOptions...)
}

func CheckHealthCmd(client healthcheckPerformer) func(_ time.Time) tea.Msg {
	return func(_ time.Time) tea.Msg {
		res, err := client.Healthcheck()
		if err != nil {
			return HealthcheckResponseMsg{
				Err: err,
			}
		}

		return HealthcheckResponseMsg{
			StatusCode: res.StatusCode,
		}
	}
}

type HealthcheckResponseMsg struct {
	StatusCode int
	Err        error
}

func (m HealthcheckResponseMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", "HealthcheckResponseMsg")
	enc.OpenNamespace("data")
	enc.AddInt("statusCode", m.StatusCode)

	if m.Err != nil {
		enc.AddString("err", m.Err.Error())
	} else {
		enc.AddString("err", "")
	}

	return nil
}
