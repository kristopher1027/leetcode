package main

import (
	//"fmt"
	"fmt"
	"strconv"
	"strings"
)

func process(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {

		if word[i] == "(hex)" && i > 0 {
			val, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				fmt.Println("invalid hex", err)
			}

			word[i-1] = strconv.FormatInt(val, 10)

			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(bin)" && i > 0 {
			val, err := strconv.ParseInt(word[i-1], 2, 64)
			if err != nil {
				fmt.Println("invalid bin", err)
			}

			word[i-1] = strconv.FormatInt(val, 10)

			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(up)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(low)" && i > 0 {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(cap)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1][:1]) + strings.ToLower(word[i-1][1:])
			word = append(word[:i], word[i+1:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}

// func main() {
// 	fmt.Println(process("HELLO WORD (low) 1E (hex) files (up) were 1010 (bin) added (cap)"))
// }
