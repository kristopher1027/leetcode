// Question 2 - Hex & Binary Conversion
// Task:
// Write a function:
// func ConvertNumbers(words []string) []string
// Behavior:
// (hex) → convert previous word from hex → decimal
// (bin) → convert previous word from binary → decimal
// Example:
// ["1E", "(hex)", "files"] → ["30", "files"]
// ["10", "(bin)", "years"] → ["2", "years"]
// Constraints:
// Use strconv.ParseInt
// Handle invalid numbers safely (ignore or skip)

package main

import (
	"fmt"
	"strconv"
)

func convdec(words []string) []string {
	result := make([]string, len(words))
	copy(result, words)

	for i := 0; i < len(result); i++ {
		switch result[i] {
		case "(bin)":
			if i > 0 {
				if num, err := strconv.ParseInt(result[i-1], 2, 64); err == nil {
					result[i-1] = fmt.Sprintf("%d", num)
				}
			}
			result = append(result[:i], result[i+1:]...)
			i--
		case "(hex)":
			if i > 0 {
				if num, err := strconv.ParseInt(result[i-1], 16, 64); err == nil {
					result[i-1] = fmt.Sprintf("%d", num)
				}
			}
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}
	return result
}

func main() {
	res := []string{"i", "have", "1010", "(bin)", "books"}
	res1 := []string{"i", "have", "1E", "(hex)", "books"}
	fmt.Println(convdec(res))
	fmt.Println(convdec(res1))
}
