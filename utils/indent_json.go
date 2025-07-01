package utils

import "encoding/json"

func IndentJson(bytes []byte) string {
	var post any
	isJson := true

	switch bytes[0] {
	case '{': // Por si es un Object
		post = make(map[string]any, 0)
	case '[': // Por si es un Array
		post = make([]map[string]any, 0)
	default:
		isJson = false
	}

	if !isJson {
		return string(bytes)
	}

	err := json.Unmarshal(bytes, &post)

	if err != nil {
		return err.Error()
	}

	ident, err := json.MarshalIndent(post, "", " ")

	if err != nil {
		return err.Error()
	}

	return string(ident)
}
