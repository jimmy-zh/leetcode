package remove_linked_list_elements

/*
Remove all elements from a linked list of integers that have value val.

Example
Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
Return: 1 --> 2 --> 3 --> 4 --> 5
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur, cursor := preHead, head
	for cursor != nil {
		if cursor.Val != val {
			cur.Next = cursor
			cur = cursor
		}
		cursor = cursor.Next
	}
	cur.Next = nil
	return preHead.Next
}
