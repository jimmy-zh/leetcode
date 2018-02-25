package two_sum

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	m := make(map[int]int)
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = k
	}
	return nil
}
