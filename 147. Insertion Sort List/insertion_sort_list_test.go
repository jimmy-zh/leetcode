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
	resultVal := []int{1, 2, 4, 5, 1000}
	head := insertionSortListIteration(firstNode)
	for i := 0; i < 5; i++ {
		if head.Val != resultVal[i] {
			t.Error("Not Sorted")
		}
		head = head.Next
	}
}

func TestInsertionSortListRecursion(t *testing.T) {
	fifthNode := &ListNode{1, nil}
	fourthNode := &ListNode{5, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{4, thirdNode}
	firstNode := &ListNode{1000, secondNode}
	resultVal := []int{1, 2, 4, 5, 1000}
	head := insertionSortListRecursion(firstNode)
	for i := 0; i < 5; i++ {
		if head.Val != resultVal[i] {
			t.Error("Not Sorted")
		}
		head = head.Next
	}
}
