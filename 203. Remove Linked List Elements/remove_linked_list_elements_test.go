package remove_linked_list_elements

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	fifthNode := &ListNode{2, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{1, thirdNode}
	firstNode := &ListNode{2, secondNode}
	result := []int{
		1,
		5,
	}
	node := removeElements(firstNode, 2)
	for i := range result {
		if result[i] != node.Val {
			t.Error("remove elements error")
		}
		node = node.Next
	}
	if node != nil {
		t.Error("remove elements error")
	}

}
