package main

import (
	"strings"
)

func fixpunc(word string) string {
	s := strings.Fields(word)
	result := ""
	for i, t := range s {
		if i > 0 && !strings.ContainsAny(t, "?.,/!.}{()}") {
			result += " "
		}
		result += t
	}
	return result + " "
}
