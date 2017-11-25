package longest_word_in_dictionary_through_deleting

/*
Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string.
If there are more than one possible results, return the longest word with the smallest lexicographical order.
If there is no possible result, return the empty string.

Example 1:
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
Example 2:
Input:
s = "abpcplea", d = ["a","b","c"]

Output:
"a"
Note:
All the strings in the input will only contain lower-case letters.
The size of the dictionary won't exceed 1,000.
The length of all the strings in the input won't exceed 1,000.
*/

func findLongestWordNoSort(s string, d []string) string {
	var res string
	for _, di := range d {
		i, j := 0, 0
		for i < len(s) && j < len(di) {
			if s[i] == di[j] {
				j++
			}
			i++
		}
		if j == len(di) && (len(di) > len(res) || (len(di) == len(res) && di < res)) {
			res = di
		}
	}
	return res
}

func findLongestWordSort(s string, d []string) string {
	for i := 1; i < len(d); i++ {
		for j := i; j > 0; j-- {
			if len(d[j]) > len(d[j-1]) || (len(d[j]) == len(d[j-1]) && d[j] < d[j-1]) {
				d[j], d[j-1] = d[j-1], d[j]
			}
		}
	}
	for _, di := range d {
		i, j := 0, 0
		for i < len(s) && j < len(di) {
			if s[i] == di[j] {
				j++
			}
			i++
		}
		if j == len(di) {
			return di
		}
	}
	return ""
}
