package invert_binary_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/src/utils"
)

func TestInvertTree(t *testing.T) {
	vals := []interface{}{}
	tVal := utils.BinaryTreeConstructor(vals)
	result := []interface{}{}
	tResult := utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(invertTree(tVal), tResult) {
		t.Error("invertTree error")
	}

	vals = []interface{}{1}
	tVal = utils.BinaryTreeConstructor(vals)
	result = []interface{}{1}
	tResult = utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(invertTree(tVal), tResult) {
		t.Error("invertTree error")
	}

	vals = []interface{}{1, 2, 3, 4, nil, 5, nil, nil, nil, 6, 7}
	tVal = utils.BinaryTreeConstructor(vals)
	result = []interface{}{1, 3, 2, nil, 5, nil, 4, 7, 6, nil, nil}
	tResult = utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(invertTree(tVal), tResult) {
		t.Error("invertTree error")
	}

}
