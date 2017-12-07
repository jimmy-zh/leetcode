package find_the_duplicate_number

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	nums := [][]int{
		[]int{},
		[]int{1},
		[]int{4, 3, 1, 2, 4},
		[]int{4, 4, 1, 3, 2},
		[]int{1, 3, 2, 4, 4},
		[]int{4, 4, 4, 4, 4},
		[]int{1, 2, 2, 3},
	}

	result := []int{
		0,
		0,
		4,
		4,
		4,
		4,
		2,
	}

	for k := range nums {
		if findDuplicateCycle(nums[k]) != result[k] {
			t.Error("find duplicate error")
		}
	}
	for k := range nums {
		if findDuplicateBinary(nums[k]) != result[k] {
			t.Error("find duplicate error")
		}
	}
}
