package DeleteNodeInLinkedList

type ListNode struct {
	 Val int
	 Next *ListNode
}
func ToList(node *ListNode) []int {
	var ans []int
	for {
		ans = append(ans, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return ans
}
func DeleteNode(node *ListNode) {
	if node.Next != nil {
		*node = *node.Next
	}
}
