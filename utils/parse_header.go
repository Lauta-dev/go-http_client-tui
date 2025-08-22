package utils

import (
	"http_client/const/prefix"
	"os"
	"strings"
)

func ParseHeaders(str string, variables map[string]string) map[string]string {
	headers := ParseKeyValueText(str)
	headersMap := make(map[string]string)

	for k, v := range headers {
		var key string = k
		var value string = v

		if strings.HasPrefix(v, prefix.VariablePrefix) {
			// Examinar las variables en la pesantañ de variables
			// Examinar os.GetEnv("")
			keyWithoutPrefix := v[1:]
			if variable := variables[keyWithoutPrefix]; variable != "" {
				value = variable
			}

			if env := os.Getenv(keyWithoutPrefix); env != "" {
				value = env
			}

		}

		// Si value sigue teniendo el prefix de @, pasa a la siguiente iteracion
		// Esto es porque el anterior if cambia el valor de "value"
		if strings.HasPrefix(value, prefix.VariablePrefix) {
			continue
		}

		headersMap[key] = value
	}

	return headersMap
}
