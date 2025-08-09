package layout

import (
	"github.com/rivo/tview"
)

type Layout struct {
	Main        *tview.Flex
	EditorPanel *Editor
	Form        *Left
	RightPanel  *ResponseWindow
	LeftPanel   *Left
}

func MainLayout() *Layout {
	mainLayout := tview.NewFlex() // Contiene los flexes "leftPanel" y "rightPanel"
	editorPanel := EditorPanel()
	editor := editorPanel.Editor
	/*
		bodyEditor := editorPanel.Body
		headerEditor := editorPanel.Header
		queryParamEditor := editorPanel.QueryParam
		pathParamEditor := editorPanel.PathParam*/

	rightPanel := NewResponseWindow()
	leftPanel := LeftPanel(editor)
	/*input := form.Input
	dropdown := form.DropDown*/

	mainLayout.AddItem(leftPanel.Container, 0, 1, false)
	mainLayout.AddItem(rightPanel.Container, 0, 1, false)

	return &Layout{
		Main:        mainLayout,
		EditorPanel: editorPanel,
		Form:        leftPanel,
		LeftPanel:   leftPanel,
		RightPanel:  rightPanel,
	}
}
