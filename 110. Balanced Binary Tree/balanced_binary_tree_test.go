package balanced_binary_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestIsBalanced(t *testing.T) {
	vals := []interface{}{}
	result := true
	tVal := utils.BinaryTreeConstructor(vals)
	if isBalanced(tVal) != result {
		t.Error("isBalanced error")
	}

	vals = []interface{}{1}
	result = true
	tVal = utils.BinaryTreeConstructor(vals)
	if isBalanced(tVal) != result {
		t.Error("isBalanced error")
	}

	vals = []interface{}{5, 8, 7, 10, nil, 6, 9, nil, nil, 4, nil, nil, nil}
	result = true
	tVal = utils.BinaryTreeConstructor(vals)
	if isBalanced(tVal) != result {
		t.Error("isBalanced error")
	}
	vals = []interface{}{5, 8, 7, nil, nil, 6, 9, 4, nil, nil, nil}
	result = false
	tVal = utils.BinaryTreeConstructor(vals)
	if isBalanced(tVal) != result {
		t.Error("isBalanced error")
	}
}
