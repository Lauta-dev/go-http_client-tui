package main

import "github.com/rivo/tview"

func Buttons(pages *tview.Pages) *tview.Flex {
	optionsContainer := tview.NewFlex().SetDirection(tview.FlexColumn)
	optionsContainer.SetBorder(true)
	optionsContainer.SetTitle(" Buttons ")
	optionsContainer.SetTitleAlign(tview.AlignLeft)
	buttons := tview.NewForm()

	buttons.AddButton("Body", func() {
		pages.SwitchToPage("body")
	})
	buttons.AddButton("Query Params", func() {
		pages.SwitchToPage("qp")
	})
	buttons.AddButton("Path Param", func() {
		pages.SwitchToPage("pp")
	})
	buttons.AddButton("Headers", func() {
		pages.SwitchToPage("header")
	})

	optionsContainer.AddItem(buttons, 0, 1, false)

	return optionsContainer

}
