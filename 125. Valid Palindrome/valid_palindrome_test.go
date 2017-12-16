package valid_palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "ab3d3ba"
	if !isPalindrome(s) {
		t.Error("should be Palindrome")
	}
	s = "aowao"
	if isPalindrome(s) {
		t.Error("should not be Palindrome")
	}
	s = "a𐅭|b张A"
	if !isPalindrome(s) {
		t.Error("should be Palindrome")
	}
	s = "A man, a plan, a canal: Panama"
	if !isPalindrome(s) {
		t.Error("should be Palindrome")
	}
}
