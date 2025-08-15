package ui

import (
	color "http_client/const/color_wrapper"
	"http_client/logic"

	"github.com/rivo/tview"
)

func TabsList() *tview.List {
	list := tview.NewList()
	list.ShowSecondaryText(false)
	list.SetBorder(true)
	list.SetTitle(" > Pestaña ")
	list.SetTitleAlign(tview.AlignLeft)
	logic.ApplySelectedBackgroundIfSupported(list, color.ColorHighlight.TrueColor())
	return list
}
