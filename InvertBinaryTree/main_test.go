package InvertBinaryTree

import (
	"testing"
)

func TestInvertTree(t *testing.T) {
	var ans = InvertTree(&TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		},
	})
	if ans.Val != 1 {
		t.Errorf("Root node value should be: %d, but actual is: %d", 1, ans.Val)
	}
	if ans.Right != nil {
		t.Errorf("Right Node value should be nil, but actual is not nil")
	}
	if ans.Left.Val != 2 {
		t.Errorf("Left Node Value should be %d, but actual is: %d",2, ans.Left.Val)
	}
}