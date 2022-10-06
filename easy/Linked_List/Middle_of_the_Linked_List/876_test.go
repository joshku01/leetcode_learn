package Middle_of_the_Linked_List

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				head: BuildList([]int{1, 2, 3, 4, 5, 6}),
			},
			want: BuildList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BuildList(nums []int) *ListNode {
	var result *ListNode
	var currentNode *ListNode
	for idx, val := range nums {
		tempNode := ListNode{Val: val, Next: nil}
		if idx == 0 {
			result = &tempNode
			currentNode = &tempNode
		} else {
			currentNode.Next = &tempNode
			currentNode = currentNode.Next
		}
	}
	return result
}
