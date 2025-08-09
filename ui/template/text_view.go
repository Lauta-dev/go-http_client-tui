package template

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func TextView(color tcell.Color, label string) *tview.TextView {
	view := tview.NewTextView()
	view.SetTextColor(color)
	view.SetBorder(true)
	view.SetTitle(" > " + label + " ")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetDynamicColors(true)

	return view
}
