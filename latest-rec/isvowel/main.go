package main

import (
	"fmt"
)

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}

// 	switch c {
// 	case 'a', 'e', 'i', 'o', 'u',
// 		'A', 'E', 'I', 'O', 'U':
// 		return true
// 	}
// 	return false
// }

func main() {
	fmt.Println(isVowel('a')) // true
	fmt.Println(isVowel('b')) // false
	fmt.Println(isVowel('U')) // true
}
