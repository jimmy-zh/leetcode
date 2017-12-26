package invert_binary_tree

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Invert a binary tree.
     4
   /   \
  2     7
 / \   / \
1   3 6   9
to
     4
   /   \
  7     2
 / \   / \
9   6 3   1
Trivia:
This problem was inspired by this original tweet by Max Howell:
Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so fuck off.
*/

func invertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	leftTmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(leftTmp)
	return root
}
