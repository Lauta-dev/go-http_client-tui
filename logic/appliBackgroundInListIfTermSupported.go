package logic

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Si la terminal soporta truecolor o 256 colores, se aplica el color.
// Caso contrario, no cambia el color (se deja el valor por defecto de tview).
func ApplySelectedBackgroundIfSupported(list *tview.List, desiredColor tcell.Color) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		return
	}

	numColors := screen.Colors()
	if numColors >= 1<<24 || numColors >= 256 {
		list.SetSelectedBackgroundColor(desiredColor)
	} else {
		list.SetSelectedBackgroundColor(tcell.ColorDarkRed.TrueColor())
	}
}
