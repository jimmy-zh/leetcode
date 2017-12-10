package intersection_of_two_linked_lists

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	fifthA := &ListNode{1, nil}
	fourthA := &ListNode{5, fifthA}
	thirdA := &ListNode{2, fourthA}
	secondA := &ListNode{4, thirdA}
	firstA := &ListNode{10, secondA}

	secondB := &ListNode{34, nil}
	firstB := &ListNode{23, secondB}
	if getIntersectionNode(firstA, firstB) != nil {
		t.Error("wrong node")
	}

	secondB = &ListNode{34, secondA}
	firstB = &ListNode{23, secondB}
	if getIntersectionNode(firstA, firstB) != secondA {
		t.Error("wrong node")
	}

	secondB = &ListNode{34, fifthA}
	firstB = &ListNode{23, secondB}
	if getIntersectionNode(firstA, firstB) != fifthA {
		t.Error("wrong node")
	}

	secondB = &ListNode{34, nil}
	firstB = &ListNode{23, secondB}
	if getIntersectionNodePro(firstA, firstB) != nil {
		t.Error("wrong node")
	}

	secondB = &ListNode{34, secondA}
	firstB = &ListNode{23, secondB}
	if getIntersectionNodePro(firstA, firstB) != secondA {
		t.Error("wrong node")
	}

	secondB = &ListNode{34, fifthA}
	firstB = &ListNode{23, secondB}
	if getIntersectionNodePro(firstA, firstB) != fifthA {
		t.Error("wrong node")
	}
}
