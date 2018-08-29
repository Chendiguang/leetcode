package leetcode

/*
153. Find Minimum in Rotated Sorted Array
You may assume no duplicate exists in the array.

Example 1:
Input: [3,4,5,1,2]
Output: 1

Example 2:
Input: [4,5,6,7,0,1,2]
Output: 0
*/

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	for nums[l] > nums[r] {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func findMin2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low < high-1 {
		m := (low + high) / 2
		if nums[m] < nums[high] {
			high = m
		} else if nums[m] > nums[high] {
			low = m
		}
	}
	if nums[low] < nums[high] {
		return nums[low]
	}
	return nums[high]
}

func findMin3(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
