package linked_list_cycle_II

import (
	"testing"
)

func TestDetectCycle(t *testing.T) {
	fifthNode := &ListNode{1, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{4, thirdNode}
	firstNode := &ListNode{1000, secondNode}

	if detectCycle(firstNode) != nil {
		t.Error("has no cycle")
	}
	//head is on the cycle
	fifthNode.Next = firstNode
	if detectCycle(firstNode) != firstNode {
		t.Error("detectCycle error")
	}
	fifthNode.Next = secondNode
	if detectCycle(firstNode) != secondNode {
		t.Error("detectCycle error")
	}
	fifthNode.Next = fourthNode
	if detectCycle(firstNode) != fourthNode {
		t.Error("detectCycle error")
	}
}
