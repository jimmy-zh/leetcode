package reverse_string

/*
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".
*/

func reverseString(s string) string {
	length := len(s)
	b := make([]byte, length)
	i, j := 0, length-1
	for i < length {
		b[j] = s[i]
		i++
		j--
	}
	return string(b)
}

func reverseStringPro(s string) string {
	length := len(s)
	b := make([]byte, length)
	i, j := 0, length-1
	for i <= j {
		b[i], b[j] = s[j], s[i]
		i++
		j--
	}
	return string(b)
}
