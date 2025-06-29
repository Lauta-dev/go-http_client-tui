package main

import "github.com/rivo/tview"

func PathParams() *tview.TextArea {
	//del/1
	ejemplo := "del, 1 , #algo: 123 <- Se ignora"
	pathParam := tview.NewTextArea().SetPlaceholder(ejemplo)
	pathParam.SetBorder(true).SetTitle(" Path Param ").SetTitleAlign(tview.AlignLeft)
	return pathParam
}
