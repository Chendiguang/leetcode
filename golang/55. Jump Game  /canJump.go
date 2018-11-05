package canJump

// DP
func canJump(nums []int) bool {
	// dp数组用于保存，到达当前点i剩余的步数
	l := len(nums)
	dp := make([]int, l)
	// dp[i] = max(dp[i-1], nums[i-1]) - 1
	for i := 1; i < l; i++ {
		if dp[i-1] > nums[i-1] {
			dp[i] = dp[i-1] - 1
		} else {
			dp[i] = nums[i-1] - 1
		}
		if dp[i] < 0 {
			return false
		}
	}
	return dp[l-1] >= 0
}

// Solution 2, 贪婪算法Greedy Algorithm
func canJump2(nums []int) bool {
	// 假定能到达终点，从后向前遍历
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
