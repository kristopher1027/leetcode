package main

import (
	"fmt"
	"strconv"
	"strings"
)

// question 1
func hexTodecimal(hexstr string) (int64, error) {
	return strconv.ParseInt(hexstr, 16, 64)
}

// question 2
func binTodecimal(hexstr string) (int64, error) {
	return strconv.ParseInt(hexstr, 2, 64)
}

// question 3
func capitaliseFirstWORD(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// question 4
func capitaliseLastprevWord(s []string, n int) []string {
	for i := len(s) - n; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}

// question 5
func isPunc(s string) bool {
	if s == "," || s == "!" {
		return true
	}
	return false
}

// question 6
func removespacefrompunc(s []string) string {
	res := ""
	for i, t := range s {
		if i > 0 && !strings.ContainsAny(t, "?.!,:;") {
			res += " "
		}
		res += t
	}
	return res
}

// question 7
func article(word string) string {
	if strings.Contains("aeiouh", strings.ToLower(string(word[0]))) {
		return "an"
	}
	return "a"
}

// question 8
func fixArticle(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiou", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}

	return strings.Join(words, " ")
}

// question 9
func fixsingleQuotes(text string) string {
	text = strings.Trim(text, "'")
	text = strings.TrimSpace(text)
	return "'" + text + "'"
}

func main() {
	// question 1
	fmt.Println(hexTodecimal("1E"))
	fmt.Println(hexTodecimal("FF"))

	// question 2
	fmt.Println(binTodecimal("10"))
	fmt.Println(binTodecimal("1010"))
	fmt.Println(binTodecimal("1111111"))

	// question 3
	fmt.Println(capitaliseFirstWORD("hELLO"))
	fmt.Println(capitaliseFirstWORD("WORLD"))

	// question 4
	fmt.Println(capitaliseLastprevWord([]string{"how", "are", "you", "doing"}, 3))

	// question 5
	fmt.Println(isPunc(","))
	fmt.Println(isPunc("!"))
	fmt.Println(isPunc("x"))

	// question 6
	fmt.Println(removespacefrompunc([]string{"how", ",", "you", "doing", "!"}))

	// question 7
	fmt.Println(article("Apple"))
	fmt.Println(article("bpple"))

	// question 8
	fmt.Println(fixArticle("this is a apple which is a orange"))

	// question 9
	input := "'  hello world      '"
	input2 := "'   awesome    '"
	result := fixsingleQuotes(input)
	result2 := fixsingleQuotes(input2)
	fmt.Printf("%q\n", result)
	fmt.Printf("%q\n", result2)

}

