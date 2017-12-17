package palindrome_linked_list

/*
Given a singly linked list, determine if it is a palindrome.

Follow up:
Could you do it in O(n) time and O(1) space?
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	//find the middle node of linked list
	one, two := head, head
	for two.Next != nil && two.Next.Next != nil {
		one, two = one.Next, two.Next.Next
	}
	//reverse the first half of linked list
	var pre *ListNode
	cursor := head
	for cursor != one {
		next := cursor.Next
		cursor.Next = pre
		pre, cursor = cursor, next
	}
	//find proper first half head and second half head
	var leftHead, rightHead *ListNode
	if two.Next == nil {
		leftHead, rightHead = pre, one.Next
	} else {
		leftHead, rightHead = one, one.Next
		one.Next = pre
	}
	//compare two half
	for leftHead != nil && rightHead != nil {
		if leftHead.Val != rightHead.Val {
			return false
		}
		leftHead, rightHead = leftHead.Next, rightHead.Next
	}
	return true
}
