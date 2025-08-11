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
	App           *tview.Application
	ResponseView  *tview.TextView
	ResponseInfo  *tview.TextView
	ContentToCopy *string // puntero para actualizar desde la función

}

func (ui *UIController) UpdateComponent(
	httpMethod string,
	headers map[string]string,
	queryParams map[string]string,
	params []string,
	body string,
	userUrl string, saveRequest bool) {

	app := ui.App
	responseView := ui.ResponseView
	responseInfo := ui.ResponseInfo

	app.QueueUpdateDraw(func() {
		responseView.Clear()
		responseInfo.Clear()

		loadingFormat := fmt.Sprintf("[%s::b]%s[%s::B]",
			color.ColorTextSecondary.String(),
			"Cargando...",
			color.ColorPrimary.String(),
		)

		fmt.Fprintln(responseView, loadingFormat)
		fmt.Fprintln(responseInfo, loadingFormat)
	})

	res, err := logic.Fetching(userUrl, httpMethod, headers, queryParams, params, body)

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

		if saveRequest {
			format := fmt.Sprintf("%s, %s \nURL: %s", utils.StatusCodesColors(res.Status), res.ContentType, res.UserUrl)
			code := strings.Split(res.Status, " ")[0]

			err := logic.SaveItems(res.UserUrl, code, res.ContentType, res.Body, httpMethod)

			if err != nil {
				fmt.Fprintf(responseInfo, "[red]%s", err.Error())
			} else {
				fmt.Fprintln(responseInfo, format)
			}
		}

		display, err := utils.PrettyStyle(res.ContentType, []byte(res.Body), responseView)
		ui.ContentToCopy = &display

		if err != nil {
			responseView.Clear()
			fmt.Fprintf(responseView, "Error al colorar la salída\n%s\n\n%s", err.Error(), display)
		}

	})
}
