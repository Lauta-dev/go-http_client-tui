package utils

import (
	"http_client/const/mime"

	"github.com/yosssi/gohtml"
)

type Ident struct {
	ToDisplay string
	Lexer     string
}

func IdentText(response []byte, contentType string) Ident {
	var toDisplay string = string(response)
	lexer := MimeToLexer(contentType)

	switch lexer {

	case mime.Html:
		toDisplay = string(gohtml.FormatBytes(response))

	case mime.Json:
		toDisplay = string(IndentJson(response))

	default:
		toDisplay = string(response)
	}

	return Ident{
		ToDisplay: toDisplay,
		Lexer:     lexer,
	}
}
