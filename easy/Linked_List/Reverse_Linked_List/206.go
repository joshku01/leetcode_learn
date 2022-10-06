package Reverse_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		prev, current, current.Next = current, current.Next, prev
	}

	return prev
}
