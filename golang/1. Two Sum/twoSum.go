package twoSum

func twoSum(nums []int, target int) []int {
	var res []int
	if nums == nil || len(nums) <= 1 {
		return res
	}

	unorderMap := make(map[int]int)
	for i, item := range nums {
		if index, ok := unorderMap[item]; ok {
			return []int{index, i}
		} else {
			unorderMap[target-item] = i
		}
	}
	return res
}
