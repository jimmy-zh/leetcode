package trim_a_binary_search_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestTrimBST(t *testing.T) {
	vals := []interface{}{7, 4, 10, 2, 5, 8, 11, 1, 3, nil, 6, nil, 9, nil, nil}
	tVal := utils.BinaryTreeConstructor(vals)
	L := 6
	R := 10
	result := []interface{}{7, 6, 10, nil, nil, 8, nil, nil, 9}
	tResult := utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(trimBST(tVal, L, R), tResult) {
		t.Error("trimBST error")
	}
}
