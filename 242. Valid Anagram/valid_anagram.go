package valid_anagram

//Note:
//You may assume the string contains only lowercase alphabets.
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

//Follow up:
//What if the inputs contain unicode characters? How would you adapt your solution to such case?
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
