// input hello my people output: hellO mY peoplE
package main

import (
	"fmt"
	"strings"
)

func art(s string) string {
	res := strings.Fields(s)
	for i, v := range res {
		res[i] = strings.ToLower(string(v[:len(v)-1])) + strings.ToUpper(string(v[len(v)-1]))

	}
	return strings.Join(res, " ")

}

func num(a []int) int {
	result := 0
	for _, v := range a {
		result ^= v
		//fmt.Println(v)
	}
	return result
}
func firstrune(s string) rune {
	r := []rune(s)
	return (r[0])
}
func lastrune(s string) rune {
	r := []rune(s)
	return (r[len(r)-1])
}
func word(s string) rune {
	r := []rune(s)
	return (r[len(r)-1])
}
func wordnum(s string) int {
	return len(strings.Fields(s))
}
func wordal(s []string) int {
	return len(s)
}

func trimspace(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "  ", " ")
	//s = strings.Trim(s, "")
	return s
}
func trimspa(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func main() {
	fmt.Println(art("hello my people"))
	fmt.Println(num([]int{1, 2, 3, 4, 1, 2, 3, 4, 5}))
	fmt.Printf("%c\n", firstrune("🌍go"))
	fmt.Printf("%c\n", lastrune("gogdhfjghk🌍"))
	re := []string{"hello", "world", "hi"}
	fmt.Println(wordnum("hello how are you"))
	fmt.Println(wordal(re))
	fmt.Println(trimspace("         hello  world      "))
	fmt.Println(trimspa("         hello                                         world            "))
}

// "hello","world","hi" => 3
