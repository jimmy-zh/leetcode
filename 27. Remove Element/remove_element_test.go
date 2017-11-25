package remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tableDriver := []map[string]interface{}{
		map[string]interface{}{
			"nums":   []int{2, 3, 4, 3, 5, 3, 6},
			"val":    3,
			"length": 4,
		},
		map[string]interface{}{
			"nums":   []int{},
			"val":    3,
			"length": 0,
		},
		map[string]interface{}{
			"nums":   []int{3, 3, 3},
			"val":    3,
			"length": 0,
		},
		map[string]interface{}{
			"nums":   []int{1, 2, 4},
			"val":    3,
			"length": 3,
		},
	}
	for _, item := range tableDriver {
		nums := item["nums"].([]int)
		val := item["val"].(int)
		length := item["length"].(int)
		if removeElement(nums, val) != length {
			t.Error("Length is wrong")
		}
		for i := 0; i < length; i++ {
			if nums[i] == val {
				t.Error("Value is wrong")
			}
		}
	}
}

func TestRemoveElementEasy(t *testing.T) {
	tableDriver := []map[string]interface{}{
		map[string]interface{}{
			"nums":   []int{2, 3, 4, 3, 5, 3, 6},
			"val":    3,
			"length": 4,
		},
		map[string]interface{}{
			"nums":   []int{},
			"val":    3,
			"length": 0,
		},
		map[string]interface{}{
			"nums":   []int{3, 3, 3},
			"val":    3,
			"length": 0,
		},
		map[string]interface{}{
			"nums":   []int{1, 2, 4},
			"val":    3,
			"length": 3,
		},
	}
	for _, item := range tableDriver {
		nums := item["nums"].([]int)
		val := item["val"].(int)
		length := item["length"].(int)
		if removeElementEasy(nums, val) != length {
			t.Error("Length is wrong")
		}
		for i := 0; i < length; i++ {
			if nums[i] == val {
				t.Error("Value is wrong")
			}
		}
	}
}
