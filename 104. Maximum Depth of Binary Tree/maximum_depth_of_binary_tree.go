package maximum_depth_of_binary_tree

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}
