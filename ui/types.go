package ui

import (
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
	showTabPage  bool
	fullScreen   bool
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

// SetShowTabPage establece el estado de la página de pestañas
func (as *AppState) SetShowTabPage(show bool) {
	as.showTabPage = show
}

// GetShowTabPage obtiene el estado actual de la página de pestañas
func (as *AppState) GetShowTabPage() bool {
	return as.showTabPage
}

// GetApp retorna la aplicación tview
func (as *AppState) GetApp() *tview.Application {
	return as.app
}

// GetMainPage retorna la página principal
func (as *AppState) GetMainPage() *tview.Pages {
	return as.mainPage
}

// SetMainPage establece la página principal
func (as *AppState) SetMainPage(page *tview.Pages) {
	as.mainPage = page
}

// SetResponseViews establece las vistas de respuesta
func (as *AppState) SetResponseViews(responseView, responseInfo *tview.TextView) {
	as.responseView = responseView
	as.responseInfo = responseInfo
}

// SetTabList establece la lista de pestañas
func (as *AppState) SetTabList(tabList *tview.List) {
	as.tabList = tabList
}

// setupStyles configura los estilos de la aplicación
func SetupStyles() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()
}
