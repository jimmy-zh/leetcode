package remove_duplicates_from_sorted_array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := [][]int{
		[]int{1, 2, 2, 3, 4, 4, 4, 5},
		[]int{1, 2, 2},
		[]int{1, 1, 2},
		[]int{1, 1, 1},
		[]int{1},
		[]int{},
	}
	result := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2},
		[]int{1, 2},
		[]int{1},
		[]int{1},
		[]int{},
	}
	for k := range nums {
		if removeDuplicates(nums[k]) != len(result[k]) {
			t.Error("remove error")
		}
		for i := 0; i < len(result[k]); i++ {
			if nums[k][i] != result[k][i] {
				t.Error("remove error")
			}
		}
	}

}
