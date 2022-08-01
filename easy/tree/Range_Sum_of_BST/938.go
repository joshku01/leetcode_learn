package Range_Sum_of_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {

	result := 0
	result = transversal(root, low, high)

	return result
}

func transversal(root *TreeNode, L int, R int) int {
	count := 0
	// 條件 大於L 小於R
	if root.Val <= R && root.Val >= L {
		count += root.Val
	}
	if root.Left != nil && root.Val != L {
		count += transversal(root.Left, L, R)
	}
	if root.Right != nil && root.Val != R {
		count += transversal(root.Right, L, R)
	}

	return count
}
