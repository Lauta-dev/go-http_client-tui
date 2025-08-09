package ui

import (
	"http_client/ui/template"

	"github.com/rivo/tview"
)

// Crea y devuelve el editor de Body que se envia
// Ejemplo: { "id": 1 }
func BodyEditor() *tview.TextArea {
	return template.TextEditor("{ 'id': 1 }", " > Body ")
}
