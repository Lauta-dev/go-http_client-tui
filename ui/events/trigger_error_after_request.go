package events

import (
	"fmt"

	"github.com/rivo/tview"
)

type RequestError struct {
	App          *tview.Application
	ResponseView *tview.TextView
	ResponseInfo *tview.TextView
}

func (requestError *RequestError) TriggerErrorAfterRequest(err error) {
	app := requestError.App
	responseInfo := requestError.ResponseInfo
	responseView := requestError.ResponseView

	app.QueueUpdateDraw(func() {
		responseInfo.Clear()
		responseView.Clear()

		fmt.Fprintln(responseView, err.Error())
	})
}
