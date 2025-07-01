package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ResponseView() *tview.TextView {
	view := tview.NewTextView()
	view.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	view.SetTextColor(tcell.ColorWhite.TrueColor())
	view.SetBorder(true)
	view.SetTitle(" > Response ")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetDynamicColors(true)

	return view
}
