package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 26; i++ {
		char := string('A' + i)

		for j := 0; j <= i; j++ {
			fmt.Print(char)
		}
		fmt.Println(" ")
	}
}
