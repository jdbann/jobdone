package glam

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/lipgloss"
)

type Renderer struct {
	style    ansi.StyleConfig
	wordWrap int

	termRenderer *glamour.TermRenderer
}

func NewRenderer() (*Renderer, error) {
	var style ansi.StyleConfig
	if lipgloss.HasDarkBackground() {
		style = dark.buildStyle()
	} else {
		style = light.buildStyle()
	}

	renderer := &Renderer{
		style:    style,
		wordWrap: MaxWidth,
	}

	if err := renderer.updateTermRenderer(); err != nil {
		return nil, err
	}

	return renderer, nil
}

func (r *Renderer) Render(in string) (string, error) {
	return r.termRenderer.Render(in)
}

func (r *Renderer) SetWordWrap(wordWrap int) error {
	r.wordWrap = wordWrap

	return r.updateTermRenderer()
}

func (r *Renderer) updateTermRenderer() error {
	tr, err := glamour.NewTermRenderer(glamour.WithStyles(r.style), glamour.WithWordWrap(r.wordWrap))
	if err != nil {
		return err
	}

	r.termRenderer = tr
	return nil
}
