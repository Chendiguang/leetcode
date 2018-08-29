package leetcode

/*
665. Non-decreasing Array
Given an array with n integers, your task is to check if it could become
non-decreasing by modifying at most 1 element.
We define an array is non-decreasing if array[i] <= array[i + 1] holds
for every i (1 <= i < n).

Example 1:
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
*/
func checkPossibility(nums []int) bool {
	// cnt统计修改的次数，修改次数超出1次时则为假
	cnt := 0
	for i := 1; i < len(nums) && cnt <= 1; i++ {
		if nums[i-1] > nums[i] {
			cnt++
			// 考虑把前面的数字变成当前的数字的情况，
			// 当前面的数字再前面没有数字，那么无疑改前面的数字是最好的，不会影响后面。
			// 如果前面的数字再前面还有数字，并且要是小于关系，那么改前面这个数字也是对后面没影响的
			if i < 2 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return cnt <= 1
}
