package nextPermutation

/*
	向后往前走，找到升序的转折点即nums[i + 1] > nums[i]
*/

func nextPermutation(nums []int) {
	numsLen := len(nums)
	var i = numsLen - 2
	for ; i >= 0; i-- {
		// 可以提前退出循坏，使得最好的匹配，能更快
		if nums[i+1] > nums[i] {
			break
		}
	}
	// 保证i的合法性
	if i >= 0 {
		j := numsLen - 1
		// 找到第一个比nums[i]大的点，交换
		for ; j >= 0; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	// 将nums[i+1]及之后的数字按照升序排列
	start, end := i+1, numsLen-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
