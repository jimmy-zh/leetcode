package two_sum_II_input_array_is_sorted

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := [][]int{
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11},
		[]int{2, 7, 11},
		[]int{},
	}
	target := []int{
		9,
		26,
		17,
		25,
		1,
		1,
	}
	result := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{1, 4},
		[]int{},
		[]int{},
		[]int{},
	}

	for i := range numbers {
		arr := twoSum(numbers[i], target[i])
		for j := range arr {
			if arr[j] != result[i][j] {
				t.Error("twoSum error")
			}
		}
	}

}
