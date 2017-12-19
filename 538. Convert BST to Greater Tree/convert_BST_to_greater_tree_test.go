package convert_BST_to_greater_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestConvertBST(t *testing.T) {
	vals := []interface{}{5, 2, 13, 1, 3, 7, 28}
	valsTree := utils.BinaryTreeConstructor(vals)
	results := []interface{}{46, 5, 41, 1, 3, 7, 28}
	resTree := utils.BinaryTreeConstructor(results)
	if !utils.BinaryTreeEqual(convertBST(valsTree), resTree) {
		t.Error("convert error")
	}
}
