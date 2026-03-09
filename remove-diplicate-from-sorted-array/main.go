package main

import "fmt"

func removeDuplicate(nums []int) int {
	count := len(nums)
	y := 1

	for x := 1; x < count; x++ {
		if nums[x] != nums[x-1] {
			nums[y] = nums[x]
			y++
		}
	}
	return y
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}

	newLength := removeDuplicate(nums)

	fmt.Println("Length after removing duplicates:", newLength)
	fmt.Println("Array after removing duplicates:", nums[:newLength])
}
