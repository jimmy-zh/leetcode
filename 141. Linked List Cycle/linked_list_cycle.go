package linked_list_cycle

/*
Given a linked list, determine if it has a cycle in it.
Follow up:
Can you solve it without using extra space?
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//two pointer O(1)
func hasCycle(head *ListNode) bool {
	one, two := head, head
	for two != nil && two.Next != nil {
		one, two = one.Next, two.Next.Next
		if one == two {
			return true
		}
	}
	return false
}
