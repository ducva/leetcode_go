package InvertBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Right
	root.Right = InvertTree(root.Left)
	root.Left = InvertTree(temp)
	return root
}
