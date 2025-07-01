package ui

import "github.com/rivo/tview"

func PathParamsEditor() *tview.TextArea {
	//del/1
	return TextEditor("1,\nid", " > Path Params ")
}
