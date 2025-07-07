package ui

import "github.com/rivo/tview"

func VariableEditor() *tview.TextArea {
	// @url/items/@file

	return TextEditor("url: http://localhost:4000,\nfile: csv", " > Variables ")
}
