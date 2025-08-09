package ui

import (
	"http_client/cmd"
	"http_client/internal/clipboard"
	"http_client/logic"
	ui "http_client/ui/components"
	"http_client/ui/events"
	"http_client/ui/layout"
	"http_client/ui/shotcust"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	responseView  *tview.TextView
	responseInfo  *tview.TextView
	app           *tview.Application
	contentToCopy string
)

func StartApp() {

	cli := cmd.Launch()

	showHelpPage := false
	fullScreen := false

	app = tview.NewApplication()
	mainPage := tview.NewPages() // Contiene toda la UI

	main := layout.MainLayout()
	responseView = main.RightPanel.ResponseView
	responseInfo = main.RightPanel.ResponseInfo

	mainLayout := main.Main
	mainPage.AddPage("main", mainLayout, true, true)
	mainPage.AddPage("help", ui.Help(), true, false)
	mainPage.AddPage("history", ui.History(app), true, false)

	requestSender := &events.RequestSender{
		App:           app,
		ResponseView:  responseView,
		ResponseInfo:  responseInfo,
		ContentToCopy: &contentToCopy,
	}

	shortcusts := &shotcust.Shortcuts{
		App:                    app,
		MainPage:               mainPage,
		SwitchPage:             main.EditorPanel.Editor,
		ResponseFlex:           main.RightPanel.Container,
		ShowRequestHistoryPage: showHelpPage,
		ShowHelpPage:           showHelpPage,
		RightPanel:             main.RightPanel.Container,
		ChangeToFullScreen:     fullScreen,

		CopyFn: func() { clipboard.Copy(contentToCopy) },
		ResponseFn: func() {
			requestSender.SendRequest(
				main.LeftPanel.Input,
				main.LeftPanel.DropDown,
				main.EditorPanel.Body,
				main.EditorPanel.Header,
				main.EditorPanel.QueryParam,
				main.EditorPanel.PathParam,
				main.EditorPanel.Variable)
		},
		FocusForm: func() { app.SetFocus(main.Form.Container) },
	}
	shortcusts.RegisterKeys()

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.ContrastBackgroundColor = tcell.ColorWhite.TrueColor()
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack.TrueColor()
	tview.Styles.PrimaryTextColor = tcell.ColorDarkSlateGray.TrueColor()

	if cli.Help {
		return
	}

	main.EditorPanel.Variable.SetText(logic.ReadEnvFile(cli.EnvFilePath), false)

	if err := app.SetRoot(mainPage, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
