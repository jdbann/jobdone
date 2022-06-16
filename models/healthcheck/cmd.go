package healthcheck

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func CheckCmd(client healthcheckPerformer) func(_ time.Time) tea.Msg {
	return func(_ time.Time) tea.Msg {
		res, err := client.Healthcheck()
		if err != nil {
			return ResponseMsg{
				Err: err,
			}
		}

		return ResponseMsg{
			StatusCode: res.StatusCode,
		}
	}
}
