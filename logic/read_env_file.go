package logic

import (
	"fmt"
	"os"
	"strings"
)

type Env struct {
	key   string
	value string
}

// Si hay algún error la función retorna una variable vacía
func ReadEnvFile(path string) string {
	if path == "" {
		return ""
	}

	ignorePrefix := "#"
	envVars := []Env{}

	content, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, ignorePrefix) || line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "" {
			continue
		}

		envVars = append(envVars, Env{
			key, value,
		})
	}

	// Crear un string secuecial
	var builder strings.Builder
	for _, v := range envVars {
		builder.WriteString(fmt.Sprintf("%s: %s,\n", v.key, v.value))
	}

	return builder.String()
}
