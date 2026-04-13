package main

import (
	"regexp"
	// "strings"
)

func fixpunc(word string) string {
	punc1 := regexp.MustCompile(`\s+([.,?/:;!])`)
	word = punc1.ReplaceAllString(word, "$1")

	//punc2 := regexp.MustCompile(`([.,?/:;!])(\s*)(\w)`)
	punc2 := regexp.MustCompile(`([.,?/:;!])([a-zA-Z0-9])`)
	word = punc2.ReplaceAllString(word, "$1 $2")

	return word

}

// func main() {
// 	fmt.Println(fixpunc("Punctuation tests are ... kinda boring ,what do you think ?"))
// }

// 	s := strings.Fields(word)
// 	result := ""
// 	for i, t := range s {
// 		if i > 0 && !strings.ContainsAny(t, "?.,/!.}{()}") {
// 			result += " "
// 		}
// 		result += t
// 	}
// 	return strings.Join(s, " ")
// }
// func multipunc(s string) string {
// 	s = strings.ReplaceAll(s, ". . .", "...")
// 	//s = strings.ReplaceAll(s, "  ... ", "...")
// 	s = strings.ReplaceAll(s, " ?!", "?!")
// 	return s
// }
