package main

import (
	"fmt"
	"strings"
)

func isGroupPunctuation(s string) bool {
	if len(s) <= 1 {
		return false
	}
	return strings.ContainsAny(s, ",.!?;:") && len(s) > 1
}

func main() {
	fmt.Println(isGroupPunctuation("...")) // true
	fmt.Println(isGroupPunctuation("!?"))  // true
	fmt.Println(isGroupPunctuation("."))   // false
}
