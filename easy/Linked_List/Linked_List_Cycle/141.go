package Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// 題意:給定一個LinkList的head,我們要判斷此LinkList是否有環(Cycle).
// 思考方向: 我們可以把經過的節點都記錄下來,遍歷整個LinkList,若有cycle則會有重複經過的點,return true ,如果都沒有重複的節點則表示無cycle.
func hasCycle(head *ListNode) bool {
	// 用hash map來紀錄是否有遍歷
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}
