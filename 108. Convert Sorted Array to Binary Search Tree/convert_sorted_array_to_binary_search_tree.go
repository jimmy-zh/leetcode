package convert_sorted_array_to_binary_search_tree

import "github.com/midnight-vivian/go-data-structures/src/utils"

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:
Given the sorted array: [-10,-3,0,5,9],
One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5

*/

func sortedArrayToBST(nums []int) *utils.TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	return &utils.TreeNode{
		Val:   nums[l/2],
		Left:  sortedArrayToBST(nums[:l/2]),
		Right: sortedArrayToBST(nums[l/2+1:]),
	}
}
