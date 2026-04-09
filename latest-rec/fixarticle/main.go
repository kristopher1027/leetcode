package main

import (
	"fmt"
	"strings"
)

func fixArticle(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiou", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(fixArticle("i have a apple, a orange, a book "))
}
