package person

import tea "github.com/charmbracelet/bubbletea"

type registerPerformer interface {
	RegisterPerson(localID string) (remoteID string, err error)
}

func RegisterCmd(client registerPerformer, localID string) tea.Cmd {
	return func() tea.Msg {
		remoteID, err := client.RegisterPerson(localID)
		if err != nil {
			return RegisterFailedMsg{
				LocalID: localID,
				Err:     err,
			}
		}

		return RegisterSucceededMsg{
			LocalID:  localID,
			RemoteID: remoteID,
		}
	}
}
