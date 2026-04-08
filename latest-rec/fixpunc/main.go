package main

import (
	"fmt"
	"strings"
)

func fixPunctuation(input []string) string {

	res := ""
	for i, t := range input {
		if i > 0 && !strings.ContainsAny(t, "?.!,:;") {
			res += " "
		}
		res += t
	}
	return res
}

func fixPunctuation2(input string) string {
	words := strings.Fields(input)
	res := ""
	for i, t := range words {
		if i > 0 && !strings.ContainsAny(t, "?.!,:;") {
			res += " "
		}
		res += t
	}
	return res
}

func main() {
	text := []string{"how", ",", "are", "you", "!", "doing"}
	fmt.Println(fixPunctuation(text)) // Output: how, are you! doing.

	text2 := "how , are you ! doing  ."
	fmt.Println(fixPunctuation2(text2))

}
