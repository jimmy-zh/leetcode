package maximum_depth_of_binary_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestMaxDepth(t *testing.T) {
	vals := []interface{}{}
	tree := utils.BinaryTreeConstructor(vals)
	result := 0
	if result != maxDepth(tree) {
		t.Error("maxDepth error")
	}
	vals = []interface{}{1}
	tree = utils.BinaryTreeConstructor(vals)
	result = 1
	if result != maxDepth(tree) {
		t.Error("maxDepth error")
	}
	vals = []interface{}{1, 2, 3, 4, nil, 5, nil, nil, nil, nil, 7}
	tree = utils.BinaryTreeConstructor(vals)
	result = 4
	if result != maxDepth(tree) {
		t.Error("maxDepth error")
	}
}
