package main

import (
	//"fmt"
	"regexp"
	"strings"
)

func fixALL(s string) string {
	word := strings.Join(strings.Fields(s), " ")
	// FIX PUNCTUATION

	res := regexp.MustCompile(`\s*([.,?;!:/])`)
	word = res.ReplaceAllString(word, "$1")

	res1 := regexp.MustCompile(`([.,?;!:/])([^\s.,?;!:/])`)
	word = res1.ReplaceAllString(word, "$1 $2")

	// FIX QUOTE
	res2 := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	word = res2.ReplaceAllString(word, `'$1'`)

	res3 := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	word = res3.ReplaceAllString(word, `"$1"`)
	// FIX ARTICLE
	word = regexp.MustCompile(`a\s+([aeiouhAEIOUH])`).ReplaceAllString(word, `an $1`)
	word = regexp.MustCompile(`A\s+([aeiouhAEIOUH])`).ReplaceAllString(word, `An $1`)
	word = regexp.MustCompile(`an\s+([^aeiouhAEIOUH])`).ReplaceAllString(word, `a $1`)
	word = regexp.MustCompile(`An\s+([^aeiouhAEIOUH])`).ReplaceAllString(word, `A $1`)
	return word

}

// func main() {
// 	fmt.Println(fixALL(` "      There it was    .A amazing rock      !     "`))
// }
