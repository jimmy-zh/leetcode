package reverse_string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "abcabc"
	if reverseString(s) != "cbacba" {
		t.Error("reverse error")
	}

	s1 := "a"
	if reverseString(s1) != "a" {
		t.Error("reverse error")
	}

	s2 := ""
	if reverseString(s2) != "" {
		t.Error("reverse error")
	}
}

func TestReverseStringPro(t *testing.T) {
	s := "abcabc"
	if reverseStringPro(s) != "cbacba" {
		t.Error("reverse error")
	}

	s1 := "a"
	if reverseStringPro(s1) != "a" {
		t.Error("reverse error")
	}

	s2 := ""
	if reverseStringPro(s2) != "" {
		t.Error("reverse error")
	}
}
