package valid_palindrome

import (
	"strings"
)

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.

Note:
Have you consider that the string might be empty? This is a good question to ask during an interview.
For the purpose of this problem, we define empty string as valid palindrome.
*/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !(s[i]-'a' >= 0 && s[i]-'a' <= 25) && !(s[i]-'0' >= 0 && s[i]-'0' <= 9) {
			i++
			continue
		}
		if !(s[j]-'a' >= 0 && s[j]-'a' <= 25) && !(s[j]-'0' >= 0 && s[j]-'0' <= 9) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}
