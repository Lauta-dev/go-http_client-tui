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

func ReadEnvFile(path string) string {
	if path == "" {
		return ""
	}

	ignoreChar := "#"
	mapp := []Env{}
	final_string := ""

	fi, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	// string
	arch := string(fi)

	s := strings.Split(arch, "\n")

	for _, v := range s {
		if strings.HasPrefix(v, ignoreChar) {
			continue
		}

		if v != "" {
			d := strings.Split(v, "=")
			key := strings.TrimSpace(d[0])
			value := strings.TrimSpace(d[1])

			mapp = append(mapp, Env{
				key, value,
			})
		}
	}

	for _, v := range mapp {
		l := fmt.Sprintf("%s: %s,\n", v.key, v.value)

		final_string += l
	}

	fmt.Println(final_string)

	return final_string
}
