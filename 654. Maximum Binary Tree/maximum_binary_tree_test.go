package maximum_binary_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	arr := []int{3, 6, 9, 4, 7, 2, 0}
	result := []interface{}{9, 6, 7, 3, nil, 4, 2, nil, nil, nil, nil, nil, 0}
	tResult := utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(constructMaximumBinaryTree(arr), tResult) {
		t.Error("ConstructMaximumBinaryTree error")
	}
}
