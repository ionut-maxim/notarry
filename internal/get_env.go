package internal

import (
	"os"
	"strings"
)

func GetEnvVariables(prefix string) []string {
	var result []string

	for _, pair := range os.Environ() {
		if strings.Contains(pair, "radarr") {
			result = append(result, pair)
		}
	}

	return result
}
