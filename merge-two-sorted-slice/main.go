package main

import "fmt"

func mergeTwoSlices(list1 []int, list2 []int) []int {
	i := 0
	j := 0
	merged := []int{}
	for i < len(list1) && j < len(list2) {
		if list1[i] <= list2[j] {
			merged = append(merged, list1[i])
			i++
		} else {
			merged = append(merged, list2[j])
			j++
		}
	}
	for i < len(list1) {
		merged = append(merged, list1[i])
		i++
	}
	for j < len(list2) {
		merged = append(merged, list2[j])
		j++
	}
	return merged
}

func main() {
	// Test case 1
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}
	fmt.Println(mergeTwoSlices(list1, list2)) // Output: [1 1 2 3 4 4]

	// Test case 2
	list1 = []int{}
	list2 = []int{}
	fmt.Println(mergeTwoSlices(list1, list2)) // Output: []

	// Test case 3
	list1 = []int{}
	list2 = []int{0}
	fmt.Println(mergeTwoSlices(list1, list2)) // Output: [0]
}
