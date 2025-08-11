package utils

import (
	chromacolor "http_client/const/chroma_color"

	"github.com/alecthomas/chroma/quick"
	"github.com/rivo/tview"
)

func PrettyStyle(contentType string, response []byte, responseView *tview.TextView) (string, error) {
	data := IdentText(response, contentType)
	// Preparar para escribir el output
	writer := tview.ANSIWriter(responseView)

	err := quick.Highlight(writer, data.ToDisplay, data.Lexer, "terminal", chromacolor.Highlight)
	return data.ToDisplay, err
}
