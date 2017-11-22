package insertion_sort_list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Iteration
func insertionSortListIteration(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	preHead := &ListNode{0, head}
	for lastOut, curOut := head, head.Next; curOut != nil; lastOut, curOut = curOut, curOut.Next {
		for lastIn, curIn := preHead, preHead.Next; curIn != curOut; lastIn, curIn = curIn, curIn.Next {
			if curOut.Val < curIn.Val {
				lastOut.Next = curOut.Next
				lastIn.Next = curOut
				curOut.Next = curIn
				curOut = lastOut
				break
			}
		}
	}
	return preHead.Next
}

//Recursion + exchange values
func insertionSortListRecursion(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	insertionSortListRecursion(head.Next)
	for i := head; i.Next != nil && i.Val > i.Next.Val; i = i.Next {
		i.Val, i.Next.Val = i.Next.Val, i.Val
	}
	return head
}
