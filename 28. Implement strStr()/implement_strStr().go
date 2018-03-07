package implement_strStr

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := range needle {
			if needle[j] != haystack[i+j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
