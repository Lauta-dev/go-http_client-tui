package utils

import (
	mime "http_client/const/mime"
	"strings"
)

// Mapear MIME types a lexers de Chroma
func MimeToLexer(contentType string) string {
	if strings.Contains(contentType, mime.JsonLong) {
		return mime.Json
	}
	if strings.Contains(contentType, mime.HtmlLong) {
		return mime.Html
	}
	// Podés agregar más tipos
	return mime.PlainText
}
