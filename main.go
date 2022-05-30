package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"jobdone.emailaddress.horse/models"
)

func main() {
	p := tea.NewProgram(models.NewApp(models.AppParams{}), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("OH NO! There has been an error: %v", err)
		os.Exit(1)
	}
}
