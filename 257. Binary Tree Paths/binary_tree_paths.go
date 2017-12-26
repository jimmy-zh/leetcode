package binary_tree_paths

import (
	"strconv"

	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given a binary tree, return all root-to-leaf paths.

For example, given the following binary tree:

   1
 /   \
2     3
 \
  5
All root-to-leaf paths are:

["1->2->5", "1->3"]
*/

func binaryTreePaths(root *utils.TreeNode) []string {
	if root == nil {
		return nil
	}
	paths := make([]string, 0)
	binaryTreePathsRecursion(root, "", &paths)
	return paths
}

func binaryTreePathsRecursion(node *utils.TreeNode, path string, paths *[]string) {
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path+strconv.Itoa(node.Val))
	}
	if node.Left != nil {
		binaryTreePathsRecursion(node.Left, path+strconv.Itoa(node.Val)+`->`, paths)
	}
	if node.Right != nil {
		binaryTreePathsRecursion(node.Right, path+strconv.Itoa(node.Val)+`->`, paths)
	}
}
