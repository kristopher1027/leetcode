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
func convertToDecimal(str string, base int) (int64, error) {
	return strconv.ParseInt(str, base, 64)
}

// QUESTION 4
func contain(s string) bool {
	if s == "bin" || s == "hex" || s == "low" || s == "cap" || s == "up" {
		return true
	}
	return false
}

// QUESTION 5
func IsPunctuation(s string) string {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "0123456789"
	specials := "/,.?::;!"
	if len(s) != 1 {
		return "invalid"
	}
	if strings.ContainsAny(s, alpha) {
		return "letters"
	}
	if strings.ContainsAny(s, num) {
		return "numbers"
	}
	if strings.ContainsAny(s, specials) {
		return "specialCharacter"
	}
	return s
}

// question 6
func capitaliseFirstWORD(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// func capitaliseFirstWORD(s string) string {
// 	if len(s) == 0 {
// 		return s
// 	}
// 	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
// }

// question 7 update to capilize each word
// func capeveryword(s string) string {
// 	res := strings.ToLower(s)
// 	return strings.Title(res)
// }

func CapitaliseEachWord(s []string) string {
	// 	words := strings.Fields(strings.ToLower(s))

	for i := range s {
		s[i] = strings.ToUpper(s[i][:1]) + s[i][1:]
	}
	return strings.Join(s, " ")
}

// question 8
func capitaliseLastprevWord(s []string, n int) []string {
	for i := len(s) - n; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}

// question 9
func isPunc(s string) bool {
	if s == "," || s == "!" {
		return true
	}
	return false
}

// func isPunc(s string) bool {
// 	return strings.ContainsAny(s, ",!.?:;")
// }

// question 10
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

// question 11
func article(word string) string {
	if strings.Contains("aeiouh", strings.ToLower(string(word[0]))) {
		return "an"
	}
	return "a"
}

// question 12
func fixArticle(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && strings.Contains("aeiou", strings.ToLower(string(words[i+1][0]))) {
			words[i] = "an"
		}
	}

	return strings.Join(words, " ")
}

// alternative for no 13
func fixart(s string) string {
	s = strings.ReplaceAll(s, "A", "An")
	s = strings.ReplaceAll(s, "An book", " A book")
	return s
}

// question 14
func fixsingleQuotes(text string) string {
	text = strings.Trim(text, "'")
	text = strings.TrimSpace(text)
	return "'" + text + "'"
}

// question 15
func num(s string) int {
	return len(strings.Fields(s))
}

// question 16
func firstword(s string) string {
	return strings.Fields(s)[0]
}

// question 17
func lastword(s string) string {
	w := strings.Fields(s)
	return w[len(w)-1]

}

// question 18
func lastchar(s string) string {
	if s == "" {
		return ""
	}

	return string(s[len(s)-1])
}

// question 19
func firstchar(s string) string {
	return string(s[0])
}

// question 20
func inte(s string) (int, error) {
	return strconv.Atoi("10")
}

// question 21
func uplastword(s string) string {
	w := strings.Fields(s)
	i := len(w) - 1
	w[i] = strings.ToUpper(w[i])
	return strings.Join(w, " ")
}

// question 22
func extraspace(s string) string {
	res := strings.Fields(s)
	return strings.Join(res, " ")
}

// question 23
func reverseword(s string) string {
	w := strings.Fields(s)
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	return strings.Join(w, " ")

}

// question 24
func duplicatewords(s string) string {
	r := []string{}
	for _, v := range strings.Fields(s) {
		r = append(r, v, v)
	}
	return strings.Join(r, " ")
}

// question 25
func replaceword(s, old, new string) string {
	w := strings.Fields(s)
	for i := range w {
		if w[i] == old {
			w[i] = new
		}
	}
	return strings.Join(w, " ")
}

// question 26
func fixPunct(s string) string {
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " !", "!")
	return s
}

// question 27
func fixQuotes(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "'")
	return "'" + strings.TrimSpace(s) + "'"
}

// question 28
func removeWord(s, target string) string {
	w := strings.Fields(s)
	var r []string
	for _, v := range w {
		if v != target {
			r = append(r, v)
		}
	}
	return strings.Join(r, " ")
}

// question 29
func uplowcap(s, mode string) string {
	switch mode {
	case "low":
		return strings.ToLower(s)
	case "up":
		return strings.ToUpper(s)
	case "cap":
		w := strings.Fields(s)
		for i := range w {
			w[i] = strings.ToUpper(w[i][:1]) + strings.ToLower(w[i][1:])
		}
		return strings.Join(w, " ")
	default:
		return s
	}
}

// question 30
func removeCommand(words []string, index int) []string {
	if index < 0 || index >= len(words) {
		return words
	}
	return append(words[:index], words[index+1:]...)
}

// question 31
func applyUp(s string) string {
	w := strings.Fields(s)
	for i := 0; i < len(w); i++ {
		if w[i] == "(up)" && i > 0 {
			w[i-1] = strings.ToUpper(w[i-1])
			w = append(w[:i], w[i+1:]...)
			i--
		}
	}
	return strings.Join(w, " ")
}

func unmatched(a []int) int {
	result := 0
	for _, v := range a {
		result ^= v
		}
	return result
}

func main() {

	fmt.Println(num("go reloaded project"))
	fmt.Println(firstword("go reloaded project"))
	fmt.Println(lastword("go reloaded project"))
	fmt.Println(inte("10"))
	fmt.Println(firstchar("hello"))
	fmt.Println(lastchar(""))
	fmt.Println(lastchar("hello"))
	fmt.Println(uplastword("hello  word"))
	fmt.Println(uplastword("hello how are you"))
	//fmt.Println(capeveryword("how are YOU DoiNg"))
	fmt.Println(extraspace(" hello  world   dear "))
	fmt.Println(reverseword("hello world"))
	fmt.Println(duplicatewords("hello world"))
	fmt.Println(replaceword("hello word com", "word", "world"))
	fmt.Println(removeCommand([]string{"go", "(up)"}, 2))
	fmt.Println(removespacefrompunc([]string{"hello", ",", "world"}))

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
	input := "' hello world '"
	input2 := "' awesome '"

	result := fixsingleQuotes(input)
	result2 := fixsingleQuotes(input2)

	fmt.Printf("%q\n", result)
	fmt.Printf("%q\n", result2)

	fmt.Println(convertToDecimal("1010", 2))
	fmt.Println(convertToDecimal("77", 8))
	fmt.Println(convertToDecimal("1E", 16))

	fmt.Println(fixart("this is A apple which is A orange A book"))

	fmt.Println(IsPunctuation("!"))
	fmt.Println(IsPunctuation("7"))
	fmt.Println(IsPunctuation("G"))
	fmt.Println(IsPunctuation(""))

	fmt.Println(contain("bin"))
	fmt.Println(contain("hex"))
	fmt.Println(contain("xyz"))

	res := []int{1, 2, 3, 4, 5, 7, 9, 1, 2, 3, 4, 5, 7}
	fmt.Println(unmatchN(res))
}
