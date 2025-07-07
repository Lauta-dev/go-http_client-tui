package ui

import (
	"github.com/rivo/tview"
	colors "http_client/const"
)

func ResponseInfo() *tview.TextView {
	responseInfo := tview.NewTextView()
	responseInfo.SetTextColor(colors.ColorTextSecondary)
	responseInfo.SetBorder(true)
	responseInfo.SetTitle(" > Response Information ")
	responseInfo.SetTitleAlign(tview.AlignLeft)
	responseInfo.SetDynamicColors(true)

	return responseInfo
}
