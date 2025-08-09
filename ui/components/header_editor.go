package ui

import (
	"http_client/ui/template"

	"github.com/rivo/tview"
)

// Crea y devuelve el editor para añadir headers
// Ejemplo de uso: x-api-key: 123, # Estos es un comentario
func HeaderEditor() *tview.TextArea {
	return template.TextEditor("x-api-key: 123,\nset-content: 123,\n#x-test: 123 <- Esto será ignorado", " > Header ")
}
