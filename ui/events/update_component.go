package events

import (
	"fmt"
	color "http_client/const/color_wrapper"
	"http_client/logic"
	"http_client/utils"
	"strings"

	"github.com/rivo/tview"
)

type UIController struct {
	App          *tview.Application
	ResponseView *tview.TextView
	ResponseInfo *tview.TextView

	ContentToCopy  *string // Body de la Request
	StatusCodeText string  // Ej OK 200
	StatusCode     int     // Ej 200
}

func (ui *UIController) clearResponses() {
	ui.ResponseView.Clear()
	ui.ResponseInfo.Clear()
}

func (ui *UIController) showLoading() {
	ui.App.QueueUpdateDraw(func() {
		ui.clearResponses()

		loadingFormat := fmt.Sprintf("[%s::b]%s[%s::B]",
			color.ColorTextSecondary.String(),
			"Cargando...",
			color.ColorPrimary.String(),
		)

		fmt.Fprintln(ui.ResponseView, loadingFormat)
		fmt.Fprintln(ui.ResponseInfo, loadingFormat)
	})
}

func (ui *UIController) showError(err error) {
	ui.App.QueueUpdateDraw(func() {
		ui.clearResponses()
		fmt.Fprintln(ui.ResponseView, err.Error())
	})
}

func (ui *UIController) displayResponse(res logic.Fetch, saveRequest bool, method string) {
	ui.App.QueueUpdateDraw(func() {
		ui.clearResponses()

		format := utils.ResponseInfoFormat(res.ContentType, res.URL, res.StatusCodeText)
		code := strings.Split(res.StatusCodeText, " ")[0]

		if saveRequest {
			if err := logic.SaveItems(res.URL, code, res.ContentType, res.Body, method); err != nil {
				fmt.Fprintf(ui.ResponseInfo, "[red]%s", err.Error())
			}
		}

		fmt.Fprintln(ui.ResponseInfo, format)

		display, err := utils.PrettyStyle(res.ContentType, []byte(res.Body), ui.ResponseView)
		ui.ContentToCopy = &display
		ui.StatusCodeText = res.StatusCodeText
		ui.StatusCode = res.StatusCode

		if err != nil {
			ui.ResponseView.Clear()
			fmt.Fprintf(ui.ResponseView, "Error al colorear la salida\n%s\n\n%s", err.Error(), display)
		}
	})
}

func (ui *UIController) UpdateComponent(
	httpMethod string,
	headers map[string]string,
	queryParams map[string]string,
	params []string,
	body string,
	userUrl string, saveRequest bool,
) {

	ui.showLoading()

	res, err := logic.Fetching(userUrl, httpMethod, headers, queryParams, params, body)

	if err != nil {
		ui.showError(err)
		return
	}

	ui.displayResponse(res, saveRequest, httpMethod)
}
