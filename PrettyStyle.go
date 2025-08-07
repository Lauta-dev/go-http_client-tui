package main

import (
	chromacolor "http_client/const/chroma_color"
	mime "http_client/const/mime"
	"http_client/utils"

	"github.com/alecthomas/chroma/quick"
	"github.com/rivo/tview"
	"github.com/yosssi/gohtml"
)

func PrettyStyle(contentType string, response []byte, responseView *tview.TextView) (string, error) {
	lexer := utils.MimeToLexer(contentType)

	// Preparar para escribir el output
	writer := tview.ANSIWriter(responseView)

	var toDisplay string = string(response)

	switch lexer {

	case mime.Html:
		toDisplay = string(gohtml.FormatBytes(response))

	case mime.Json:
		toDisplay = string(utils.IndentJson(response))

	default:
		toDisplay = string(response)
	}

	err := quick.Highlight(writer, toDisplay, lexer, "terminal", chromacolor.Highlight)
	return toDisplay, err
}
