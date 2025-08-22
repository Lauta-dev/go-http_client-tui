package utils

// El cuerpo es el siguiente:
// "x-api-key: 123; Auth: bearer 123123123"
func ParseQueryParams(str string) map[string]string {
	h := ParseKeyValueText(str)

	return h
}
