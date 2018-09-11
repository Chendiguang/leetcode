package removeDuplicates

// two pointers
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 3 {
		return l
	}
	// 指向改动后，符合条件的末尾位置
	res := 2
	for i := 2; i < l; i++ {
		if nums[res-2] != nums[i] || nums[res-1] != nums[i] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
