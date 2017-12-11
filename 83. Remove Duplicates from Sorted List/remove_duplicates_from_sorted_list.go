package remove_duplicates_from_sorted_list

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//modify duplicate node's value
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur, cursor := head, head.Next
	for cursor != nil {
		if cur.Val != cursor.Val {
			cur.Next.Val = cursor.Val
			cur = cur.Next
		}
		cursor = cursor.Next
	}
	cur.Next = nil
	return head
}

//drop duplicate node directly
func deleteDuplicatesPro(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur, cursor := head, head.Next
	for cursor != nil {
		if cur.Val != cursor.Val {
			cur.Next = cursor
			cur = cursor
		}
		cursor = cursor.Next
	}
	cur.Next = nil
	return head
}
