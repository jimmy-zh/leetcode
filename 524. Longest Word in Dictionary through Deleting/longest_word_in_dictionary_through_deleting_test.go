package longest_word_in_dictionary_through_deleting

import (
	"testing"
)

func TestFindLongestWordNoSort(t *testing.T) {
	s := "bgaewdf"
	d := []string{"ga", "ge", "abw", "awf", "ewf", "bgaewdfa"}
	if findLongestWordNoSort(s, d) != "awf" {
		t.Error("error result")
	}
}

func TestFindLongestWordSort(t *testing.T) {
	s := "bgaewdf"
	d := []string{"ga", "ge", "abw", "awf", "ewf", "bgaewdfa"}
	if findLongestWordSort(s, d) != "awf" {
		t.Error("error result")
	}
}
