package convert_sorted_array_to_binary_search_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestSortedArrayToBST(t *testing.T) {
	vals := []int{-10, -3, 0, 5, 9}
	valsTree := sortedArrayToBST(vals)
	result := []interface{}{0, -3, 9, -10, nil, 5}
	resTree := utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(valsTree, resTree) {
		t.Error("sortedArrayToBST error")
	}

	vals = []int{1}
	valsTree = sortedArrayToBST(vals)
	result = []interface{}{1}
	resTree = utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(valsTree, resTree) {
		t.Error("sortedArrayToBST error")
	}

	vals = []int{}
	valsTree = sortedArrayToBST(vals)
	result = []interface{}{}
	resTree = utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(valsTree, resTree) {
		t.Error("sortedArrayToBST error")
	}

}
