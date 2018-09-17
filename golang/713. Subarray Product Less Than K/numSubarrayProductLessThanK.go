package Sub

import "fmt"

/*
Your are given an array of positive integers nums.

Count and print the number of (contiguous) subarrays where the
product of all the elements in the subarray is less than k.
*/

// Slidng window
// 维护一个滑动窗口，保证窗口内元素的值小于k
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	left, res, prod := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		// 窗口中元素个数==当前窗口去除重复后的组合数
		fmt.Println(nums[left : right+1])
		res += right - left + 1
	}
	return res
}
