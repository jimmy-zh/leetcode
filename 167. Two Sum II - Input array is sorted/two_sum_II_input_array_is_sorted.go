package two_sum_II_input_array_is_sorted

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2
*/

func twoSumBinarySearch(numbers []int, target int) []int {
	for i := range numbers {
		complement := target - numbers[i]
		l, r := i+1, len(numbers)-1
		for l <= r {
			m := l + (r-l)/2
			switch {
			case complement == numbers[m]:
				return []int{i + 1, m + 1}
			case complement > numbers[m]:
				l = m + 1
			case complement < numbers[m]:
				r = m - 1
			}
		}
	}
	return []int{}
}

func twoSumTwoPointer(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
