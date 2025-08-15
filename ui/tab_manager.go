package ui

import (
	"fmt"
	"http_client/ui/layout"
	"http_client/utils"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

// TabManager maneja todas las operaciones relacionadas con pestañas
type TabManager struct {
	appState *AppState
}

// NewTabManager crea un nuevo manejador de pestañas
func NewTabManager(appState *AppState) *TabManager {
	return &TabManager{
		appState: appState,
	}
}

// CreateInitialTab crea la primera pestaña de la aplicación
func (tm *TabManager) CreateInitialTab(id string) {
	tm.appState.tabs[id] = &Tab{
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

// SaveCurrentTabState guarda el estado actual de la pestaña activa
func (tm *TabManager) SaveCurrentTabState(main *layout.Layout) {
	inputText := main.LeftPanel.Input.GetText()
	method := main.LeftPanel.DropDown.GetTitle()
	headers := main.EditorPanel.Header.GetText()
	queryParams := main.EditorPanel.QueryParam.GetText()
	pathParams := main.EditorPanel.PathParam.GetText()
	variables := main.EditorPanel.Variable.GetText()
	body := main.EditorPanel.Body.GetText()
	responseBody := main.RightPanel.ResponseView.GetText(true)
	responseInfo := main.RightPanel.ResponseInfo.GetText(true)

	tm.appState.tabs[tm.appState.currentTab] = &Tab{
		ID:           tm.appState.currentTab,
		URL:          inputText,
		Method:       method,
		Headers:      headers,
		QueryParams:  queryParams,
		PathParams:   pathParams,
		Variables:    variables,
		Body:         body,
		ResponseBody: responseBody,
		ResponseInfo: responseInfo,
	}

	tm.updateTabListItem(responseInfo, inputText)
}

// updateTabListItem actualiza el texto mostrado en la lista de pestañas
func (tm *TabManager) updateTabListItem(responseInfo, inputText string) {
	statusCode := strings.SplitN(responseInfo, ",", 2)
	if len(statusCode) > 0 {
		mainText := fmt.Sprintf("[#ffffff] [%s] - %s", statusCode[0], inputText)
		tm.appState.tabList.SetItemText(tm.appState.currentListID, mainText, tm.appState.currentTab)
	}
}

// LoadTabState carga el estado de una pestaña específica
func (tm *TabManager) LoadTabState(tab *Tab, main *layout.Layout) {
	// Limpiar vistas
	tm.appState.responseView.Clear()
	tm.appState.responseInfo.Clear()

	// Cargar datos básicos
	main.LeftPanel.Input.SetText(tab.URL)
	main.LeftPanel.DropDown.SetTitle(tab.Method)
	main.EditorPanel.Header.SetText(tab.Headers, false)
	main.EditorPanel.QueryParam.SetText(tab.QueryParams, false)
	main.EditorPanel.PathParam.SetText(tab.PathParams, false)
	main.EditorPanel.Variable.SetText(tab.Variables, false)
	main.EditorPanel.Body.SetText(tab.Body, false)

	// Cargar información de respuesta
	tm.loadResponseInfo(tab, main)
	tm.loadResponseBody(tab)
}

// loadResponseInfo carga la información de la respuesta HTTP
func (tm *TabManager) loadResponseInfo(tab *Tab, main *layout.Layout) {
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

// loadResponseBody carga el cuerpo de la respuesta HTTP
func (tm *TabManager) loadResponseBody(tab *Tab) {
	if tab.ResponseBody == "" {
		tm.appState.responseView.SetText("")
		return
	}

	if tab.ResponseInfo != "" {
		contentTypeParts := strings.SplitN(tab.ResponseInfo, ",", 2)
		if len(contentTypeParts) >= 2 {
			contentType := strings.TrimSpace(contentTypeParts[1])
			utils.PrettyStyle(contentType, []byte(tab.ResponseBody), tm.appState.responseView)
			return
		}
	}

	tm.appState.responseView.SetText(tab.ResponseBody)
}

// CreateNewTab crea una nueva pestaña vacía
func (tm *TabManager) CreateNewTab() {
	id := uuid.New().String()

	tm.appState.tabs[id] = &Tab{
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

	tm.appState.tabList.AddItem(id, id, 0, nil)
}

// SetupTabListHandlers configura los manejadores de eventos para la lista de pestañas
func (tm *TabManager) SetupTabListHandlers(main *layout.Layout) {
	// Manejador de selección de pestaña
	tm.appState.tabList.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		tabID := secondaryText
		tab, exists := tm.appState.tabs[tabID]
		if !exists {
			return
		}

		tm.appState.currentListID = index
		tm.appState.currentTab = tabID

		// Cargar estado de la pestaña de forma asíncrona
		go func() {
			tm.appState.app.QueueUpdateDraw(func() {
				tm.LoadTabState(tab, main)
			})
		}()

		tm.appState.mainPage.SwitchToPage("main")
		tm.appState.showTabPage = false
	})

	// Manejador de entrada de teclado
	tm.appState.tabList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			tm.CreateNewTab()
		}
		return event
	})
}
