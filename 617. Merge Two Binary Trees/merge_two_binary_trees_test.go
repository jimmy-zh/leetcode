package merge_two_binary_trees

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestMergeTrees(t *testing.T) {
	vals := [][][]interface{}{
		[][]interface{}{
			{},
			{},
			{},
		},

		[][]interface{}{
			{1},
			{},
			{1},
		},
		[][]interface{}{
			{},
			{1},
			{1},
		},
		[][]interface{}{
			{1, 3, 2, 5},
			{5, 4, 8, nil, nil, 7, 10},
			{6, 7, 10, 5, nil, 7, 10},
		},
		[][]interface{}{
			{8, 4, 2, nil, 1, 5, nil, nil, nil, nil, 3},
			{9, 7, 6, nil, 6, nil, 8, 4, 5, 10, nil},
			{17, 11, 8, nil, 7, 5, 8, 4, 5, nil, 3, 10},
		},
	}

	for i := range vals {
		t1 := utils.BinaryTreeConstructor(vals[i][0])
		t2 := utils.BinaryTreeConstructor(vals[i][1])
		t3 := utils.BinaryTreeConstructor(vals[i][2])
		if !utils.BinaryTreeEqual(mergeTrees(t1, t2), t3) {
			t.Error("mergeTrees error")
		}
	}

}
