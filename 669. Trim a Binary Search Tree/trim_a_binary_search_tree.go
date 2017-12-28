package trim_a_binary_search_tree

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2
Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
*/

func trimBST(root *utils.TreeNode, L int, R int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
