package delete_node_in_a_BST

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

/*
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
Note: Time complexity should be O(height of tree).

Example:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

Given key to delete is 3. So we find the node with value 3 and delete it.

One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

    5
   / \
  4   6
 /     \
2       7

Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
*/

func deleteNode(root *utils.TreeNode, key int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}
		if root.Left.Right == nil {
			root.Left.Right = root.Right
			return root.Left
		}
		pn, n := root.Left, root.Left.Right
		for n != nil && n.Right != nil {
			pn, n = n, n.Right
		}
		pn.Right = n.Left
		n.Left = root.Left
		n.Right = root.Right
		return n
	}
	return root
}
