package maxSubArray

func maxSubArray(nums []int) int {
	// 使用两个变量来更新保存最大的和
	maxSum, curSum := nums[0], nums[0]
	for _, n := range nums[1:] {
		if curSum <= 0 {
			curSum = n
		} else {
			curSum += n
		}

		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum

}
