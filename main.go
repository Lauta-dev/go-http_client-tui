package main

import (
	"fmt"

	"http_client/internal/clipboard"
	"http_client/ui"
	"http_client/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	view          *tview.TextView
	footerInfo    *tview.Frame
	app           *tview.Application
	contentToCopy string

	//
	status string // ex 200, OK
)

func updateComponent(userUrl string, verb string, h map[string]string, qp map[string]string, p []string, body string) {
	app.QueueUpdateDraw(func() {
		view.Clear()
		fmt.Fprintln(view, "[yellow]Cargando...[#FFFFFF]")
	})

	data := Fetching(userUrl, verb, h, qp, p, body)

	app.QueueUpdateDraw(func() {
		view.Clear()
		response := StatusCodesColors(status) + "\n\n" + data
		fmt.Fprintln(view, response)
	})

}

func SendInfo(
	formInput *tview.InputField,
	dropdown *tview.DropDown,
	bodyContent *tview.TextArea,
	headerPage *tview.TextArea,
	queryParamPage *tview.TextArea,
	pathParamPage *tview.TextArea,

) {
	url := formInput.GetText()
	_, selected := dropdown.GetCurrentOption()
	body := bodyContent.GetText()
	header := utils.ParseHeader(headerPage.GetText())
	queryParams := utils.ParseHeader(queryParamPage.GetText())
	params := utils.ParsePathParams(pathParamPage.GetText())

	go updateComponent(url, selected, header, queryParams, params, body)
}

func main() {
	app = tview.NewApplication()
	flex := tview.NewFlex()

	bodyEditor := ui.BodyEditor()
	headerEditor := ui.HeaderEditor()
	queryParamEditor := ui.QueryParamsEditor()
	pathParamEditor := ui.PathParamsEditor()
	form, dropdown, input := ui.Form()
	helpPage := ui.Help()

	mainPage := tview.NewPages()
	workspacePages := tview.NewPages()
	view = ui.ResponseView()

	showHelpPage := false

	der := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(workspacePages, 0, 2, false)

	Keys(app, workspacePages, der)

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()

	workspacePages.
		AddPage("body", bodyEditor, true, false).
		AddPage("response", view, true, true).
		AddPage("header", headerEditor, true, false).
		AddPage("qp", queryParamEditor, true, false).
		AddPage("pp", pathParamEditor, true, false)

	// Parte izq
	flex.AddItem(
		tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(form, 0, 5, false),

		0,
		1,
		false)

	// Parte der
	flex.AddItem(der, 0, 1, false)

	fullScreen := false
	der.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'f':
				fullScreen = !fullScreen
				der.SetFullScreen(fullScreen)
				return nil
			}
		}
		return event
	})

	mainPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyF1:
			if showHelpPage {
				showHelpPage = false
				mainPage.SwitchToPage("main")
			} else {
				showHelpPage = true
				mainPage.SwitchToPage("help")
			}
		}

		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'd':
				workspacePages.SwitchToPage("response")
				SendInfo(input, dropdown, bodyEditor, headerEditor, queryParamEditor, pathParamEditor)

			case 'c':
				clipboard.Copy(contentToCopy)
			case 'q':
				app.Stop()
			}
		}
		return event
	})

	mainPage.AddPage("main", flex, true, true)
	mainPage.AddPage("help", helpPage, true, false)

	if err := app.SetRoot(mainPage, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
