package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <sample.txt> <result.txt>")
		return
	}
	val, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error input:", err)
	}
	input := string(val)
	result := processor(input)

	if err != os.WriteFile(os.Args[2], []byte(result), 0664) {
		fmt.Println("error", err)
	}

	fmt.Println("processed successful:")
}
func processor(input string) string {
	input = proces(input)
	input = applycase(input)
	input = fixpunc(input)
	input = article(input)
	input = fixquote(input)
	input = multipunc(input)

	return input + "\n"
}
