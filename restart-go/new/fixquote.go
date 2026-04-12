package main

import (
	"strings"
)

func fixquote(s string) string {
	word := strings.Join(strings.Fields(s), " ")

	word = strings.TrimSpace(word)
	word = strings.Trim(word, "'")
	word = strings.TrimSpace(word)
	return "'" + word + "'"

}
