package ui

import (
	color "http_client/const/color_wrapper"
	"http_client/ui/template"

	"github.com/rivo/tview"
)

func ResponseInfo() *tview.TextView {
	return template.TextView(color.ColorTextPrimary, "Response")
}
