package minimum_depth_of_binary_tree

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
*/

func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := minDepth(root.Left)
	rDepth := minDepth(root.Right)
	if (lDepth != 0 && lDepth < rDepth) || rDepth == 0 {
		return lDepth + 1
	}
	return rDepth + 1
}
