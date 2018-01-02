package average_of_levels_in_binary_tree

import (
	"container/list"

	"github.com/midnight-vivian/go-data-structures/data-structures/queue"
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
Note:
The range of node's value is in the range of 32-bit signed integer.
*/

func averageOfLevelsMyLib(root *utils.TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	q := queue.NewQueueSlice()
	q.Enqueue(root)
	for !q.Empty() {
		count, sum := q.Size(), 0
		for i := 0; i < count; i++ {
			n := q.Dequeue().(*utils.TreeNode)
			sum += n.Val
			if n.Left != nil {
				q.Enqueue(n.Left)
			}
			if n.Right != nil {
				q.Enqueue(n.Right)
			}
		}
		res = append(res, float64(sum)/float64(count))
	}
	return res
}

func averageOfLevelsStdLib(root *utils.TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	q := list.New()
	q.PushFront(root)
	for q.Len() != 0 {
		count, sum := q.Len(), 0
		for i := 0; i < count; i++ {
			e := q.Back()
			n := e.Value.(*utils.TreeNode)
			q.Remove(e)
			sum += n.Val
			if n.Left != nil {
				q.PushFront(n.Left)
			}
			if n.Right != nil {
				q.PushFront(n.Right)
			}
		}
		res = append(res, float64(sum)/float64(count))
	}
	return res
}

func averageOfLevelsNoLib(root *utils.TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	q := []*utils.TreeNode{root}
	for len(q) != 0 {
		count, sum := len(q), 0
		for i := 0; i < count; i++ {
			n := q[0]
			sum += n.Val
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, float64(sum)/float64(count))
	}
	return res
}
