package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Help() *tview.Frame {

	options := map[string]string{
		"Body":                        "Alt-b",
		"Query Param":                 "Alt-q",
		"Path Param":                  "Alt-p",
		"Header":                      "Alt-h",
		"Response":                    "Alt-r",
		"Enviar petición":             "Alt-d",
		"Cambiar a pantalla completa": "Alt-f",
		"Copiar respuesta":            "Alt-c",
		"salir":                       "alt-q",
	}

	editorBodyQpHeader := "Cada editor ([red]Body[#FFFFFF], [red]Query Param[#FFFFFF] y [red]Header[#FFFFFF]) tiene que tener este formato:\nx-api-key: [red]123[#FFFFFF],\nx-test: [red]123[#FFFFFF],\n#x-text: [red]123 [gray]<- Esto será ignorado[#FFFFFF]"
	editorPathParam := "Los Path Param tiene este formato:\n[red]users[#FFFFFF], [white]1[#FFFFFF].\n\nNO se puede ignorar, esto es así porque los Path Params son esto [red]example.com/users/1[#FFFFFF]"

	dialog := "Cambiar entre ventanas\n\n"

	for k, v := range options {
		whiteSpace := ""
		maxLength := 40 - len(k)

		for i := 0; i < maxLength; i++ {
			whiteSpace += " "
		}

		dialog += k + whiteSpace + "[red]" + strings.ToUpper(v) + "[#FFFFFF]" + "\n"
	}

	dialog += "\n" + editorBodyQpHeader + "\n\n" + editorPathParam

	text := tview.NewTextView().SetDynamicColors(true)
	text.SetBackgroundColor(tcell.ColorBlack.TrueColor())
	text.SetTextColor(tcell.ColorWhite.TrueColor())
	text.SetBorder(true).SetTitle(" > Help ").SetTitleAlign(tview.AlignLeft)
	text.SetBorderPadding(1, 1, 1, 1)

	fmt.Fprintln(text, dialog)

	frame := tview.NewFrame(text).SetBorders(1, 1, 1, 1, 1, 1).AddText("> Para volver presione (F1)", false, tview.AlignLeft, tcell.ColorWhite.TrueColor())

	frame.SetBackgroundColor(tcell.ColorBlack.TrueColor())

	return frame

}
