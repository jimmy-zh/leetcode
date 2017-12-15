package reverse_linked_list

import (
	"testing"
)

func TestReverseListRecursion(t *testing.T) {
	fifthNode := &ListNode{2, nil}
	fourthNode := &ListNode{4, fifthNode}
	thirdNode := &ListNode{6, fourthNode}
	secondNode := &ListNode{3, thirdNode}
	firstNode := &ListNode{1, secondNode}
	result := []int{
		2,
		4,
		6,
		3,
		1,
	}
	node := reverseListRecursion(firstNode)
	for i := range result {
		if result[i] != node.Val {
			t.Error("reverse list error")
		}
		node = node.Next
	}
	if node != nil {
		t.Error("reverse list error")
	}

}

func TestReverseListIteration(t *testing.T) {
	fifthNode := &ListNode{2, nil}
	fourthNode := &ListNode{4, fifthNode}
	thirdNode := &ListNode{6, fourthNode}
	secondNode := &ListNode{3, thirdNode}
	firstNode := &ListNode{1, secondNode}
	result := []int{
		2,
		4,
		6,
		3,
		1,
	}
	node := reverseListIteration(firstNode)
	for i := range result {
		if result[i] != node.Val {
			t.Error("reverse list error")
		}
		node = node.Next
	}
	if node != nil {
		t.Error("reverse list error")
	}
}
