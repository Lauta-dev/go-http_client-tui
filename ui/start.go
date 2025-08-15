package ui

import (
	"fmt"
	"http_client/cmd"
	"http_client/internal/clipboard"
	"http_client/logic"
	component "http_client/ui/components"
	"http_client/ui/events"
	"http_client/ui/layout"
	"http_client/ui/shotcust"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

type Tabs struct {
	ID             string
	Input          string
	DropDownOption string
	Headers        string
	QueryParam     string
	PathParam      string
	Variables      string
	Body           string
	ResponseBody   string
	ResponseInfo   string
}

var (
	responseView  *tview.TextView
	responseInfo  *tview.TextView
	app           *tview.Application
	contentToCopy string
)

func StartApp() {

	cli := cmd.Launch()

	firstId := uuid.New().String()
	var currentId string = firstId
	var currentListId int

	showHelpPage := false
	showTabPage := false
	fullScreen := false

	app = tview.NewApplication()
	mainPage := tview.NewPages() // Contiene toda la UI

	main := layout.MainLayout()
	responseView = main.RightPanel.ResponseView
	responseInfo = main.RightPanel.ResponseInfo

	mainLayout := main.Main

	tabList := component.TabsList()

	// Primer elemento que hace referencia al primer estado de la app
	tabList.AddItem("Not Found", firstId, 0, nil)

	tabsMap := map[string]Tabs{}

	mainPage.AddPage("main", mainLayout, true, true)
	mainPage.AddPage("help", component.Help(), true, false)
	mainPage.AddPage("tab", tabList, true, false)
	if cli.ActHistory {
		mainPage.AddPage("history", component.History(app), true, false)

	}

	requestSender := &events.RequestSender{
		App:           app,
		ResponseView:  responseView,
		ResponseInfo:  responseInfo,
		ContentToCopy: &contentToCopy,
	}

	shortcusts := &shotcust.Shortcuts{
		App:                    app,
		MainPage:               mainPage,
		SwitchPage:             main.EditorPanel.Editor,
		ResponseFlex:           main.RightPanel.Container,
		ShowRequestHistoryPage: showHelpPage,
		ShowHelpPage:           showHelpPage,
		RightPanel:             main.RightPanel.Container,
		ChangeToFullScreen:     fullScreen,
		ShowTabPage:            showTabPage,

		CopyFn: func() { clipboard.Copy(contentToCopy) },
		ResponseFn: func() {
			requestSender.SendRequest(
				main.LeftPanel.Input,
				main.LeftPanel.DropDown,
				main.EditorPanel.Body,
				main.EditorPanel.Header,
				main.EditorPanel.QueryParam,
				main.EditorPanel.PathParam,
				main.EditorPanel.Variable, cli.ActHistory)
		},
		FocusForm: func() { app.SetFocus(main.Form.Container) },
		SaveStateFn: func() {
			// Este se usa para editar la pestaña actual
			inputText := main.LeftPanel.Input.GetText()
			dropDown := main.LeftPanel.DropDown.GetTitle()

			headers := main.EditorPanel.Header.GetText()
			queryParams := main.EditorPanel.QueryParam.GetText()
			pathParams := main.EditorPanel.PathParam.GetText()
			variables := main.EditorPanel.Variable.GetText()
			body := main.EditorPanel.Body.GetText()

			responseBody := main.RightPanel.ResponseView.GetText(true)
			responseInfo := main.RightPanel.ResponseInfo.GetText(true)

			tabsMap[currentId] = Tabs{
				ID:             currentId,
				Input:          inputText,
				DropDownOption: dropDown,
				Headers:        headers,
				QueryParam:     queryParams,
				PathParam:      pathParams,
				Variables:      variables,
				Body:           body,
				ResponseBody:   responseBody,
				ResponseInfo:   responseInfo,
			}

			statusCode := strings.SplitN(responseInfo, ",", 2)
			mainText := fmt.Sprintf("[#ffffff] (%s) - %s ", statusCode[0], inputText)

			tabList.SetItemText(currentListId, mainText, currentId)
		},
	}
	shortcusts.RegisterKeys()

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()

	if cli.Help {
		return
	}

	tabList.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		id := secondaryText
		tab := tabsMap[id]
		currentListId = index

		// Cada que se selecciona un elemento de la lista con "enter", cambia a la pestaña principal y actualiza los valores
		go func() {
			app.QueueUpdateDraw(func() {
				main.LeftPanel.Input.SetText(tab.Input)
				main.LeftPanel.DropDown.SetTitle(tab.DropDownOption)
				main.EditorPanel.Header.SetText(tab.Headers, false)
				main.EditorPanel.QueryParam.SetText(tab.QueryParam, false)
				main.EditorPanel.PathParam.SetText(tab.PathParam, false)
				main.EditorPanel.Variable.SetText(tab.Variables, false)
				main.EditorPanel.Body.SetText(tab.Body, false)

				main.RightPanel.ResponseView.SetText(tab.ResponseBody)
				main.RightPanel.ResponseInfo.SetText(tab.ResponseInfo)

			})
		}()

		mainPage.SwitchToPage("main")
		currentId = id
		shortcusts.ShowTabPage = false
	})

	tabList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			id := uuid.New().String()

			tabsMap[id] = Tabs{
				ID:             id,
				Input:          "",
				DropDownOption: "",
				Headers:        "",
				QueryParam:     "",
				PathParam:      "",
				Variables:      "",
				Body:           "",
				ResponseBody:   "",
				ResponseInfo:   "",
			}

			tabList.AddItem(id, id, 0, nil)
		}

		return event

	})

	main.EditorPanel.Variable.SetText(logic.ReadEnvFile(cli.EnvFilePath), false)

	if err := app.SetRoot(mainPage, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
