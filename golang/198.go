package leetcode

/*
198. House Robber
dp 问题, 本质是在一列数组中取出一个或多个不相邻数，使其和最大。
*/

// solution 1, 维护两个变量
func rob(nums []int) int {
	preMax, curMax := 0, 0
	for i := 0; i < len(nums); i++ {
		tmp := curMax
		if preMax+nums[i] > curMax {
			curMax = preMax + nums[i]
		}
		preMax = tmp
	}
	return curMax
}

// solution 2, 维护一个数组
// 状态转移方程：dp[i] = max(dp[i - 1], dp[i - 2] + num[i - 1])
// dp[i]表示打劫到第i间房屋时累计取得的金钱最大值
func rob2(nums []int) int {
	l := len(nums)
	if l < 1 {
		return 0
	}
	dp := make([]int, l+1)
	dp[1] = nums[0]
	for i := 2; i < l+1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[l]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// same as 1, 2
func rob3(nums []int) int {
	even, odd := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			odd = max(odd+nums[i], even)
		} else {
			even = max(even+nums[i], odd)
		}
	}
	return max(even, odd)
}
