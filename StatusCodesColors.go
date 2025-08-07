package main

import (
	"fmt"
	colors "http_client/const/color_wrapper"
	"strconv"
	"strings"
)

func StatusCodesColors(statusCode string) string {
	status := strings.Split(statusCode, " ")
	code, _ := strconv.Atoi(status[0])

	var color string
	switch {

	case code >= 100 && code < 200:
		color = fmt.Sprintf("[%s::b]", colors.ColorTextPrimary.String())

	case code >= 200 && code < 300:
		color = fmt.Sprintf("[%s::b]", colors.ColorSuccess.String())

	case code >= 300 && code < 400:
		color = fmt.Sprintf("[%s::b]", colors.ColorWarning.String())

	case code >= 400:
		color = fmt.Sprintf("[%s::b]", colors.ColorError.String())

	}

	return fmt.Sprintf("%s%s", color, statusCode)

}
