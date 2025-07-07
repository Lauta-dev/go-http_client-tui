package ui

import (
	"github.com/rivo/tview"
	constants "http_client/const"
)

func ResponseView() *tview.TextView {
	view := tview.NewTextView()
	view.SetTextColor(constants.ColorTextPrimary)
	view.SetBorder(true)
	view.SetTitle(" > Response ")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetDynamicColors(true)

	return view
}
