package symmetric_tree

import (
	"github.com/midnight-vivian/go-data-structures/src/data-structures/queue"
	"github.com/midnight-vivian/go-data-structures/src/data-structures/stack"
	"github.com/midnight-vivian/go-data-structures/src/utils"
)

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
Note:
Bonus points if you could solve it both recursively and iteratively.
*/

func isSymmetricRecursion(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecursionInner(root.Left, root.Right)
}

func isSymmetricRecursionInner(n1, n2 *utils.TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil || n1.Val != n2.Val {
		return false
	}
	if isSymmetricRecursionInner(n1.Left, n2.Right) && isSymmetricRecursionInner(n1.Right, n2.Left) {
		return true
	}
	return false
}

func isSymmetricIterationDFS(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	s := stack.NewStackSlice()
	s.Push(root.Left)
	s.Push(root.Right)
	for !s.Empty() {
		n1, n2 := s.Pop().(*utils.TreeNode), s.Pop().(*utils.TreeNode)
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		s.Push(n1.Left)
		s.Push(n2.Right)
		s.Push(n1.Right)
		s.Push(n2.Left)
	}
	return true
}

func isSymmetricIterationBFS(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	q := queue.NewQueueSlice()
	q.Enqueue(root.Left)
	q.Enqueue(root.Right)
	for !q.Empty() {
		n1, n2 := q.Dequeue().(*utils.TreeNode), q.Dequeue().(*utils.TreeNode)
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		q.Enqueue(n1.Left)
		q.Enqueue(n2.Right)
		q.Enqueue(n1.Right)
		q.Enqueue(n2.Left)
	}
	return true
}
