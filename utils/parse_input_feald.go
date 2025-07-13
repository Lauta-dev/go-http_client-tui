package utils

import (
	"fmt"
	"strings"
)

func ParseInput(input string, variables map[string]string) (string, error) {
	/*re := regexp.MustCompile(`@[^/\s]+(?=[/ ]|$)`)
	matches := re.FindAllString(input, -1)*/

	s := strings.Split(input, "/")
	noMatch := ""

	for i, v := range s {
		if strings.HasPrefix(v, "@") {
			key := v[1:] // Quitar la @
			val, ok := variables[key]
			if ok {
				s[i] = val
			} else {
				noMatch += v + "\n" // Solo agrega el no encontrado
			}
		}
	}

	if noMatch != "" {
		return "", fmt.Errorf("Algunos elemenos no se encuentar en la pestaña de variables: \n%s", noMatch)
	}

	return strings.Join(s, "/"), nil
}
