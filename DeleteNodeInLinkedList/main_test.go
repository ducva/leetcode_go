package DeleteNodeInLinkedList

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	var head = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	DeleteNode(head.Next)
	var actual = ToList(head)
	var expected = ToList(&ListNode{
		Val:  4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Failed. actual: %v", actual)
	}
}