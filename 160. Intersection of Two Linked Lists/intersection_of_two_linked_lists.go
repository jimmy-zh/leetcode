package intersection_of_two_linked_lists

/*
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:
A:          a1 → a2
                   ↘
                     c1 → c2 → c3
                   ↗
B:     b1 → b2 → b3
begin to intersect at node c1.

Notes:
If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	one, two := headA, headA
	for {
		one = one.Next
		if one == nil {
			one = headB
		}
		if two.Next == nil {
			two = headB.Next
		} else if two.Next.Next == nil {
			two = headB
		} else {
			two = two.Next.Next
		}
		if one == two {
			break
		}
	}
	one = headA
	for one != two {
		one = one.Next
		if one == nil {
			one = headB
		}
		two = two.Next
		if two == nil {
			two = headB
		}
	}
	if one == headB {
		return nil
	}
	return one
}

func getIntersectionNodePro(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
