package ui

import (
	"github.com/rivo/tview"
)

func BodyEditor() *tview.TextArea {
	editor := tview.NewTextArea()
	editor.SetPlaceholder("{'name': 'lautaro'}")
	editor.SetTitle(" > Body ")
	editor.SetBorder(true)
	editor.SetTitleAlign(tview.AlignLeft)

	return editor
}
