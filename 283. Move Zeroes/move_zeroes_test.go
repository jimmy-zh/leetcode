package move_zeroes

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	input := []int{1, 2, 0, 3, 0, 4, 5}
	output := []int{1, 2, 3, 4, 5, 0, 0}
	moveZeroes(input)
	for i := range input {
		if input[i] != output[i] {
			t.Error("move error")
		}
	}

}
