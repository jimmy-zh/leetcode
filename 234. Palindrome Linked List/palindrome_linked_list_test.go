package palindrome_linked_list

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fifthNode := &ListNode{10, nil}
	fourthNode := &ListNode{3, fifthNode}
	thirdNode := &ListNode{2, fourthNode}
	secondNode := &ListNode{3, thirdNode}
	firstNode := &ListNode{10, secondNode}
	if !isPalindrome(firstNode) {
		t.Error("should be palindrome")
	}
	fifthNode = &ListNode{10, nil}
	fourthNode = &ListNode{2, fifthNode}
	thirdNode = &ListNode{2, fourthNode}
	secondNode = &ListNode{3, thirdNode}
	firstNode = &ListNode{10, secondNode}
	if isPalindrome(firstNode) {
		t.Error("should not be palindrome")
	}
	fourthNode = &ListNode{10, nil}
	thirdNode = &ListNode{3, fourthNode}
	secondNode = &ListNode{3, thirdNode}
	firstNode = &ListNode{10, secondNode}
	if !isPalindrome(firstNode) {
		t.Error("should be palindrome")
	}
	fourthNode = &ListNode{10, nil}
	thirdNode = &ListNode{2, fourthNode}
	secondNode = &ListNode{3, thirdNode}
	firstNode = &ListNode{10, secondNode}
	if isPalindrome(firstNode) {
		t.Error("should not be palindrome")
	}
	firstNode = &ListNode{10, nil}
	if !isPalindrome(firstNode) {
		t.Error("should be palindrome")
	}
	secondNode = &ListNode{10, nil}
	firstNode = &ListNode{10, secondNode}
	if !isPalindrome(firstNode) {
		t.Error("should be palindrome")
	}
	secondNode = &ListNode{3, nil}
	firstNode = &ListNode{10, secondNode}
	if isPalindrome(firstNode) {
		t.Error("should not be palindrome")
	}
}
