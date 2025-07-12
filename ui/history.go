package ui

/*
CREATE TABLE request_history (
  id TEXT PRIMARY KEY,
  url TEXT,
  method TEXT,
  status_code INTEGER,
  content_type TEXT,
  response_body TEXT,
  created_at TEXT DEFAULT CURRENT_TIMESTAMP
);
DROP TABLE IF EXISTS request_history;
*/

import (
	"fmt"
	constants "http_client/const"
	"http_client/logic"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func History() *tview.Flex {
	list := tview.NewList()
	list.SetSecondaryTextColor(constants.ColorBackground.TrueColor())
	list.SetBorder(true)
	list.SetTitle(" > Historial de Request ('F2' para volver, 'F3' para actualizar lista)")

	flex := tview.NewFlex()
	responseView := ResponseView()

	for _, v := range logic.Ggg() {
		code := strconv.Itoa(v.StatusCode)

		mainText := fmt.Sprintf("(%s, %s) - %s", v.Method, code, v.URL)
		list.AddItem(mainText, v.ID, 'a', nil)
	}

	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		responseView.Clear()

		row, err := logic.GetItemById(secondaryText)

		if err != nil {
			fmt.Fprintf(responseView, "[red]%s", err.Error())
			return
		}

		fmt.Fprintf(responseView, "%s", row.ResponseBody)

	})

	flex.AddItem(list, 0, 1, true)
	flex.AddItem(responseView, 0, 1, false)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyF3:
			list.Clear()
			for _, v := range logic.Ggg() {
				code := strconv.Itoa(v.StatusCode)

				mainText := fmt.Sprintf("(%s, %s) - %s", v.Method, code, v.URL)
				list.AddItem(mainText, v.ID, 'a', nil)
			}
		}

		return event
	})

	return flex
}
