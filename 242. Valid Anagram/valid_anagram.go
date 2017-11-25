package valid_anagram

/*
Given two strings s and t, write a function to determine if t is an anagram of s.
For example,
s = "anagram", t = "nagaram", return true.
s = "rat", t = "car", return false.

Note:
You may assume the string contains only lowercase alphabets.
Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var char26 [26]int
	for i := range s {
		char26[s[i]-'a']++
		char26[t[i]-'a']--
	}
	for i := range char26 {
		if char26[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagramUnicode(s string, t string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
