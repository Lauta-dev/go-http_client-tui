package ui

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func helpFile() string {
	file := "../help.txt"
	contenido, err := os.ReadFile(file)
	if err != nil {
		return err.Error()
	}
	return string(contenido)

}

func Help() *tview.Frame {
	dialog := helpFile()

	text := tview.NewTextView().SetDynamicColors(true)
	text.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	text.SetTextColor(tcell.ColorWhite.TrueColor())
	text.SetBorder(true).SetTitle(" > Help ").SetTitleAlign(tview.AlignLeft)
	text.SetBorderPadding(1, 1, 1, 1)

	fmt.Fprintln(text, dialog)

	frame := tview.NewFrame(text).SetBorders(1, 1, 1, 1, 1, 1).AddText("> Para volver presione (F1)", false, tview.AlignLeft, tcell.ColorWhite.TrueColor())

	frame.SetBackgroundColor(tcell.ColorBlack.TrueColor())

	return frame

}
