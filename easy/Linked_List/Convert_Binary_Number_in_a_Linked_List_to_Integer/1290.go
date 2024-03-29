package Convert_Binary_Number_in_a_Linked_List_to_Integer

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	result := 0
	binaryStr := ""
	for head != nil {
		binaryStr += strconv.Itoa(head.Val)
		head = head.Next
	}
	result64, _ := strconv.ParseInt(binaryStr, 2, 0)
	result = int(result64)
	return result
}
