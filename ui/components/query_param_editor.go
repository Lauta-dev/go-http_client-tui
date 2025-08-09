package ui

import (
	"http_client/ui/template"

	"github.com/rivo/tview"
)

// Crea y devuelve el editor para los Query Params
// Esto se usa así: "limit: 1"
func QueryParamsEditor() *tview.TextArea {
	//Algo?limit=1
	return template.TextEditor("limit: 1,\nx: algo, #algo: 123 <- Se ignora", " > Query Params ")
}
