package ui

import (
	"github.com/rivo/tview"
	color "http_client/const/color_wrapper"
)

func ResponseView() *tview.TextView {
	view := tview.NewTextView()
	view.SetTextColor(color.ColorTextPrimary)
	view.SetBorder(true)
	view.SetTitle(" > Response ")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetDynamicColors(true)

	return view
}
