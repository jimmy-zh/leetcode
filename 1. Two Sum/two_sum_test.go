package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	vals := [][]int{
		[]int{},
		[]int{1},
		[]int{1, 2},
		[]int{3, 6, 7, 2, 1},
	}

	targets := []int{
		1,
		1,
		3,
		7,
	}
	result := [][]int{
		nil,
		nil,
		[]int{0, 1},
		[]int{1, 4},
	}

	for i := range result {
		res := twoSum(vals[i], targets[i])
		if result[i] == nil && res != nil {
			t.Error("twoSum error")
			continue
		}
		if len(res) != len(result[i]) {
			t.Error("twoSum error")
			continue
		}
		for k := range result[i] {
			if result[i][k] != res[k] {
				t.Error("twoSum error")
			}
		}
	}
}
