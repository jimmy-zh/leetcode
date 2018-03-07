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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	last, cur := head, head.Next
	for cur != nil {
		if cur.Val != last.Val {
			last.Next, last = cur, cur
		}
		cur = cur.Next
	}
	last.Next = nil
	return head
}
