package ui

import (
	"github.com/rivo/tview"
)

func HeaderEditor() *tview.TextArea {
	return TextEditor("x-api-key: 123,\nset-content: 123,\n#x-test: 123 <- Esto será ignorado", " > Header ")
}
