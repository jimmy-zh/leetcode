package valid_anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	ss := "nagaram"
	if !IsAnagram(s, ss) {
		t.Error("should be Anagram")
	}
	s = "rat"
	ss = "car"
	if IsAnagram(s, ss) {
		t.Error("should'nt be Anagram")
	}
	s = "rat"
	ss = "ratt"
	if IsAnagram(s, ss) {
		t.Error("should'nt be Anagram")
	}
}

func TestIsAnagramUnicode(t *testing.T) {
	s := "anagram"
	ss := "nagaram"
	if !IsAnagramUnicode(s, ss) {
		t.Error("should be Anagram")
	}
	s = "*张zz"
	ss = "张z*z"
	if !IsAnagramUnicode(s, ss) {
		t.Error("should be Anagram")
	}
	s = "*张(("
	ss = "*张((("
	if IsAnagramUnicode(s, ss) {
		t.Error("should'nt be Anagram")
	}
}
