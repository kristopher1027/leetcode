// Question 1 - Basic Command Parsing
// Task:
// Write a Go function:

// func ToUpperLastWord(words []string) []string

// Input:
// ["hello", "world", "(up)"]

// Output:
// ["hello", "WORLD"]

// Requirements:
// Detect (up)

// Apply it ONLY to the previous word

// Remove (up) from output

// Tests:

// ["go", "is", "fun", "(up)"] → ["go", "is", "FUN"]

// ["hello", "(up)"] → ["HELLO"]

package main

import (
	"fmt"
	"strings"
)

func ToUpperLastWord(words []string) []string {
	result := make([]string, len(words))
	copy(result, words)

	for i := 0; i < len(result); i++ {
		switch result[i] {
		case "(cap)":
			if i > 0 {
				word := result[i-1]
				result[i-1] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}
			result = append(result[:i], result[i+1:]...)
			i--
		case "(low)":
			if i > 0 {
				result[i-1] = strings.ToLower(result[i-1])
			}
			result = append(result[:i], result[i+1:]...)
			i--
		case "(up)":
			if i > 0 {
				result[i-1] = strings.ToUpper(result[i-1])
			}
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}
	return result
}

func main() {
	res := []string{"HELLO", "(cap)", "WORD", "(low)", "hello", "(up)"}
	fmt.Println(ToUpperLastWord(res))
}
