package utils

import "fmt"

func ResponseInfoFormat(contentType string, url string, code string) string {
	return fmt.Sprintf("%s, %s \nURL: %s", StatusCodesColors(code), contentType, url)
}
