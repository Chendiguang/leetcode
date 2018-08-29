package leetcode

/*
分三步，首先整个数组反转，然后以k为分界线，对分割的两部分进行反转
Time Complexity: O(n), Space Complexity: O(1)
*/

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
