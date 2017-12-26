package same_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestIsSameTree(t *testing.T) {
	valsT := [][][]interface{}{
		[][]interface{}{
			[]interface{}{nil},
			[]interface{}{nil},
		},
		[][]interface{}{
			[]interface{}{1, 2, 3, nil, 5, 4, nil},
			[]interface{}{1, 2, 3, nil, 5, 4, nil},
		},
		[][]interface{}{
			[]interface{}{1, 2, 3, nil, 5, 4, nil},
			[]interface{}{1, 2, 3, nil, 5, 4, 7},
		},
		[][]interface{}{
			[]interface{}{1, 2, 3, nil, 5, 4, nil},
			[]interface{}{1, 2, 3, nil, 5, nil, nil},
		},
	}

	results := []bool{
		true,
		true,
		false,
		false,
	}

	funcs := []func(*utils.TreeNode, *utils.TreeNode) bool{
		isSameTreeRecursion,
		isSameTreeIterationDFS,
		isSameTreeIterationBFS,
	}

	for i := range funcs {
		for k, v := range valsT {
			vals1 := v[0]
			t1 := utils.BinaryTreeConstructor(vals1)
			vals2 := v[1]
			t2 := utils.BinaryTreeConstructor(vals2)
			if results[k] != funcs[i](t1, t2) {
				t.Error("isSameTree error")
			}
		}
	}
}
