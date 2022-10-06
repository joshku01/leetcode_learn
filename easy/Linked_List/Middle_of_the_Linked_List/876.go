package Middle_of_the_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.
func middleNode(head *ListNode) *ListNode {
	//思考方向,定義兩個指針,速度不同,快的為慢的兩倍,當快的到達終點,慢的即為中間點
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}
