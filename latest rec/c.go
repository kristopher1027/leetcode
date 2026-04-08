package main

import (
	"fmt"
	"strings"
)

func fixPunctuation(input []string) string {
	//words := strings.Fields(input)
	res := ""
	for i, t := range input {
		if i > 0 && !strings.ContainsAny(t, "?.!,:;") {
			res += " "
		}
		res += t
	}
	return res
}

// 	for _, w := range words {
// 		if strings.ContainsAny(w, ",.!?;:") && len(result) > 0 {
// 			result[len(result)-1] += w
// 		} else {
// 			result = append(result, w)
// 		}
// 	}
// 	return strings.Join(result, " ")
// }

func main() {
	text2 := []string{"how", ",", "are", "you", "!", "doing"}
	//text := "how , are you ! doing  ."
	//fmt.Println(fixPunctuation(text)) // Output: how, are you! doing.
	fmt.Println(fixPunctuation(text2))
}
