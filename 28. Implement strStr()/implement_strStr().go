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
	if len(needle) == 0 {
		return 0
	}
	i, begin := 0, -1
	for j := 0; j < len(haystack); j++ {
		if haystack[j] == needle[i] {
			if begin == -1 {
				begin = j
			}
			i++
		}
		if i == len(needle) {
			if j-begin == len(needle)-1 {
				return begin
			}
			j, i, begin = begin, 0, -1
		}
	}
	return -1
}
