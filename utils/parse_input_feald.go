package utils

import (
	"strings"
)

func ParseInput(input string, variables map[string]string) string {
	/*re := regexp.MustCompile(`@[^/\s]+(?=[/ ]|$)`)
	matches := re.FindAllString(input, -1)*/

	s := strings.Split(input, "/")

	for i, v := range s {

		for k, variable := range variables {
			if v == "@"+k {
				s[i] = variable
			}
		}

	}

	return strings.Join(s, "/")
}
