package main

import (
	"fmt"

	colors "http_client/const"
	"http_client/internal/clipboard"
	"http_client/ui"
	"http_client/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	responseView  *tview.TextView
	responseInfo  *tview.TextView
	app           *tview.Application
	contentToCopy string

	//
	status      string // ex 200, OK
	contentType string // Content-Type: Json
	completeUrl string // URL completa con la que se hace el fetching
)

func updateComponent(userUrl string, verb string, h map[string]string, qp map[string]string, p []string, body string, tmp string) {
	app.QueueUpdateDraw(func() {
		responseView.Clear()
		responseInfo.Clear()

		loadingFormat := fmt.Sprintf("[%s]%s[%s]",
			colors.ColorTextSecondary.String(),
			"Cargando...",
			colors.ColorPrimary.String(),
		)

		fmt.Fprintln(responseView, loadingFormat)
		fmt.Fprintln(responseInfo, loadingFormat)
	})

	data := Fetching(tmp, verb, h, qp, p, body)

	app.QueueUpdateDraw(func() {
		// Set footer info
		responseInfo.Clear()

		format := fmt.Sprintf("%s, %s \nURL: %s", StatusCodesColors(status), contentType, completeUrl)

		fmt.Fprintln(responseInfo, format)

		// Set reponse info
		responseView.Clear()
		response := data
		fmt.Fprintln(responseView, response)
	})

}

func SendInfo(
	formInput *tview.InputField,
	dropdown *tview.DropDown,
	bodyContent *tview.TextArea,
	headerPage *tview.TextArea,
	queryParamPage *tview.TextArea,
	pathParamPage *tview.TextArea,
	varr *tview.TextArea,
) {
	url := formInput.GetText()
	_, selected := dropdown.GetCurrentOption()
	body := bodyContent.GetText()
	header := utils.ParseHeader(headerPage.GetText())
	queryParams := utils.ParseHeader(queryParamPage.GetText())
	params := utils.ParsePathParams(pathParamPage.GetText())

	tmp := utils.ParseInput(url, utils.ParseHeader(varr.GetText()))

	go updateComponent(url, selected, header, queryParams, params, body, tmp)
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
	responseView = ui.ResponseView()
	responseInfo = ui.ResponseInfo()
	variableEditor := ui.VariableEditor()

	mainPage := tview.NewPages()
	workspacePages := tview.NewPages()

	showHelpPage := false

	windows := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(workspacePages, 0, 2, false)

	responseWindow := tview.NewFlex().SetDirection(tview.FlexRow)
	responseWindow.AddItem(responseView, 0, 8, false)
	responseWindow.AddItem(responseInfo, 0, 1, false)

	Keys(app, workspacePages, windows)

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()

	workspacePages.
		AddPage("body", bodyEditor, true, false).
		AddPage("response", responseWindow, true, true).
		AddPage("header", headerEditor, true, false).
		AddPage("qp", queryParamEditor, true, false).
		AddPage("pp", pathParamEditor, true, false).
		AddPage("var", variableEditor, true, false)

	flex.AddItem(form, 0, 1, false)
	flex.AddItem(windows, 0, 1, false)

	fullScreen := false
	windows.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'f':
				fullScreen = !fullScreen
				windows.SetFullScreen(fullScreen)
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
				SendInfo(input, dropdown, bodyEditor, headerEditor, queryParamEditor, pathParamEditor, variableEditor)

			case 'n':
				clipboard.Copy(contentToCopy)
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
