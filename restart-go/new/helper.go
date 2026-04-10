package main

import (
	"fmt"
	"strconv"
	"strings"
)

func proces(s string) string {
	word := strings.Fields(s)
	result := make([]string, len(word))
	copy(result, word)

	for i := 0; i < len(result); i++ {
		switch result[i] {
		case "(up)":
			if i > 0 {
				result[i-1] = strings.ToUpper(result[i-1])
			}
			result = append(result[:i], result[i+1:]...)
			i--
		case "(low)":
			if i > 0 {
				result[i-1] = strings.ToLower(result[i-1])
			}
			result = append(result[:i], result[i+1:]...)
			i--
		case "(cap)":
			if i > 0 {
				word := result[i-1]
				result[i-1] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
			result = append(result[:i], result[i+1:]...)
			i--
		case "(bin)":
			if i > 0 {
				val, err := strconv.ParseInt(result[i-1], 2, 64)
				if err == nil {
					result[i-1] = fmt.Sprintf("%d", val)
				}
				result = append(result[:i], result[i+1:]...)
				i--

			}
		case "(hex)":
			if i > 0 {
				val, err := strconv.ParseInt(result[i-1], 16, 64)
				if err == nil {
					result[i-1] = fmt.Sprintf("%d", val)
				}
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(result, " ")

}
