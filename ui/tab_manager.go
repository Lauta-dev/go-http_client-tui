package ui

import (
	"fmt"
	"strconv"
	"strings"

	color "http_client/const/color_wrapper"
	"http_client/ui/layout"
	"http_client/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

// TODO: Hacer un correcto formateo de los items, si se añade uno ya tenga el estilos "[200] - CustomName"
// Si se edita el texto lo mismo debe tener el mismo estilo

// TODO: También crear opción para eliminar pestaña
// TETAS >>>>> CULOS

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
		ID:             id,
		URL:            "",
		Method:         "",
		MethodID:       0,
		Headers:        "",
		QueryParams:    "",
		PathParams:     "",
		Variables:      "",
		Body:           "",
		ResponseBody:   "",
		ResponseInfo:   "",
		CustomName:     "",
		StatusCodeText: "",
		StatusCode:     0,
		ContentType:    "",
	}
}

func ItemNameFormat(statusCodeText, inputText string) string {
	return fmt.Sprintf(
		"[%s] [%s] - %s ",
		color.ColorTextPrimary.String(),
		statusCodeText,
		inputText,
	)
}

// SaveCurrentTabState guarda el estado actual de la pestaña activa
func (tm *TabManager) SaveCurrentTabState(main *layout.Layout) {
	if tm.appState == nil {
		return
	}

	inputText := main.LeftPanel.Input.GetText()
	methodId, method := main.LeftPanel.DropDown.GetCurrentOption()
	headers := main.EditorPanel.Header.GetText()
	queryParams := main.EditorPanel.QueryParam.GetText()
	pathParams := main.EditorPanel.PathParam.GetText()
	variables := main.EditorPanel.Variable.GetText()
	body := main.EditorPanel.Body.GetText()
	responseBody := main.RightPanel.ResponseView.GetText(true)
	responseInfo := main.RightPanel.ResponseInfo.GetText(true)

	var contentType string = ""
	var statusCodeText string = ""
	var statusCode int = 0

	if responseInfo != "" {
		// [200 OK, text/html] [https://example.com]
		t := strings.SplitN(responseInfo, "\n", 2)

		// [200 OK] [text/html]
		d := strings.SplitN(t[0], ",", 2)

		contentType = strings.TrimSpace(d[1])                         // text/html
		statusCodeText = strings.TrimSpace(d[0])                      // 200 OK
		statusCode, _ = strconv.Atoi(strings.SplitN(d[0], " ", 2)[0]) // 200 "int"
	}

	tabs := tm.appState.tabs
	currentTab := tm.appState.currentTab
	tab, exists := tm.appState.tabs[currentTab]

	if !exists {
		return
	}

	tabs[currentTab] = &Tab{
		ID:             currentTab,
		URL:            inputText,
		Method:         method,
		MethodID:       methodId,
		Headers:        headers,
		QueryParams:    queryParams,
		PathParams:     pathParams,
		Variables:      variables,
		Body:           body,
		ResponseBody:   responseBody,
		ResponseInfo:   responseInfo,
		CustomName:     tab.CustomName,
		StatusCodeText: statusCodeText,
		StatusCode:     statusCode,
		ContentType:    contentType,
	}

	if tab.CustomName != "" {
		f := ItemNameFormat(tab.StatusCodeText, tab.CustomName)
		tm.updateTabListItem(f)
		tm.showRequestInfo(currentTab)
		return
	}

	if inputText == "" && statusCodeText == "" {
		tm.updateTabListItem(" " + currentTab + " ")
		return
	}

	f := ItemNameFormat(statusCodeText, inputText)

	tm.showRequestInfo(currentTab)
	tm.updateTabListItem(f)
}

func (tm *TabManager) showRequestInfo(id string) {

	if tm.appState == nil {
		return
	}

	as := tm.appState
	tabs := tm.appState.tabs

	tab, exists := tabs[id]

	if !exists {
		return
	}

	if tab.URL == "" {
		as.tabInfo.TextArea.SetText("Por favor, haga una Request")
		return
	}

	c, err := utils.ParseUrl(
		tab.Variables,
		tab.URL,
		utils.ParseHeader(tab.QueryParams),
		utils.ParsePathParams(tab.PathParams),
	)

	if err != nil {
		c = tab.URL
	}

	info := fmt.Sprintf("URL: %s\nMétodo: %s\nTipo de contenido: %s\nCódigo: %s\n\n%s",
		c,
		tab.Method,
		tab.ContentType,
		tab.StatusCodeText,
		tab.ResponseBody,
	)
	as.tabInfo.TextArea.SetText(info)
}

// updateTabListItem actualiza el texto mostrado en la lista de pestañas
func (tm *TabManager) updateTabListItem(mainText string) {
	tm.appState.tabList.SetItemText(tm.appState.currentListID, mainText, tm.appState.currentTab)
}

// LoadTabState carga el estado de una pestaña específica
func (tm *TabManager) LoadTabState(tab *Tab, main *layout.Layout) {
	// Limpiar vistas
	tm.appState.responseView.Clear()
	tm.appState.responseInfo.Clear()

	// Cargar datos básicos
	main.LeftPanel.Input.SetText(tab.URL)
	main.LeftPanel.DropDown.SetCurrentOption(tab.MethodID)
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
		tab := tm.appState.tabs[tm.appState.currentTab]
		utils.PrettyStyle(tab.ContentType, []byte(tab.ResponseBody), tm.appState.responseView)
		return
	}

	tm.appState.responseView.SetText(tab.ResponseBody)
}

// CreateNewTab crea una nueva pestaña vacía
func (tm *TabManager) CreateNewTab() {
	id := uuid.New().String()

	tm.appState.tabs[id] = &Tab{
		ID:             id,
		URL:            "",
		Method:         "",
		MethodID:       0,
		Headers:        "",
		QueryParams:    "",
		PathParams:     "",
		Variables:      "",
		Body:           "",
		ResponseBody:   "",
		ResponseInfo:   "",
		CustomName:     "",
		StatusCodeText: "",
		StatusCode:     0,
		ContentType:    "",
	}

	tm.appState.tabList.AddItem(" "+id+" ", id, 0, nil)
}

// SetupTabListHandlers configura los manejadores de eventos para la lista de pestañas
func (tm *TabManager) SetupTabListHandlers(main *layout.Layout) {

	if tm == nil || tm.appState == nil || main == nil {
		return
	}

	list := tm.appState.tabList
	tabs := tm.appState.tabs
	as := tm.appState

	if list == nil {
		return
	}

	var id int
	var tabId string

	// Manejador de selección de pestaña
	list.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		tabID := secondaryText
		tab, exists := tabs[tabID]
		if !exists {
			return
		}

		as.currentListID = index
		as.currentTab = tabID

		// Cargar estado de la pestaña de forma asíncrona
		go func() {
			as.app.QueueUpdateDraw(func() {
				tm.LoadTabState(tab, main)
			})
		}()

		as.mainPage.SwitchToPage("main")
		as.showTabPage = false

		tabId = secondaryText
		id = index
	})

	// Cambia la información de la Request al cambiar de item
	list.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		tm.showRequestInfo(secondaryText)
		tabId = secondaryText
		id = index
	})

	// Manejador de entrada de teclado
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			tm.CreateNewTab()
		case 'e':

			if tabId == "" {
				return nil
			}

			tabInfo := as.tabInfo
			tabInfo.DetailsPage.SwitchToPage("edit-item")
			as.app.SetFocus(as.tabInfo.Input)

			tabInfo.DetailsPage.SetTitle(" > Edición | 'ENTER' para asignar 'ESC' para salir ")
			tabInfo.Input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyEnter:
					newText := tabInfo.Input.GetText()

					if newText == "" {
						return nil
					}
					tab := tabs[tabId]
					mainText := ItemNameFormat(tab.StatusCodeText, newText)

					tab.CustomName = newText

					tabInfo.List.SetItemText(id, mainText, tabId)
					tabInfo.DetailsPage.SwitchToPage("info")
					as.app.SetFocus(tabInfo.List)
					tabInfo.Input.SetText("")

				case tcell.KeyEsc:
					tabInfo.DetailsPage.SwitchToPage("info")
					tabInfo.DetailsPage.SetTitle(" > Info | 'E' para editar item ")
					as.app.SetFocus(tabInfo.List)
					tabInfo.Input.SetText("")
				}
				return event
			})

		}
		return event
	})
}
