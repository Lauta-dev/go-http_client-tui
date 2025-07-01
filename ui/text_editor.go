package ui

import "github.com/rivo/tview"

func TextEditor(placeholder string, title string) *tview.TextArea {
	//del/1
	pathParam := tview.NewTextArea().SetPlaceholder(placeholder)
	pathParam.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignLeft)
	return pathParam
}
