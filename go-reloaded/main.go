package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	val, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error in input.txt")
		return
	}
	input := string(val)
	result := processor(input)
	if err != os.WriteFile("output.txt", []byte(result), 0644) {

		fmt.Println("error in output.txt")
		return
	}
	fmt.Println("processed successful")
}
