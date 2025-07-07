package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func helpFile() string {
	return `================================================================================
                           [blue::b]GO HTTP CLIENT TUI - AYUDA[-::-]
================================================================================

> [yellow::b]CONTROLES PRINCIPALES[-::-]
================================================================================
[green::b]F1[-::-]          - Mostrar/Ocultar ayuda
[green::b]Alt+F[-::-]       - Pantalla completa para respuesta
[green::b]Alt+B[-::-]       - Vista Body
[green::b]Alt+Q[-::-]       - Vista Query Params
[green::b]Alt+P[-::-]       - Vista Path Params
[green::b]Alt+H[-::-]       - Vista Headers
[green::b]Alt+D[-::-]       - Enviar petición
[green::b]Alt+J[-::-]       - Vista variables
[green::b]Alt+N[-::-]       - Copiar petición
[red::b]Ctrl+C[-::-]      - Salir

> [yellow::b]FORMATO PARA BODY, HEADERS Y QUERY PARAMS[-::-]
================================================================================
• [white::b]Formato:[-::-] clave: valor
• [white::b]Comentarios:[-::-] # texto (ignorado)
[green::b]Ejemplos:[-::-]
Content-Type: application/json
Authorization: Bearer token123
x-api-key: abc123
[gray]# Este es un comentario[-]

> [cyan::b]PATH PARAMETERS[-::-]
================================================================================
• Una línea por parámetro
• Se añaden a la URL en orden
[green::b]Ejemplo:[-::-]
[white::b]URL:[-::-] https://api.com
[white::b]Parámetros:[-::-]
  users,
  123
[white::b]Resultado:[-::-] https://api.com/users/123

> [magenta::b]VARIABLES[-::-]
================================================================================
• [white::b]Definir:[-::-] clave: valor
• [white::b]Usar:[-::-] @variable en URL
[green::b]Ejemplo:[-::-]
[white::b]Variables:[-::-]
  HOST: https://api.com,
[white::b]Uso:[-::-]
  URL: @HOST/users

> [orange::b]MÉTODOS HTTP[-::-]
================================================================================
[green::b]GET[-::-]     - Obtener recursos
[blue::b]POST[-::-]    - Crear recursos
[yellow::b]PUT[-::-]     - Actualizar recursos
[red::b]DELETE[-::-]  - Eliminar recursos

> [purple::b]CÓDIGOS DE ESTADO[-::-]
================================================================================
[white::b]1xx[-::-] - Información (Blanco)
[green::b]2xx[-::-] - Éxito (Verde)
[yellow::b]3xx[-::-] - Redirección (Amarillo)
[red::b]4xx/5xx[-::-] - Error (Rojo)

> [cyan::b]EJEMPLO RÁPIDO[-::-]
================================================================================
[blue::b]POST[-::-] a https://jsonplaceholder.typicode.com/posts
[white::b]Headers:[-::-]
  Content-Type: application/json
[white::b]Body:[-::-]
  [cyan]{[-]
    [yellow]"title"[-]: [green]"test"[-],
    [yellow]"body"[-]: [green]"content"[-],
    [yellow]"userId"[-]: [blue]1[-]
  [cyan]}[-]
[green::b]Presiona Alt+D para enviar[-::-]
	`
}

func Help() *tview.Frame {
	dialog := helpFile()

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
