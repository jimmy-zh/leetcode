package insertion_sort_list

import (
	"testing"
)

func TestInsertionSortListIteration(t *testing.T) {
	fifthNode := &ListNode{1, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{4, thirdNode}
	firstNode := &ListNode{1000, secondNode}

	head := insertionSortListIteration(firstNode)
	if head.Val != 1 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 2 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 4 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 5 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 1000 {
		t.Error("Not Sorted")
	}
}

func TestInsertionSortListRecursion(t *testing.T) {
	fifthNode := &ListNode{1, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{4, thirdNode}
	firstNode := &ListNode{1000, secondNode}

	head := insertionSortListRecursion(firstNode)
	if head.Val != 1 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 2 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 4 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 5 {
		t.Error("Not Sorted")
	}
	head = head.Next
	if head.Val != 1000 {
		t.Error("Not Sorted")
	}
}
