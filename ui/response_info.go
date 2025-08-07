package ui

import (
	"github.com/rivo/tview"
	color "http_client/const/color_wrapper"
)

func ResponseInfo() *tview.TextView {
	responseInfo := tview.NewTextView()
	responseInfo.SetTextColor(color.ColorTextSecondary)
	responseInfo.SetBorder(true)
	responseInfo.SetTitle(" > Response Information ")
	responseInfo.SetTitleAlign(tview.AlignLeft)
	responseInfo.SetDynamicColors(true)

	return responseInfo
}
