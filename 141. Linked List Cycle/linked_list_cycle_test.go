package linked_list_cycle

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	fifthNode := &ListNode{1, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{4, thirdNode}
	firstNode := &ListNode{1000, secondNode}

	if hasCycle(firstNode) {
		t.Error("has no cycle")
	}
	//head is on the cycle
	fifthNode.Next = firstNode
	if !hasCycle(firstNode) {
		t.Error("has cycle")
	}
	fifthNode.Next = secondNode
	if !hasCycle(firstNode) {
		t.Error("has cycle")
	}
	fifthNode.Next = fourthNode
	if !hasCycle(firstNode) {
		t.Error("has cycle")
	}
}
