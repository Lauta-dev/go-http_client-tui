package shotcust

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Shortcuts struct {
	App          *tview.Application
	SwitchPage   *tview.Pages
	ResponseFlex *tview.Flex
	MainPage     *tview.Pages
	RightPanel   *tview.Flex
	CopyFn       func()
	FocusForm    func()
	ResponseFn   func()
	SaveStateFn  func()

	ShowHelpPage           bool
	ShowRequestHistoryPage bool
	ChangeToFullScreen     bool
	ShowTabPage            bool
}

func (s *Shortcuts) RegisterKeys() {
	s.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {

			// Cambio entre ventanas
			switch event.Rune() {
			case 'r':
				s.SwitchPage.SwitchToPage("response")
				s.App.SetFocus(s.SwitchPage)
			case 'p':
				s.SwitchPage.SwitchToPage("pp")
				s.App.SetFocus(s.SwitchPage)
			case 'q':
				s.SwitchPage.SwitchToPage("qp")
				s.App.SetFocus(s.SwitchPage)
			case 'h':
				s.SwitchPage.SwitchToPage("header")
				s.App.SetFocus(s.SwitchPage)
			case 'b':
				s.SwitchPage.SwitchToPage("body")
				s.App.SetFocus(s.SwitchPage)
			case 'j':
				s.SwitchPage.SwitchToPage("var")
				s.App.SetFocus(s.SwitchPage)
			}
		}
		return event
	})

	s.MainPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {

		// Cambiar a vista de ayuda
		case tcell.KeyF1:
			if s.ShowHelpPage {
				s.ShowHelpPage = false
				s.MainPage.SwitchToPage("main")
			} else {
				s.ShowHelpPage = true
				s.MainPage.SwitchToPage("help")
			}

		// Cambiar a vista de historial de request
		case tcell.KeyF2:
			if s.ShowRequestHistoryPage {
				s.ShowRequestHistoryPage = false
				s.MainPage.SwitchToPage("main")
			} else {
				s.ShowRequestHistoryPage = true
				s.MainPage.SwitchToPage("history")
			}

		case tcell.KeyF3:
			if s.ShowTabPage {
				s.ShowTabPage = false
				s.MainPage.SwitchToPage("main")
			} else {
				s.ShowTabPage = true
				s.MainPage.SwitchToPage("tab")
				s.SaveStateFn()
			}
		}

		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'd':
				s.App.SetFocus(s.ResponseFlex)
				s.ResponseFn()

			case 'n':
				s.CopyFn()

			case 'i':
				s.FocusForm()
			}
		}
		return event
	})

	s.RightPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Modifiers() == tcell.ModAlt {
			switch event.Rune() {
			case 'f':
				s.ChangeToFullScreen = !s.ChangeToFullScreen
				s.RightPanel.SetFullScreen(s.ChangeToFullScreen)
				return nil
			}
		}
		return event
	})
}
