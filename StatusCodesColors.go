package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StatusCodesColors(statusCode string) string {
	defaultColor := "[#FFFFFF]"
	//1xx
	//a := "[white]"
	//2xx
	//b := "[green]"
	// 3xx
	//c := "[yellow]"
	// 4xx/5xx
	//d := "[red]"

	status := strings.Split(statusCode, " ")
	code, _ := strconv.Atoi(status[0])

	var color string
	switch {
	case code >= 100 && code < 200:
		color = defaultColor
	case code >= 200 && code < 300:
		color = "[green]"
	case code >= 300 && code < 400:
		color = "[yellow]"
	case code >= 400:
		color = "[red]"
	}

	return fmt.Sprintf("%s%s%s", color, statusCode, defaultColor)

}
