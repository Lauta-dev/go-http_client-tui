package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Keys(app *tview.Application, switchPage *tview.Pages, der *tview.Flex) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'r':
				switchPage.SwitchToPage("response")
				app.SetFocus(switchPage)
			case 'p':
				switchPage.SwitchToPage("pp")
				app.SetFocus(switchPage)
			case 'q':
				switchPage.SwitchToPage("qp")
				app.SetFocus(switchPage)
			case 'h':
				switchPage.SwitchToPage("header")
				app.SetFocus(switchPage)
			case 'b':
				switchPage.SwitchToPage("body")
				app.SetFocus(switchPage)

			case 'd':

			}
		}
		return event
	})
}
