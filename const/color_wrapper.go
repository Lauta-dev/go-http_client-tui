package constants

import (
	"github.com/gdamore/tcell/v2"
)

var (
	// Base
	ColorPrimary   tcell.Color = tcell.GetColor("#89b4fa") // Blue
	ColorSecondary tcell.Color = tcell.GetColor("#cba6f7") // Mauve
	ColorAccent    tcell.Color = tcell.GetColor("#fab387") // Peach

	// Estados
	ColorSuccess tcell.Color = tcell.GetColor("#a6e3a1") // Green
	ColorError   tcell.Color = tcell.GetColor("#f38ba8") // Red
	ColorWarning tcell.Color = tcell.GetColor("#f9e2af") // Yellow
	ColorInfo    tcell.Color = tcell.GetColor("#89dceb") // Sky

	// Texto
	ColorTextPrimary   tcell.Color = tcell.GetColor("#cdd6f4") // Text
	ColorTextSecondary tcell.Color = tcell.GetColor("#bac2de") // Subtext1
	ColorTextMuted     tcell.Color = tcell.GetColor("#a6adc8") // Subtext0

	// Fondo
	ColorBackground tcell.Color = tcell.GetColor("#1e1e2e") // Base
	ColorPanel      tcell.Color = tcell.GetColor("#313244") // Surface0
	ColorHighlight  tcell.Color = tcell.GetColor("#45475a") // Surface1
	ColorDivider    tcell.Color = tcell.GetColor("#585b70") // Surface2

	// Otros
	ColorBorder    tcell.Color = tcell.GetColor("#74c7ec") // Sapphire
	ColorSelection tcell.Color = tcell.GetColor("#b4befe") // Lavender
	ColorScrollbar tcell.Color = tcell.GetColor("#585b70") // Surface2
)
