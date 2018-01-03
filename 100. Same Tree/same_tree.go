package same_tree

import (
	"github.com/midnight-vivian/go-data-structures/data-structures/queue"
	"github.com/midnight-vivian/go-data-structures/data-structures/stack"
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given two binary trees, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:
Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true

Example 2:
Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false

Example 3:
Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/

func isSameTreeRecursion(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTreeRecursion(p.Left, q.Left) && isSameTreeRecursion(p.Right, q.Right)
}

func isSameTreeIterationDFS(p *utils.TreeNode, q *utils.TreeNode) bool {
	s := stack.NewStackSlice()
	s.Push(p)
	s.Push(q)
	for !s.Empty() {
		n1, n2 := s.Pop().(*utils.TreeNode), s.Pop().(*utils.TreeNode)
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		s.Push(n1.Left)
		s.Push(n2.Left)
		s.Push(n1.Right)
		s.Push(n2.Right)
	}
	return true
}

func isSameTreeIterationBFS(p *utils.TreeNode, q *utils.TreeNode) bool {
	qq := queue.NewQueueSlice()
	qq.Enqueue(p)
	qq.Enqueue(q)
	for !qq.Empty() {
		n1, n2 := qq.Dequeue().(*utils.TreeNode), qq.Dequeue().(*utils.TreeNode)
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		qq.Enqueue(n1.Left)
		qq.Enqueue(n2.Left)
		qq.Enqueue(n1.Right)
		qq.Enqueue(n2.Right)
	}
	return true
}
