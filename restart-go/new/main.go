package main

import (
	"fmt"
	"os"
)

func main() {
	val, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("error input:", err)
	}
	if err != os.WriteFile("result.txt", val, 0664) {
		fmt.Println("error", err)
	}
	fmt.Println(string(val))
}
