package main

import (
	"github.com/rivo/tview"
)

func A() *tview.TextArea {
	ejemplo := "x-api-key: 123,\nset-content: 123,\n#x-test: 123 <- Esto será ignorado"
	bodyHeaders := tview.NewTextArea().SetPlaceholder(ejemplo)
	bodyHeaders.SetBorder(true).SetTitle(" Cuerpo de la petición ").SetTitleAlign(tview.AlignLeft)
	return bodyHeaders
}
