package longest_word_in_dictionary_through_deleting

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
