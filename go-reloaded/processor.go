package main

import (
	"strings"
)

func processor(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)
		line = fixALL(line)
		line = process(line)
		line = commandN(line)

		line = process(line)

		//line = commandN(line)

		lines[i] = line
	}
	return strings.Join(lines, "\n")
}
