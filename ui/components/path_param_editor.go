package ui

import (
	"http_client/ui/template"

	"github.com/rivo/tview"
)

// Crea y devuelve el editor para los parámetros
// Se usa así: user, id, 1. En url es esto user/id/1
func PathParamsEditor() *tview.TextArea {
	return template.TextEditor("1,\nid", " > Path Params ")
}
