package utils

import (
	"fmt"
	"http_client/const/prefix"
	"net/url"
	"os"
	"strings"
)

// ReplaceVariablesInURL reemplaza @variable en URLs con valores proporcionados o variables de entorno.
func ReplaceVariablesInURL(URL string, variables map[string]string) (string, error) {
	parts := strings.Split(URL, "/")
	notFound := ""

	for index, part := range parts {
		if strings.HasPrefix(part, prefix.VariablePrefix) {
			key := part[1:] // Quitar la @

			val, ok := variables[key]

			if ok {
				parts[index] = val
				continue
			}

			if env := os.Getenv(key); env != "" {
				parts[index] = env
				continue
			}

			notFound += part + "\n" // Solo agrega el no encontrado
		}
	}

	if notFound != "" {
		return "", fmt.Errorf("Algunos elemenos no se encuentar en la pestaña de variables: \n%s", notFound)
	}

	return strings.Join(parts, "/"), nil
}

// AddPathParam Agrega los path params para que sea example.com/user/1
func AddPathParam(params []string, URL string) string {
	pathParam := ""
	for _, v := range params {
		pathParam += "/" + v
	}

	return URL + pathParam
}

// AddQueryParam parsea los datos para que sean example.com?user=1
func AddQueryParam(URL *url.URL, queryParams map[string]string) string {
	q := URL.Query()

	for k, v := range queryParams {
		if q.Get(k) == "" {
			q.Add(k, v)
		} else {
			q.Set(k, v)
		}
	}

	URL.RawQuery = q.Encode()
	return URL.String()
}

// AddVariableToURL Añade las variables faltantes a la URL @url -> http://localhost:4000
func AddVariableToURL(URL, variables string) (string, error) {
	url, err := ReplaceVariablesInURL(
		URL,
		ParseKeyValueText(variables),
	)

	if err != nil {
		return "", err
	}

	return url, nil
}

// LLama a AddPathParam, AddQueryParam y AddVariableToURL para retornar una URL completa
func ParseUrl(variables, URL string, queryParams map[string]string, pathParam []string) (string, error) {
	c, err := AddVariableToURL(URL, variables)

	if err != nil {
		return "", err
	}

	c = AddPathParam(pathParam, c)
	u, err := url.Parse(c)

	if err != nil {
		return "", err
	}

	c = AddQueryParam(u, queryParams)

	return c, nil
}
