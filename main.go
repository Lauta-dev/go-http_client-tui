package main

import (
	"flag"
	"fmt"
	"strings"

	colors "http_client/const"
	"http_client/internal/clipboard"
	"http_client/logic"
	"http_client/ui"
	"http_client/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	responseView  *tview.TextView
	responseInfo  *tview.TextView
	history       *tview.List
	app           *tview.Application
	contentToCopy string

	//
	status      string // ex 200, OK
	contentType string // Content-Type: Json
	completeUrl string // URL completa con la que se hace el fetching
)

func updateComponent(
	verb string,
	h map[string]string,
	qp map[string]string,
	p []string,
	body string,
	userUrl string) {
	app.QueueUpdateDraw(func() {
		responseView.Clear()
		responseInfo.Clear()

		loadingFormat := fmt.Sprintf("[%s::b]%s[%s::B]",
			colors.ColorTextSecondary.String(),
			"Cargando...",
			colors.ColorPrimary.String(),
		)

		fmt.Fprintln(responseView, loadingFormat)
		fmt.Fprintln(responseInfo, loadingFormat)
	})

	response, err := Fetching(userUrl, verb, h, qp, p, body)

	if err != nil {
		app.QueueUpdateDraw(func() {
			responseInfo.Clear()
			responseView.Clear()

			fmt.Fprintln(responseView, err.Error())
		})

		return
	}

	app.QueueUpdateDraw(func() {
		responseInfo.Clear()
		responseView.Clear()

		format := fmt.Sprintf("%s, %s \nURL: %s", StatusCodesColors(status), contentType, completeUrl)
		code := strings.Split(status, " ")[0]

		err := logic.SaveItems(completeUrl, code, contentType, response, verb)

		if err != nil {
			fmt.Fprintf(responseInfo, "[red]%s", err.Error())
		} else {
			fmt.Fprintln(responseInfo, format)
		}

		fmt.Fprintln(responseView, response)
	})
}

func TriggerErrorAfterRequest(err error) {
	app.QueueUpdateDraw(func() {
		responseInfo.Clear()
		responseView.Clear()

		fmt.Fprintln(responseView, err.Error())
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
	_, selected := dropdown.GetCurrentOption()
	body := bodyContent.GetText()
	header := utils.ParseHeader(headerPage.GetText())
	queryParams := utils.ParseHeader(queryParamPage.GetText())
	params := utils.ParsePathParams(pathParamPage.GetText())

	url, err := utils.ParseInput(formInput.GetText(), utils.ParseHeader(varr.GetText()))

	if err != nil {
		go TriggerErrorAfterRequest(err)
		return
	}

	go updateComponent(selected, header, queryParams, params, body, url)
}

func main() {

	envFilePath := flag.String("env-file", "", "Cargar archivo .env")
	helpTrigger := flag.Bool("help", false, "Muestra ayuda")
	flag.Parse()

	if *helpTrigger {
		HelpTrigger()
		return
	}

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
	history := ui.History(app)

	mainPage := tview.NewPages()
	workspacePages := tview.NewPages()

	flexx := tview.NewFlex()
	flexx.SetDirection(tview.FlexRow)

	flexx.AddItem(form, 0, 1, true)
	flexx.AddItem(
		workspacePages.
			AddPage("body", bodyEditor, true, true).
			AddPage("header", headerEditor, true, false).
			AddPage("qp", queryParamEditor, true, false).
			AddPage("pp", pathParamEditor, true, false).
			AddPage("var", variableEditor, true, false),
		0, 1, false)

	showHelpPage := false

	responseWindow := tview.NewFlex().SetDirection(tview.FlexRow)
	responseWindow.AddItem(responseView, 0, 8, false)
	responseWindow.AddItem(responseInfo, 0, 1, false)
	windows := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(responseWindow, 0, 2, false)

	Keys(app, workspacePages, windows)

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()

	flex.AddItem(flexx, 0, 1, false)
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
		case tcell.KeyF2:
			if showHelpPage {
				showHelpPage = false
				mainPage.SwitchToPage("main")
			} else {
				showHelpPage = true
				mainPage.SwitchToPage("history")
			}
		}

		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'd':
				app.SetFocus(responseWindow)
				SendInfo(input, dropdown, bodyEditor, headerEditor, queryParamEditor, pathParamEditor, variableEditor)

			case 'n':
				clipboard.Copy(contentToCopy)

			case 'i':
				app.SetFocus(form)
			}
		}
		return event
	})

	mainPage.AddPage("main", flex, true, true)
	mainPage.AddPage("help", helpPage, true, false)
	mainPage.AddPage("history", history, true, false)

	variableEditor.SetText(logic.ReadEnvFile(*envFilePath), false)

	if err := app.SetRoot(mainPage, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
