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
	if head == nil {
		return false
	}
	oneStep, twoStep := head, head
	for twoStep != nil && twoStep.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
		if oneStep == twoStep {
			return true
		}
	}
	return false
}
