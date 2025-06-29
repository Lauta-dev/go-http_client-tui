package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Help() *tview.Frame {
	dialog := "Cambiar entre botones\n\nBody Alt-b\nQuery Param Alt-q\nPath Param Alt-p\nHeaders Alt-h\nResponse Alt-r\nSend Request Alt-d\nSet full screen Alt-f"
	text := tview.NewTextView().SetDynamicColors(true)
	text.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	text.SetTextColor(tcell.ColorWhite.TrueColor())

	fmt.Fprintln(text, dialog)

	frame := tview.NewFrame(text).SetBorders(1, 1, 1, 1, 1, 1).AddText("> Para volver presione (F1)", false, tview.AlignLeft, tcell.ColorWhite.TrueColor())

	frame.SetBackgroundColor(tcell.ColorBlack.TrueColor())

	return frame

}
