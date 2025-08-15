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
	"http_client/utils"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
	"github.com/rivo/tview"
)

// Tab representa una pestaña con toda la información de una petición HTTP
type Tab struct {
	ID           string
	URL          string
	Method       string
	Headers      string
	QueryParams  string
	PathParams   string
	Variables    string
	Body         string
	ResponseBody string
	ResponseInfo string
}

// AppState mantiene el estado global de la aplicación
type AppState struct {
	app           *tview.Application
	mainPage      *tview.Pages
	responseView  *tview.TextView
	responseInfo  *tview.TextView
	contentToCopy string

	// Estado de pestañas
	currentTab    string
	currentListID int
	tabs          map[string]*Tab
	tabList       *tview.List

	// Estado de páginas
	showHelpPage bool
	showTabPage  bool // true -> mostrar layout, false -> mostrar lista de pestañas
	fullScreen   bool
}

// Guarda la pestaña actual
func (as *AppState) saveCurrentTabState(main *layout.Layout) {
	id := as.currentTab

	url := main.LeftPanel.Input.GetText()
	method := main.LeftPanel.DropDown.GetTitle()

	headers := main.EditorPanel.Header.GetText()
	queryParams := main.EditorPanel.QueryParam.GetText()
	pathParams := main.EditorPanel.PathParam.GetText()
	variables := main.EditorPanel.Variable.GetText()
	body := main.EditorPanel.Body.GetText()

	responseBody := main.RightPanel.ResponseView.GetText(true)
	responseInfo := main.RightPanel.ResponseInfo.GetText(true)

	as.tabs[id] = &Tab{
		ID:           id,
		URL:          url,
		Method:       method,
		Headers:      headers,
		QueryParams:  queryParams,
		PathParams:   pathParams,
		Variables:    variables,
		Body:         body,
		ResponseBody: responseBody,
		ResponseInfo: responseInfo,
	}

	as.updateTabListItem(responseInfo, url)

}

// updateTabListItem actualiza el nombre de la pestaña actual
func (as *AppState) updateTabListItem(responseInfo, inputText string) {
	statusCode := strings.SplitN(responseInfo, ",", 2)
	if len(statusCode) > 0 {
		mainText := fmt.Sprintf("[#ffffff] [%s] - %s", statusCode[0], inputText)
		as.tabList.SetItemText(as.currentListID, mainText, as.currentTab)
	}
}

// loadTabState carga el estado de la pestaña seleccionada
func (as *AppState) loadTabState(tab *Tab, main *layout.Layout) {
	// Limpiar vistas
	as.responseView.Clear()
	as.responseInfo.Clear()

	// Cargar datos básicos
	main.LeftPanel.Input.SetText(tab.URL)
	main.LeftPanel.DropDown.SetTitle(tab.Method)
	main.EditorPanel.Header.SetText(tab.Headers, false)
	main.EditorPanel.QueryParam.SetText(tab.QueryParams, false)
	main.EditorPanel.PathParam.SetText(tab.PathParams, false)
	main.EditorPanel.Variable.SetText(tab.Variables, false)
	main.EditorPanel.Body.SetText(tab.Body, false)

	// Cargar información de respuesta
	as.loadResponseInfo(tab, main)
	as.loadResponseBody(tab)
}

// TODO: Hacer mas accesible el conten type y status code

// loadResponseInfo carga la información de la respuesta HTTP
func (as *AppState) loadResponseInfo(tab *Tab, main *layout.Layout) {
	if tab.ResponseInfo == "" {
		return
	}

	statusAndContentType := strings.SplitN(tab.ResponseInfo, ",", 2)
	if len(statusAndContentType) >= 2 {
		contentType := strings.TrimSpace(strings.Split(statusAndContentType[1], "\n")[0])
		statusCode := strings.TrimSpace(statusAndContentType[0])

		formattedInfo := utils.ResponseInfoFormat(contentType, tab.URL, statusCode)
		main.RightPanel.ResponseInfo.SetText(formattedInfo)
	}
}

// TODO: Hacer mas accesible el conten type y status code

// loadResponseBody carga el cuerpo de la respuesta HTTP
func (as *AppState) loadResponseBody(tab *Tab) {
	if tab.ResponseBody == "" {
		as.responseView.SetText("")
		return
	}

	if tab.ResponseInfo != "" {
		contentTypeParts := strings.SplitN(tab.ResponseInfo, ",", 2)
		if len(contentTypeParts) >= 2 {
			contentType := strings.TrimSpace(contentTypeParts[1])
			utils.PrettyStyle(contentType, []byte(tab.ResponseBody), as.responseView)
			return
		}
	}

	as.responseView.SetText(tab.ResponseBody)
}

func (as *AppState) createNewTab() {
	id := uuid.New().String()

	as.tabs[id] = &Tab{
		ID:           id,
		URL:          "",
		Method:       "",
		Headers:      "",
		QueryParams:  "",
		PathParams:   "",
		Variables:    "",
		Body:         "",
		ResponseBody: "",
		ResponseInfo: "",
	}

	as.tabList.AddItem(id, id, 0, nil)
}

// setupRequestSender configura el enviador de peticiones
func setupRequestSender(appState *AppState, main *layout.Layout) *events.RequestSender {
	return &events.RequestSender{
		App:           appState.app,
		ResponseView:  appState.responseView,
		ResponseInfo:  appState.responseInfo,
		ContentToCopy: &appState.contentToCopy,
	}
}

// setupShortcuts configura los atajos de teclado
func setupShortcuts(appState *AppState, main *layout.Layout, cli *cmd.CliOptions) {
	requestSender := setupRequestSender(appState, main)

	shortcuts := &shotcust.Shortcuts{
		App:                    appState.app,
		MainPage:               appState.mainPage,
		SwitchPage:             main.EditorPanel.Editor,
		ResponseFlex:           main.RightPanel.Container,
		ShowRequestHistoryPage: appState.showHelpPage,
		ShowHelpPage:           appState.showHelpPage,
		RightPanel:             main.RightPanel.Container,
		ChangeToFullScreen:     appState.fullScreen,
		ShowTabPage:            &appState.showTabPage,

		CopyFn: func() {
			clipboard.Copy(appState.contentToCopy)
		},
		ResponseFn: func() {
			requestSender.SendRequest(
				main.LeftPanel.Input,
				main.LeftPanel.DropDown,
				main.EditorPanel.Body,
				main.EditorPanel.Header,
				main.EditorPanel.QueryParam,
				main.EditorPanel.PathParam,
				main.EditorPanel.Variable,
				cli.ActHistory,
			)
		},
		FocusForm: func() {
			appState.app.SetFocus(main.Form.Container)
		},
		SaveStateFn: func() {
			appState.saveCurrentTabState(main)
		},
	}

	shortcuts.RegisterKeys()
}

// setupTabListHandlers configura los manejadores de eventos para la lista de pestañas
func setupTabListHandlers(appState *AppState, main *layout.Layout) {
	// Manejador de selección de pestaña
	appState.tabList.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		tabID := secondaryText
		tab, exists := appState.tabs[tabID]
		if !exists {
			return
		}

		appState.currentListID = index
		appState.currentTab = tabID

		// Cargar estado de la pestaña de forma asíncrona
		go func() {
			appState.app.QueueUpdateDraw(func() {
				appState.loadTabState(tab, main)
			})
		}()

		appState.showTabPage = false
		appState.mainPage.SwitchToPage("main")
	})

	// Manejador de entrada de teclado
	appState.tabList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			appState.createNewTab()
		}
		return event
	})
}

func setupStyles() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()
}

// NewAppState crea una nueva instancia del estado de la aplicación
func NewAppState() *AppState {
	firstID := uuid.New().String()

	return &AppState{
		app:           tview.NewApplication(),
		tabs:          make(map[string]*Tab),
		currentTab:    firstID,
		currentListID: 0,
		showHelpPage:  false,
		showTabPage:   false,
		fullScreen:    false,
	}
}

func (as *AppState) createInitialTab(id string) {
	as.tabs[id] = &Tab{
		ID:           id,
		URL:          "",
		Method:       "",
		Headers:      "",
		QueryParams:  "",
		PathParams:   "",
		Variables:    "",
		Body:         "",
		ResponseBody: "",
		ResponseInfo: "",
	}
}

func StartApp() {
	cli := cmd.Launch()
	if cli.Help {
		return
	}

	// Crear estado de la aplicación
	appState := NewAppState()

	// Configurar layout principal
	main := layout.MainLayout()
	appState.responseView = main.RightPanel.ResponseView
	appState.responseInfo = main.RightPanel.ResponseInfo

	// Configurar páginas
	appState.mainPage = tview.NewPages()
	appState.tabList = component.TabsList()

	// Configurar pestaña inicial
	appState.createInitialTab(appState.currentTab)
	appState.tabList.AddItem("Not Found", appState.currentTab, 0, nil)

	// Agregar páginas
	appState.mainPage.AddPage("main", main.Main, true, true)
	appState.mainPage.AddPage("help", component.Help(), true, false)
	appState.mainPage.AddPage("tab", appState.tabList, true, false)

	if cli.ActHistory {
		appState.mainPage.AddPage("history", component.History(appState.app), true, false)
	}

	// Configurar eventos y atajos
	setupRequestSender(appState, main)
	setupShortcuts(appState, main, &cli)
	setupTabListHandlers(appState, main)

	// Configurar variables de entorno
	main.EditorPanel.Variable.SetText(logic.ReadEnvFile(cli.EnvFilePath), false)

	// Configurar estilos y ejecutar
	setupStyles()

	if err := appState.app.SetRoot(appState.mainPage, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
