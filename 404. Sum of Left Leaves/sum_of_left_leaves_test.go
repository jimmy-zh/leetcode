package sum_of_left_leaves

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestSumOfLeftLeaves(t *testing.T) {
	vals := []interface{}{}
	tree := utils.BinaryTreeConstructor(vals)
	result := 0
	if result != sumOfLeftLeaves(tree) {
		t.Error("sumOfLeftLeaves error")
	}

	vals = []interface{}{1}
	tree = utils.BinaryTreeConstructor(vals)
	result = 0
	if result != sumOfLeftLeaves(tree) {
		t.Error("sumOfLeftLeaves error")
	}

	vals = []interface{}{1, 2, 3, 4, nil, 5, 6, nil, nil, nil, 7, 8, nil}
	tree = utils.BinaryTreeConstructor(vals)
	result = 12
	if result != sumOfLeftLeaves(tree) {
		t.Error("sumOfLeftLeaves error")
	}
}
