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
	if needle == "" {
		return 0
	}
	i, j, imatch := 0, 0, -1
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			if j == 0 {
				imatch = i
			}
			i, j = i+1, j+1
			if j == len(needle) {
				return imatch
			}
		} else if imatch != -1 {
			i, imatch, j = imatch+1, -1, 0
		} else {
			i++
		}
	}
	return -1
}
