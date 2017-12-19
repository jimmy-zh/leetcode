package convert_BST_to_greater_tree

import "github.com/midnight-vivian/go-data-structures/src/utils"

/*
Given a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus sum of all keys greater than the original key in BST.

Example:
Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
*/

func convertBST(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	root.Right = convertBST(root.Right)
	root.Left = convertBST(root.Left)
	if root.Right != nil {
		root.Val += root.Right.Val
	}
	return root
}
