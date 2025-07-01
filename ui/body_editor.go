package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BodyEditor() *tview.TextArea {
	editor := tview.NewTextArea()
	editor.
		SetPlaceholder("{'name': 'lautaro'}").
		SetTitle(" > Body ").
		SetBorder(true).
		SetBackgroundColor(tcell.ColorBlack.TrueColor()).
		SetTitleAlign(tview.AlignLeft)

	return editor
}
