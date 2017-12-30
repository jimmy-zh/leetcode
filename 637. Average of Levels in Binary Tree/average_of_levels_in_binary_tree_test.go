package average_of_levels_in_binary_tree

import (
	"testing"

	"github.com/midnight-vivian/go-data-structures/utils"
)

func TestAverageOfLevels(t *testing.T) {
	vals := []interface{}{8, 7, 4, 5, 6, 9, nil, nil, nil, nil, nil, nil, 3}
	tVal := utils.BinaryTreeConstructor(vals)
	result := []float64{float64(8) / 1, float64(11) / 2, float64(20) / 3, float64(3) / 1}
	myResult := averageOfLevels(tVal)
	for i := range result {
		if result[i] != myResult[i] {
			t.Error("AverageOfLevels error")
		}
	}

	vals = []interface{}{}
	tVal = utils.BinaryTreeConstructor(vals)
	myResult = averageOfLevels(tVal)
	if len(myResult) != 0 {
		t.Error("AverageOfLevels error")
	}

	vals = []interface{}{8}
	tVal = utils.BinaryTreeConstructor(vals)
	result = []float64{float64(8) / 1}
	myResult = averageOfLevels(tVal)
	for i := range result {
		if result[i] != myResult[i] {
			t.Error("AverageOfLevels error")
		}
	}
}
