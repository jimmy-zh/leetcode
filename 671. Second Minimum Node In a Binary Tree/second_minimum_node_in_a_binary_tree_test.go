package second_minimum_node_in_a_binary_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestFindSecondMinimumValue(t *testing.T) {
	vals := [][]interface{}{
		[]interface{}{},
		[]interface{}{
			1,
		},
		[]interface{}{
			2, 2, 5, 2, 4, 5, 6, 2, 3, nil, nil, nil, nil, 6, 8,
		},
		[]interface{}{
			2, 2, 2,
		},
	}
	result := []int{
		-1,
		-1,
		3,
		-1,
	}
	for i := range vals {
		if findSecondMinimumValue(utils.BinaryTreeConstructor(vals[i])) != result[i] {
			t.Error("findSecondMinimumValue error")
		}
	}

}
