package convert_BST_to_greater_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestConvertBST(t *testing.T) {
	vals := []interface{}{5, 2, 13, 1, nil, 7, 28, nil, nil, nil, 8, nil, nil}
	valsTree := utils.BinaryTreeConstructor(vals)
	results := []interface{}{46, 2, 41, 1, nil, 15, 28, nil, nil, nil, 8, nil, nil}
	resTree := utils.BinaryTreeConstructor(results)
	if !utils.BinaryTreeEqual(convertBST(valsTree), resTree) {
		t.Error("convert error")
	}
}
