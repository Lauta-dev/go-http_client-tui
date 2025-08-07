package ui

import (
	color "http_client/const/color_wrapper"
	http "http_client/const/http_verbs"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Form() (*tview.Form, *tview.DropDown, *tview.InputField) {
	form := tview.NewForm()
	dropdown := tview.NewDropDown().SetLabel("Método")
	input := tview.NewInputField().SetLabel("URL")
	input.SetPlaceholder("http://example.com")
	input.SetPlaceholderTextColor(tcell.ColorGray.TrueColor())

	form.SetFieldBackgroundColor(color.ColorBackground)
	form.SetFieldTextColor(color.ColorTextPrimary)
	form.SetBorder(true)
	form.SetTitle(" > Request ('F1' para ayuda) ")
	form.SetTitleAlign(tview.AlignLeft)

	dropdown.SetBackgroundColor(color.ColorBackground)
	dropdown.SetFieldBackgroundColor(color.ColorBackground)
	dropdown.SetTitleColor(color.ColorTextPrimary)
	dropdown.SetFieldTextColor(color.ColorTextSecondary)

	dropdown.SetOptions(http.MethodList(), nil)
	dropdown.SetCurrentOption(0)
	form.AddFormItem(dropdown)
	form.AddFormItem(input)

	return form, dropdown, input
}
