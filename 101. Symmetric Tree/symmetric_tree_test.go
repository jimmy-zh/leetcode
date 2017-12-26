package symmetric_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestIsSymmetric(t *testing.T) {
	valsT := [][]interface{}{
		[]interface{}{},
		[]interface{}{1},
		[]interface{}{1, 2, 2, 4, 3, 3, 4},
		[]interface{}{1, 2, 3, 4, 3, 3, 4},
		[]interface{}{1, 2, 2, 4, 3, nil, 4},
	}

	results := []bool{
		true,
		true,
		true,
		false,
		false,
	}

	funcs := []func(*utils.TreeNode) bool{
		isSymmetricRecursion,
		isSymmetricIterationDFS,
		isSymmetricIterationBFS,
	}

	for i := range funcs {
		for k, v := range valsT {
			tree := utils.BinaryTreeConstructor(v)
			if results[k] != funcs[i](tree) {
				t.Error("isSymmetric error")
			}
		}
	}
}
