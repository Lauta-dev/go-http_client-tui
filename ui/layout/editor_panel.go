package layout

import (
	component "http_client/ui/components"

	"github.com/rivo/tview"
)

type Editor struct {
	Body       *tview.TextArea
	Header     *tview.TextArea
	QueryParam *tview.TextArea
	PathParam  *tview.TextArea
	Variable   *tview.TextArea

	Editor *tview.Pages
}

func EditorPanel() *Editor {
	body := component.BodyEditor()
	header := component.HeaderEditor()
	queryParam := component.QueryParamsEditor()
	pathParam := component.PathParamsEditor()
	variable := component.VariableEditor()

	editor := tview.NewPages()
	editor.
		AddPage("body", body, true, true).
		AddPage("header", header, true, false).
		AddPage("qp", queryParam, true, false).
		AddPage("pp", pathParam, true, false).
		AddPage("var", variable, true, false)

	return &Editor{
		Body:       body,
		Header:     header,
		QueryParam: queryParam,
		PathParam:  pathParam,
		Variable:   variable,

		Editor: editor,
	}
}
