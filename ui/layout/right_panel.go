package layout

import (
	component "http_client/ui/components"

	"github.com/rivo/tview"
)

type ResponseWindow struct {
	Container    *tview.Flex
	ResponseView *tview.TextView
	ResponseInfo *tview.TextView
}

func NewResponseWindow() *ResponseWindow {
	responseView := component.ResponseView()
	responseInfo := component.ResponseInfo()

	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	rightPanel.AddItem(responseView, 0, 8, false)
	rightPanel.AddItem(responseInfo, 0, 1, false)

	return &ResponseWindow{
		Container:    rightPanel,
		ResponseView: responseView,
		ResponseInfo: responseInfo,
	}
}
