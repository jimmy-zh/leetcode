package find_the_duplicate_number

/*
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Note:
You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
Credits:
Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
*/

//runtime complexity:O(n)
func findDuplicateCycle(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	one, two := 0, 0
	for {
		one = nums[one]
		two = nums[nums[two]]
		if one == two {
			break
		}
	}
	one = 0
	for one != two {
		one = nums[one]
		two = nums[two]
	}
	return one

}

//runtime complexity:O(nlog(n))
func findDuplicateBinary(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	lo, hi := 1, len(nums)-1
	for lo < hi {
		m := (lo + hi) / 2
		count := 0
		for i := range nums {
			if nums[i] <= m {
				count++
			}
		}
		if count <= m {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return lo
}

//this wrong solution(binary partition) will lead infinite loop
func findDuplicateBinary2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	lo, hi := 1, len(nums)-1
	for lo < hi {
		m := (lo + hi) / 2
		count := 0
		for i := range nums {
			if nums[i] < m {
				count++
			}
		}
		if count < m {
			lo = m
		} else {
			hi = m - 1
		}
	}
	return lo
}
