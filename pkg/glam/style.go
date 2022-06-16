package glam

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/lipgloss"
	"jobdone.emailaddress.horse/utils/colors"
)

const MaxWidth = 80

type theme struct {
	baseStyle   ansi.StyleConfig
	selectColor func(lipgloss.AdaptiveColor) *string
}

var (
	dark = theme{
		baseStyle:   glamour.DarkStyleConfig,
		selectColor: func(c lipgloss.AdaptiveColor) *string { return &c.Dark },
	}
	light = theme{
		baseStyle:   glamour.LightStyleConfig,
		selectColor: func(c lipgloss.AdaptiveColor) *string { return &c.Light },
	}
)

func (t theme) buildStyle() ansi.StyleConfig {
	style := t.baseStyle

	style.Document.StylePrimitive.Color = t.selectColor(colors.Indigo11)
	style.Document.StylePrimitive.BackgroundColor = t.selectColor(colors.Indigo1)
	style.Document.StylePrimitive.BlockPrefix = ""
	style.Document.StylePrimitive.BlockSuffix = ""
	style.Document.Margin = uintPtr(0)

	style.Heading.StylePrimitive.Color = t.selectColor(colors.Indigo12)
	style.H1.StylePrimitive.BackgroundColor = t.selectColor(colors.Indigo9)

	style.HorizontalRule.Color = t.selectColor(colors.Indigo6)

	return style
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }
