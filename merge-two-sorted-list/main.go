package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.val <= list2.val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return dummy.Next
}

func main() {

	list1 := [1, 2, 4]
	list2 := [1, 3, 4]
	fmt.Println(mergeTwoList(list1, list2)) // Output: [1 1 2 3 4 4]

	list1 = []int{}
	list2 = []int{}
	fmt.Println(mergeTwoList(list1, list2))

	list1 = []int{}
	list2 = []int{0}
	fmt.Println(mergeTwoList(list1, list2))
}
