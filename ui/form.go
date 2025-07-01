package ui

import (
	constants "http_client/const"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Form() (*tview.Form, *tview.DropDown, *tview.InputField) {
	form := tview.NewForm()
	dropdown := tview.NewDropDown().SetLabel("Método")
	input := tview.NewInputField().SetLabel("URL")

	form.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	form.SetFieldBackgroundColor(tcell.ColorDarkSlateGray.TrueColor())
	form.SetTitleColor(tcell.ColorWhite.TrueColor())
	form.SetFieldTextColor(tcell.ColorWhite.TrueColor())
	form.SetBorder(true)
	form.SetTitle(" > Request ")
	form.SetTitleAlign(tview.AlignLeft)
	form.SetButtonBackgroundColor(tcell.ColorBlack.TrueColor())

	dropdown.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	dropdown.SetFieldBackgroundColor(tcell.ColorBlack.TrueColor())
	dropdown.SetTitleColor(tcell.ColorWhite.TrueColor())
	dropdown.SetFieldTextColor(tcell.ColorWhite.TrueColor())
	dropdown.SetFieldStyle(tcell.StyleDefault.Blink(true))

	input.SetPlaceholder("http://example.com")
	input.SetPlaceholderTextColor(tcell.ColorGray.TrueColor())

	dropdown.SetOptions(constants.MethodList(), nil)
	form.AddFormItem(dropdown)
	form.AddFormItem(input)

	return form, dropdown, input
}
