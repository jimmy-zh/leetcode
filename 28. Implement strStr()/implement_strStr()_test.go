package implement_strStr

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	haystacks := []string{"hello", "vivian", "vivian", "hello", "hel", "", "hello", "mississippi", "mississippi"}
	needle := []string{"ll", "vi", "an", "elo", "hello", "hello", "", "issip", "pi"}
	result := []int{2, 0, 4, -1, -1, -1, 0, 4, 9}
	for i := range haystacks {
		if strStr(haystacks[i], needle[i]) != result[i] {
			t.Error("strstr error")
		}
	}
}
