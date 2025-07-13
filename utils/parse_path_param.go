package utils

import "strings"

func ParsePathParams(param string) []string {
	params := []string{}

	ignoreChar := "#"
	headers := strings.Split(param, ",")

	for _, v := range headers {
		if v == "" || strings.Contains(v, ignoreChar) {
			continue
		}

		params = append(params, v)
	}

	return params
}
