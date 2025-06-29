package main

import "github.com/rivo/tview"

func QueryParams() *tview.TextArea {
	//Algo?limit=1
	ejemplo := "limit: 1,\nx: algo, #algo: 123 <- Se ignora"
	bodyQueryParams := tview.NewTextArea().SetPlaceholder(ejemplo)
	bodyQueryParams.SetBorder(true).SetTitle(" Query Params ").SetTitleAlign(tview.AlignLeft)
	return bodyQueryParams
}
