package utils

import "strings"

// El cuerpo es el siguiente:
// "x-api-key: 123; Auth: bearer 123123123"
func ParseHeader(headerStr string) map[string]string {
	headersMap := make(map[string]string)
	ignoreChar := "#"

	if headerStr == "" {
		return headersMap
	}

	headers := strings.Split(headerStr, ",")

	for _, h := range headers {
		parts := strings.SplitN(strings.TrimSpace(h), ":", 2)

		if len(parts) == 2 {

			if strings.Contains(parts[0], ignoreChar) {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headersMap[key] = value
		}

	}

	return headersMap
}
