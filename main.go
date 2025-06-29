package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Objectivo: Usando tview crear una TUI para el cliente HTTP

/*
|------------------------------------------------------------------------------------------|
| Lista de Verbos                      |     STATUS CODE: 200 (OK)                         |
| 			                               |                                                   |
| Input URL                            |---------------------------------------------------|
|--------------------------------------|                                                   |
|                                      |                                                   |
| Botones                              |                                                   |
| [Headers] [Body] [Params]            |                                                   |
|                                      |           JSON|HTML|Text|Lo que sea               |
|                                      |                                                   |
|                                      |                                                   |
|                                      |                                                   |
|                                      |                                                   |
|                                      |                                                   |
|                                      |                                                   |
|                                      |                                                   |
|                                      |																									 |
|                                      |                                                   |
|------------------------------------------------------------------------------------------|
*/

var (
	view       *tview.TextView
	footerInfo *tview.Frame
	app        *tview.Application

	// Variables para la request
	urlTo        string
	selectedVerb string
	body         string
	header       map[string]string
	params       []string          // :id
	queryParams  map[string]string // ?id=

	//
	responseStatus string // ex 200, OK
)

func getItems() string {
	pathParam := ""
	for _, v := range params {
		pathParam += "/" + v
	}

	urlTo = urlTo + pathParam

	u, err := url.Parse(urlTo)

	if err != nil {
		return "Error al parsear la URL " + err.Error()
	}

	q := u.Query()

	for k, v := range queryParams {
		if q.Get(k) == "" {
			q.Add(k, v)

		} else {
			q.Set(k, v)
		}
	}

	u.RawQuery = q.Encode()
	finalUrl := u.String()

	var post any
	var isJson = true

	req, err := http.NewRequest(selectedVerb, finalUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "Error al crear peticion " + err.Error()
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Error al enviar peticion " + err.Error()
	}

	responseStatus = resp.Status

	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "Error al leer respuesta " + err.Error()
	}

	switch bytes[0] {
	case '{': // Por si es un Object
		post = make(map[string]any, 0)
	case '[': // Por si es un Array
		post = make([]map[string]any, 0)
	default:
		isJson = false
	}

	if !isJson {
		return string(bytes)
	}

	erre := json.Unmarshal(bytes, &post)

	if erre != nil {
		return "Error al parsear JSON"
	}

	ident, err := json.MarshalIndent(post, "", " ")

	return string(ident)
}

/*
 ----------
 | 200 OK |
 ----------

 -----
 Respuesta
*/

func updateTime() {
	app.QueueUpdateDraw(func() {
		view.Clear()
		fmt.Fprintln(view, "Cargando...")
	})
	sal := getItems()

	app.QueueUpdateDraw(func() {
		view.Clear()
		response := StatusCodesColors(responseStatus) + "\n\n" + string(sal)
		fmt.Fprintln(view, response)
	})

}

func main() {
	httpVerb := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
	}
	app = tview.NewApplication()
	flex := tview.NewFlex()
	form := tview.NewForm()
	formInput := tview.NewInputField().SetLabel("URL")
	formInput.SetText("http://localhost:4000")
	bodyContent := tview.NewTextArea().SetPlaceholder("{'name': 'lautaro'}")
	mainPage := tview.NewPages()
	switchPage := tview.NewPages()
	headerBody := A()
	qParams := QueryParams()
	pathParam := PathParams()
	der := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(switchPage, 0, 2, false)
	helpPage := Help()
	showHelpPage := false

	// Aplicar color global de selección (tema de la app)
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()

	dropdown := tview.NewDropDown().SetLabel("Método")
	dropdown.SetOptions(httpVerb, nil)
	form.AddFormItem(dropdown)

	form.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	form.SetFieldBackgroundColor(tcell.ColorDarkSlateGray.TrueColor())
	form.SetTitleColor(tcell.ColorWhite.TrueColor())
	form.SetFieldTextColor(tcell.ColorWhite.TrueColor())
	form.SetBorder(true)
	form.SetTitle(" Request ")
	form.SetTitleAlign(tview.AlignLeft)
	form.SetButtonBackgroundColor(tcell.ColorBlack.TrueColor())

	form.AddButton("Send", func() {
		urlTo = formInput.GetText()
		_, v := dropdown.GetCurrentOption()
		selectedVerb = v
		body = bodyContent.GetText()
		switchPage.SwitchToPage("response")
		header = ParseHeader(headerBody.GetText())
		queryParams = ParseHeader(qParams.GetText())
		params = ParseParams(pathParam.GetText())

		go updateTime()
	})

	form.AddButton("Response", func() {
		switchPage.SwitchToPage("response")
	})

	form.AddFormItem(formInput)

	view = tview.NewTextView()
	view.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	view.SetTextColor(tcell.ColorWhite.TrueColor())
	view.SetBorder(true)
	view.SetTitle(" Response ")
	view.SetTitleAlign(tview.AlignLeft)
	view.SetDynamicColors(true)
	statusView := tview.NewFrame(der).SetBorders(0, 0, 0, 0, 0, 0).AddText(responseStatus, true, tview.AlignLeft, tcell.ColorWhite)

	switchPage.
		AddPage("body", bodyContent, true, false).
		AddPage("response", view, true, false).
		AddPage("header", headerBody, true, false).
		AddPage("qp", qParams, true, false).
		AddPage("pp", pathParam, true, false)

	// Parte izq
	flex.AddItem(
		tview.NewFlex().SetDirection(tview.FlexRow).AddItem(Buttons(switchPage), 0, 1, false).
			AddItem(form, 0, 5, false),

		0,
		1,
		false)

	// Parte der
	flex.AddItem(statusView, 0, 1, false)

	fullScreen := false
	der.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'f':
				// Hiciste Alt+F
				// Alternar fullscreen, por ejemplo
				fullScreen = !fullScreen
				der.SetFullScreen(fullScreen)
				return nil // bloquear evento si querés
			}
		}
		return event

	})

	dropdown.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	dropdown.SetFieldBackgroundColor(tcell.ColorBlack.TrueColor())
	dropdown.SetTitleColor(tcell.ColorWhite.TrueColor())
	dropdown.SetFieldTextColor(tcell.ColorWhite.TrueColor())
	dropdown.SetFieldStyle(tcell.StyleDefault.Blink(true))

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
		return event
	})

	mainPage.AddPage("main", flex, true, true)
	mainPage.AddPage("help", helpPage, true, false)

	if err := app.SetRoot(mainPage, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
