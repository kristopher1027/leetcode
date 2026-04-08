// package main

// import "fmt"

// func isCommand(s string) bool {
// 	if s == "(up)" || s == "bin" || s == "(cap)" || s == "(cap)" {
// 		return true
// 	}
// 	return false
// }
// func main() {
// 	fmt.Println(isCommand("(up)"))   // true
// 	fmt.Println(isCommand("(cap)"))  // true
// 	fmt.Println(isCommand("hello"))  // false
// 	fmt.Println(isCommand("(down)")) // false
// }

package main

import (
	"fmt"
	"strings"
)

func isCommand(s string) bool {
	if len(s) < 4 || !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
		return false
	}
	cmd := strings.Trim(s, "()")
	switch cmd {
	case "up", "low", "cap":
		return true
	}
	return false
}

func main() {
	fmt.Println(isCommand("    (up)   ")) // true
	fmt.Println(isCommand("(cap)"))       // true
	fmt.Println(isCommand("hello"))       // false
	fmt.Println(isCommand("(down)"))      // false
}
