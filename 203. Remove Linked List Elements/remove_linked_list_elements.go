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
	preHead := &ListNode{0, head}
	for pre, cur := preHead, head; cur != nil; cur = cur.Next {
		if cur.Val != val {
			pre = cur
		} else {
			pre.Next = cur.Next
		}
	}
	return preHead.Next
}
