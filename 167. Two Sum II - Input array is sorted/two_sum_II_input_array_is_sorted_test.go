package two_sum_II_input_array_is_sorted

import (
	"testing"
)

func TestTwoSumTwoPointer(t *testing.T) {
	numbers := [][]int{
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11},
		[]int{2, 7, 11},
		[]int{1},
		[]int{},
		[]int{-7, -2, 1, 3, 7},
	}
	target := []int{
		9,
		26,
		17,
		25,
		1,
		1,
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
		[]int{},
		[]int{2, 4},
	}

	for i := range numbers {
		arr := twoSumTwoPointer(numbers[i], target[i])
		if len(arr) != len(result[i]) {
			t.Error("twoSum error")
		}
		for j := range arr {
			if arr[j] != result[i][j] {
				t.Error("twoSum error")
			}
		}
	}

}

func TestTwoSumBinarySearch(t *testing.T) {
	numbers := [][]int{
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11, 15},
		[]int{2, 7, 11},
		[]int{2, 7, 11},
		[]int{1},
		[]int{},
		[]int{-7, -2, 1, 3, 7},
	}
	target := []int{
		9,
		26,
		17,
		25,
		1,
		1,
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
		[]int{},
		[]int{2, 4},
	}

	for i := range numbers {
		arr := twoSumBinarySearch(numbers[i], target[i])
		if len(arr) != len(result[i]) {
			t.Error("twoSum error")
		}
		for j := range arr {
			if arr[j] != result[i][j] {
				t.Error("twoSum error")
			}
		}
	}

}
