package ui

import "github.com/rivo/tview"

func QueryParamsEditor() *tview.TextArea {
	//Algo?limit=1
	return TextEditor("limit: 1,\nx: algo, #algo: 123 <- Se ignora", " > Query Params ")
}
