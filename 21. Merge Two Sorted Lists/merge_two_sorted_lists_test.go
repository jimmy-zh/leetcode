package merge_two_sorted_lists

import (
	"testing"
)

func TestMergeTwoListsTwoPointer(t *testing.T) {
	l1fourth := &ListNode{11, nil}
	l1third := &ListNode{10, l1fourth}
	l1second := &ListNode{7, l1third}
	l1first := &ListNode{5, l1second}

	l2third := &ListNode{12, nil}
	l2second := &ListNode{10, l2third}
	l2first := &ListNode{4, l2second}
	result := []int{
		4, 5, 7, 10, 10, 11, 12,
	}
	l3 := mergeTwoListsTwoPointer(l1first, l2first)
	testHelp(t, l3, result)

	l1fourth = &ListNode{11, nil}
	l1third = &ListNode{10, l1fourth}
	l1second = &ListNode{7, l1third}
	l1first = &ListNode{5, l1second}

	l2third = &ListNode{8, nil}
	l2second = &ListNode{4, l2third}
	l2first = &ListNode{3, l2second}
	result = []int{
		3, 4, 5, 7, 8, 10, 11,
	}
	l3 = mergeTwoListsTwoPointer(l1first, l2first)
	testHelp(t, l3, result)
	l1first = nil
	l2first = nil
	result = []int{}
	l3 = mergeTwoListsTwoPointer(l1first, l2first)
	testHelp(t, l3, result)
}
func TestMergeTwoListsRecursion(t *testing.T) {
	l1fourth := &ListNode{11, nil}
	l1third := &ListNode{10, l1fourth}
	l1second := &ListNode{7, l1third}
	l1first := &ListNode{5, l1second}

	l2third := &ListNode{12, nil}
	l2second := &ListNode{10, l2third}
	l2first := &ListNode{4, l2second}
	result := []int{
		4, 5, 7, 10, 10, 11, 12,
	}
	l3 := mergeTwoListsRecursion(l1first, l2first)
	testHelp(t, l3, result)

	l1fourth = &ListNode{11, nil}
	l1third = &ListNode{10, l1fourth}
	l1second = &ListNode{7, l1third}
	l1first = &ListNode{5, l1second}

	l2third = &ListNode{8, nil}
	l2second = &ListNode{4, l2third}
	l2first = &ListNode{3, l2second}
	result = []int{
		3, 4, 5, 7, 8, 10, 11,
	}
	l3 = mergeTwoListsRecursion(l1first, l2first)
	testHelp(t, l3, result)

	l1first = nil
	l2first = nil
	result = []int{}
	l3 = mergeTwoListsRecursion(l1first, l2first)
	testHelp(t, l3, result)
}

func testHelp(t *testing.T, l3 *ListNode, result []int) {
	for i := range result {
		if result[i] != l3.Val {
			t.Error("merge error")
		}
		l3 = l3.Next
	}
	if l3 != nil {
		t.Error("merge error")
	}
}
