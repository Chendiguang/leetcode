package leetcode

/*
209. Minimum Size Subarray Sum
Given an array of n positive integers and a positive integer s,
find the minimal length of a contiguous subarray of which the sum ≥ s.
If there isn't one, return 0 instead.
*/

// two pointer, Timecomplexity: O(n)
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	min, tmpSum, cur := n+1, 0, 0
	for i, j := 0, 0; j < n; j++ {
		tmpSum += nums[j]
		for tmpSum >= s {
			cur = j - i + 1
			if cur < min {
				min = cur
			}
			tmpSum -= nums[i]
			i++
		}
	}
	if min == n+1 {
		return 0
	}
	return min
}
