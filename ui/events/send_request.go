package events

import (
	"github.com/rivo/tview"
	"http_client/utils"
)

type RequestSender struct {
	App          *tview.Application
	ResponseView *tview.TextView
	ResponseInfo *tview.TextView

	ContentToCopy  *string
	StatusCodeText string // Ej OK 200
	StatusCode     int    // Ej 200
}

func (rs *RequestSender) SendRequest(
	formInput *tview.InputField,
	dropdown *tview.DropDown,
	bodyContent *tview.TextArea,
	headerPage *tview.TextArea,
	queryParamPage *tview.TextArea,
	pathParamPage *tview.TextArea,
	varr *tview.TextArea,
	saveRequest bool,
) {
	_, selected := dropdown.GetCurrentOption()
	body := bodyContent.GetText()
	header := utils.ParseHeaders(headerPage.GetText(), utils.ParseQueryParams(varr.GetText()))
	queryParams := utils.ParseQueryParams(queryParamPage.GetText())
	params := utils.ParsePathParams(pathParamPage.GetText())

	url, err := utils.ReplaceVariablesInURL(formInput.GetText(), utils.ParseQueryParams(varr.GetText()))

	if err != nil {
		triggerErrorAfterRequest := &RequestError{
			App:          rs.App,
			ResponseView: rs.ResponseView,
			ResponseInfo: rs.ResponseInfo,
		}
		go triggerErrorAfterRequest.TriggerErrorAfterRequest(err)
		return
	}

	uiController := &UIController{
		App:           rs.App,
		ResponseView:  rs.ResponseView,
		ResponseInfo:  rs.ResponseInfo,
		ContentToCopy: &body,
	}

	go uiController.UpdateComponent(selected, header, queryParams, params, body, url, saveRequest)
}
