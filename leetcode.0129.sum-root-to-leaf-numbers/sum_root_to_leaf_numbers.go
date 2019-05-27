package leetcode0129

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumLeftRight(0, root)
}

func sumLeftRight(pre int, root *TreeNode) int {
	val := pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return val
	}
	if root.Left == nil {
		return sumLeftRight(val, root.Right)
	}
	if root.Right == nil {
		return sumLeftRight(val, root.Left)
	}
	return sumLeftRight(val, root.Left) + sumLeftRight(val, root.Right)
}
