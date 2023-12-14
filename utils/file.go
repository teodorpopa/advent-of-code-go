package utils

import (
	"os"
	"strings"
)

func ReadFile(path string) string {
	buf, _ := os.ReadFile(path)
	return string(buf)
}

func ReadLines(input string) []string {
	return strings.Split(input, "\n")
}
