package balanced_binary_tree

import "github.com/midnight-vivian/go-data-structures/utils"

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
*/

func isBalanced(root *utils.TreeNode) bool {
	b, _ := isBalancedInner(root)
	return b
}

func isBalancedInner(node *utils.TreeNode) (isBalanced bool, height int) {
	if node == nil {
		return true, 0
	}
	lb, lh := isBalancedInner(node.Left)
	if !lb {
		return false, 0
	}
	rb, rh := isBalancedInner(node.Right)
	if !rb {
		return false, 0
	}
	if lh >= rh && lh-rh <= 1 {
		return true, lh + 1
	}
	if rh == lh+1 {
		return true, rh + 1
	}
	return false, 0
}
