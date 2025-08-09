package ui

import (
	"http_client/ui/template"

	"github.com/rivo/tview"
)

// Crea y devuelve el editor de Variables
// Este se usa en el input de la url con @url/items
func VariableEditor() *tview.TextArea {
	return template.TextEditor("url: http://localhost:4000,\nfile: csv", " > Variables ")
}
