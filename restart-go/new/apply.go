package main

import (
	"strconv"
	"strings"
)

func applycase(s string) string {
	word := strings.Fields(s)

	var result []string

	for i := 0; i < len(word); i++ {
		if strings.HasPrefix(word[i], "(") {
			val := strings.TrimSuffix(word[i], ",")
			command := strings.Trim(val, "()")

			n := 1
			if i+1 < len(word) {
				val := strings.Trim(word[i+1], ")")
				if re, err := strconv.Atoi(val); err == nil {
					n = re
					i++

				}
			}
			if n > len(result) {
				n = len(result)
			}
			for j := len(result) - n; j < len(result); j++ {
				switch command {
				case "low":
					result[j] = strings.ToLower(result[j])
				case "up":
					result[j] = strings.ToUpper(result[j])
				case "cap":
					if len(result[j]) > 0 {
						result[j] = strings.ToUpper(result[j][:1]) + strings.ToUpper(result[j][1:])
					}
				}
			}
		} else {
			result = append(result, word[i])
		}

	}
	return strings.Join(result, " ")
}
