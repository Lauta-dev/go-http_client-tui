package template

import "github.com/rivo/tview"

func TextEditor(placeholder string, title string) *tview.TextArea {
	editor := tview.NewTextArea().SetPlaceholder(placeholder)
	editor.
		SetBorder(true).
		SetTitle(title).
		SetTitleAlign(tview.AlignLeft)
	return editor
}
