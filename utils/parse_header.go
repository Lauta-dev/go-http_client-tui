package utils

import (
	"os"
	"strings"
)

func ParseHeadera(str string, variables map[string]string) map[string]string {
	headersMap := make(map[string]string)
	prefix := "#"
	varPrefix := "@"

	if str == "" {
		return headersMap
	}

	headers := strings.Split(str, "\n")

	for _, h := range headers {
		if strings.HasPrefix(h, prefix) {
			continue
		}

		parts := strings.SplitN(strings.TrimSpace(h), ": ", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if strings.HasPrefix(value, varPrefix) {
			// Examinar las variables en la pesantañ de variables
			// Examinar os.GetEnv("")
			removePrefix := value[1:]
			if variables[removePrefix] != "" {
				value = variables[removePrefix]
			}

			if os.Getenv(removePrefix) != "" {
				value = os.Getenv(removePrefix)
			}

		}

		// Si value sigue teniendo el prefix de @, pasa a la siguiente iteracion
		// Esto es porque el anterior if cambia el valor de "value"
		if strings.HasPrefix(value, varPrefix) {
			continue
		}

		headersMap[key] = value
	}

	return headersMap
}
