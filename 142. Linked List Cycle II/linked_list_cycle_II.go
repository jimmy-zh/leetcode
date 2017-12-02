package linked_list_cycle_II

/*
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

Note: Do not modify the linked list.

Follow up:
Can you solve it without using extra space?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	cycle := false
	one, two := head, head
	for two != nil && two.Next != nil {
		one, two = one.Next, two.Next.Next
		if one == two {
			cycle = true
			break
		}
	}
	if !cycle {
		return nil
	}
	one = head
	for {
		one = one.Next
		two = two.Next
		if one == two {
			return one
		}
	}
}
