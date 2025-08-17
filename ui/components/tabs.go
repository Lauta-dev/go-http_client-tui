package ui

import (
	color "http_client/const/color_wrapper"
	"http_client/logic"

	"github.com/rivo/tview"
)

type TabListBuilder struct {
	PageList    *tview.Pages
	Input       *tview.InputField
	TextArea    *tview.TextView
	List        *tview.List
	DetailsPage *tview.Pages
	Container   *tview.Flex
}

func TabsList() TabListBuilder {
	input := tview.NewInputField()
	textView := tview.NewTextView()

	list := tview.NewList()
	list.ShowSecondaryText(false)

	list.SetSelectedBackgroundColor(color.ColorHighlight.TrueColor())
	list.SetMainTextColor(color.ColorTextPrimary.TrueColor())
	list.SetSelectedTextColor(color.ColorTextPrimary.TrueColor())
	logic.ApplySelectedBackgroundIfSupported(list, color.ColorHighlight.TrueColor())

	pageList := tview.NewPages()
	pageList.AddPage("list", list, true, true)
	pageList.SetBorder(true)
	pageList.SetTitle(" > Pestaña ")
	pageList.SetTitleAlign(tview.AlignLeft)

	detailsPage := tview.NewPages()
	detailsPage.SetBorder(true)
	detailsPage.SetTitle(" > Info | 'E' para editar item ")
	detailsPage.SetTitleAlign(tview.AlignLeft)
	detailsPage.AddPage("info", textView, true, true)
	detailsPage.AddPage("edit-item", input, true, false)

	container := tview.NewFlex()
	container.AddItem(list, 0, 1, true)
	container.AddItem(detailsPage, 0, 1, false)

	return TabListBuilder{
		PageList:    pageList,
		Input:       input,
		List:        list,
		TextArea:    textView,
		DetailsPage: detailsPage,
		Container:   container,
	}
}
