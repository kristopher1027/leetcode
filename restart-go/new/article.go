package main

import (
	"strings"
)

func article(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiouhAEIOUH", (string(words[i+1][0]))) {
			words[i] = "an"
		}
	}
	return strings.Join(words, " ")
}

// func main() {
// 	fmt.Println(article("i have a apple, a orange, a book "))
// }
