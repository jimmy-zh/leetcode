package delete_node_in_a_linked_list

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	fifthNode := &ListNode{1, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{4, thirdNode}
	firstNode := &ListNode{1000, secondNode}
	values := []int{1000, 4, 5, 1}
	deleteNode(thirdNode)
	n := firstNode
	for i := range values {
		if n == nil {
			t.Fatal("delete node error")
		}
		if values[i] != n.Val {
			t.Error("delete node error")
		}
		n = n.Next
	}

}
