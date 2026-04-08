// Question 3

// Task:

// Support commands like:

// Code
// (up, 2)
// (low, 3)
// (cap, 4)
// Example:
// input:
// ["this", "is", "so", "fun", "(up,", "2)"]
// Output:
// ["this", "is", "SO", "FUN"]
// Requirements:

// Detect (command, n)

// Apply to last n words

// Handle edge cases:

// n > number of words

// invalid numbers

// Bonus:
// Combine with single commands:

// Code
// (up) vs (up 2)

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func process(input []string) []string {
	var result []string
	for i := 0; i < len(input); i++ {

		if strings.HasPrefix(input[i], "(") {
			val := strings.TrimSuffix(input[i], ",")
			command := strings.Trim(val, "( )")
			n := 1
			if i+1 < len(input) {
				yea := strings.Trim(input[i+1], ")")
				if val, err := strconv.Atoi(yea); err == nil {
					n = val
					i++
				}
			}
			if n > len(result) {
				n = len(result)
			}
			start := len(result) - n
			for j := start; j < len(result); j++ {
				switch command {
				case "up":
					result[j] = strings.ToUpper(result[j])
				case "low":
					result[j] = strings.ToLower(result[j])
				case "cap":
					if len(result[j]) > 0 {
						result[j] = strings.ToUpper(result[j][:1]) + strings.ToLower(result[j][1:])
					}
				}
			}
		} else {
			result = append(result, input[i])
		}
	}
	return result
}

func main() {
	fmt.Println(process([]string{"this", "is", "so", "fun", "(up,", "2)"}))
	fmt.Println(process([]string{"hello", "world", "(up)"}))
}
