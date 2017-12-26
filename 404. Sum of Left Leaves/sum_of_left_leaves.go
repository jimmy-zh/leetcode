package sum_of_left_leaves

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Find the sum of all left leaves in a given binary tree.

Example:
    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

func sumOfLeftLeaves(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Right) + sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	return sum
}
