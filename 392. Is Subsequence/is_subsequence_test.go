package is_subsequence

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	input := [][]string{
		[]string{
			"",
			"",
		},
		[]string{
			"",
			"o",
		},
		[]string{
			"o",
			"",
		},
		[]string{
			"abc",
			"ahbgdc",
		},
		[]string{
			"abe",
			"ahbgdc",
		},
		[]string{
			"abh",
			"ahbgdc",
		},
	}
	result := []bool{
		true,
		true,
		false,
		true,
		false,
		false,
	}
	for i := range result {
		if isSubsequence(input[i][0], input[i][1]) != result[i] {
			t.Error("IsSubsequence error")
		}
	}
}
