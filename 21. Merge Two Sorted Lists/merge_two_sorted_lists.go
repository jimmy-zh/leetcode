package merge_two_sorted_lists

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsTwoPointer(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{Val: 0, Next: nil}
	pre, cur1, cur2 := preHead, l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			pre.Next, pre, cur1 = cur1, cur1, cur1.Next
		} else {
			pre.Next, pre, cur2 = cur2, cur2, cur2.Next
		}
	}
	if cur1 == nil {
		pre.Next = cur2
	} else {
		pre.Next = cur1
	}
	return preHead.Next
}

func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l2.Next = mergeTwoListsRecursion(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoListsRecursion(l1.Next, l2)
	return l1
}
