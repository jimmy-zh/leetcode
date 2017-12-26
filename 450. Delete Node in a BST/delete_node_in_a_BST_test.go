package delete_node_in_a_BST

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestDeleteNode(t *testing.T) {
	vals := []interface{}{10, 4, 15, 3, 8, 11, 22, nil, nil, nil, nil, nil, 13, nil, nil, 12}
	tVals := utils.BinaryTreeConstructor(vals)
	result := []interface{}{10, 4, 15, nil, 8, 11, 22, nil, nil, nil, 13, nil, nil, 12}
	tResult := utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(deleteNode(tVals, 3), tResult) {
		t.Error("DeleteNode error")
	}

	tVals = utils.BinaryTreeConstructor(vals)
	result = []interface{}{10, 3, 15, nil, 8, 11, 22, nil, nil, nil, 13, nil, nil, 12}
	tResult = utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(deleteNode(tVals, 4), tResult) {
		t.Error("DeleteNode error")
	}

	tVals = utils.BinaryTreeConstructor(vals)
	result = []interface{}{10, 4, 13, 3, 8, 11, 22, nil, nil, nil, nil, nil, 12, nil, nil}
	tResult = utils.BinaryTreeConstructor(result)
	if !utils.BinaryTreeEqual(deleteNode(tVals, 15), tResult) {
		t.Error("DeleteNode error")
	}

}
