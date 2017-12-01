package reverse_vowels_of_a_string

/*
Write a function that takes a string as input and reverse only the vowels of a string.
Example 1:
Given s = "hello", return "holle".
Example 2:
Given s = "leetcode", return "leotcede".
Note:
The vowels does not include the letter "y".
*/

func reverseVowels(s string) string {
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	b := []byte(s)
	i, j := -1, len(b)
	for {
		for {
			i++
			if i == len(b) || m[b[i]] {
				break
			}
		}
		for {
			j--
			if j < 0 || m[b[j]] {
				break
			}
		}
		if i >= j {
			break
		}
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
