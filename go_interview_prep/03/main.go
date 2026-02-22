package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers returns a new list whose digits (reverse order) are the sum of l1 and l2.
// Input lists are reverse order (head = ones place); output is also reverse order.
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}

	return dummy.Next
}

func sliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	tail := head
	for i := 1; i < len(vals); i++ {
		tail.Next = &ListNode{Val: vals[i]}
		tail = tail.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var out []int
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func main() {
	// 342 + 465 = 807 => [2,4,3] + [5,6,4] => [7,0,8]
	l1 := sliceToList([]int{2, 4, 3})
	l2 := sliceToList([]int{5, 6, 4})
	fmt.Println(listToSlice(addTwoNumbers(l1, l2)))

	// 9 + 9 = 18 => [9] + [9] => [8,1]
	fmt.Println(listToSlice(addTwoNumbers(sliceToList([]int{9}), sliceToList([]int{9}))))
}
