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
	input.SetText("http://localhost:4000/items/csv")
	input.SetPlaceholder("http://example.com")
	input.SetPlaceholderTextColor(tcell.ColorGray.TrueColor())

	form.SetFieldBackgroundColor(constants.ColorBackground)
	form.SetFieldTextColor(constants.ColorTextPrimary)
	form.SetBorder(true)
	form.SetTitle(" > Request ('F1' para ayuda) ")
	form.SetTitleAlign(tview.AlignLeft)

	dropdown.SetBackgroundColor(constants.ColorBackground)
	dropdown.SetFieldBackgroundColor(constants.ColorBackground)
	dropdown.SetTitleColor(constants.ColorTextPrimary)
	dropdown.SetFieldTextColor(constants.ColorTextSecondary)

	dropdown.SetOptions(constants.MethodList(), nil)
	form.AddFormItem(dropdown)
	form.AddFormItem(input)

	return form, dropdown, input
}
