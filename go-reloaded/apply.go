package main

import (
	//"fmt"

	"strconv"
	"strings"
)

func commandN(s string) string {
	var result []string
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		if strings.HasPrefix(word[i], "(") {
			val := strings.TrimSuffix(word[i], ",")
			command := strings.Trim(val, "()")

			n := 1
			if i+1 < len(word) {
				res := strings.Trim(word[i+1], `")"`)
				num, _ := strconv.Atoi(res)
				n = num
				i++
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
					result[j] = strings.ToUpper(result[j][:1]) + strings.ToLower(result[j][1:])
				}
			}

		} else {
			result = append(result, word[i])
		}
	}
	return strings.Join(result, " ")
}

// func main() {
// 	fmt.Println(commandN("HELLO WORD (low, 2)  files (up, 1)  were added (cap, 2)"))
// 	fmt.Println(commandN("This is so exciting (up, 2)"))
// }
