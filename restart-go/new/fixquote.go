package main

import (
	"strings"
)

func fixquote(s string) string {
	s = strings.ReplaceAll(s, " '", "'")
	s = strings.ReplaceAll(s, "' ", "'")
	s = strings.ReplaceAll(s, ":'", ": '")
	return s

}
