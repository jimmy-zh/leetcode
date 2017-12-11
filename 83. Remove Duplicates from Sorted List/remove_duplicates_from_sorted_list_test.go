package remove_duplicates_from_sorted_list

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	fifthNode := &ListNode{4, nil}
	fourthNode := &ListNode{4, fifthNode}
	thirdNode := &ListNode{3, fourthNode}
	secondNode := &ListNode{3, thirdNode}
	firstNode := &ListNode{1, secondNode}
	result := []int{
		1,
		3,
		4,
	}
	deleteDuplicatesPro(firstNode)
	node := firstNode
	for i := range result {
		if result[i] != node.Val {
			t.Error("delete duplicates error")
		}
		node = node.Next
	}
	if node != nil {
		t.Error("delete duplicates error")
	}
}
