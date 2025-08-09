package layout

import (
	component "http_client/ui/components"

	"github.com/rivo/tview"
)

type Left struct {
	DropDown  *tview.DropDown
	Input     *tview.InputField
	Container *tview.Flex
}

func LeftPanel(editorPanel *tview.Pages) *Left {
	form, dropdown, input := component.Form()
	panel := tview.NewFlex().SetDirection(tview.FlexRow)
	panel.AddItem(form, 0, 1, true)
	panel.AddItem(editorPanel, 0, 1, false)

	return &Left{
		DropDown:  dropdown,
		Input:     input,
		Container: panel,
	}
}
