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
			// "(cap,","2)"

			n := 1
			if i+1 < len(word) {
				yea := strings.Trim(word[i+1], ")")
				if val, err := strconv.Atoi(yea); err == nil {
					n = val
					i++
				}
			}
			if n > len(result) {
				n = len(result)
			}
			for j := len(result) - n; j < len(result); j++ {
				switch command {
				case "up":
					result[j] = strings.ToUpper(result[j])
				case "low":
					result[j] = strings.ToLower(result[j])
				case "cap":
					if len(result[j]) > 0 {

						result[j] = strings.ToUpper(result[j][:1]) + strings.ToLower(result[j][1:])
					}
				}

			}

		} else {
			result = append(result, word[i])
		}

	}
	return strings.Join(result, " ")
}

// func main() {
// 	fmt.Println(applycase([]string{"this", "is", "so", "fun", "(up,", "2)"}))
// 	fmt.Println(applycase([]string{"hello", "world", "(up)"}))
// }
