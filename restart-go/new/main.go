package main

import (
	"fmt"
	"os"
	"strings"
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
	result := strings.Fields(string(val))
	result = applycase(result)
	result = proces(result)
	result1 := strings.Join(result, " ")

	if err != os.WriteFile(os.Args[2], []byte(result1), 0664) {
		fmt.Println("error", err)
	}

	fmt.Println(string(val))
}
