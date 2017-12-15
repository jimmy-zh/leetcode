package reverse_linked_list

/*
Reverse a singly linked list.
click to show more hints.

Hint:
A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}

func reverseListIteration(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre, head = head, next
	}
	return pre
}
