package Convert_Binary_Number_in_a_Linked_List_to_Integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	var out int
	for _, v := range *head {
		out = out + v
	}

	return out
}
