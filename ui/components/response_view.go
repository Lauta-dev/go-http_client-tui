package ui

import (
	color "http_client/const/color_wrapper"
	"http_client/ui/template"

	"github.com/rivo/tview"
)

func ResponseView() *tview.TextView {
	return template.TextView(color.ColorTextPrimary, "Response")
}
