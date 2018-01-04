package subtree_of_anothert_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestIsSubtree(t *testing.T) {
	vals := [][][]interface{}{
		[][]interface{}{
			{},
			{},
		},
		[][]interface{}{
			{1},
			{},
		},
		[][]interface{}{
			{},
			{1},
		},
		[][]interface{}{
			{3, 4, 5, 1, 2},
			{4, 1, 2},
		},
		[][]interface{}{
			{3, 4, 5, 1, 2, nil, nil, 5},
			{4, 1, 2},
		},
	}
	result := []bool{
		true,
		false,
		false,
		true,
		false,
	}
	for i := range vals {
		ts := utils.BinaryTreeConstructor(vals[i][0])
		tt := utils.BinaryTreeConstructor(vals[i][1])
		if isSubtree(ts, tt) != result[i] {
			t.Error("isSubtree error")
		}
	}
}
