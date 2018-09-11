package twoSum2

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l <= r {
		tmp := numbers[l] + numbers[r]
		if tmp == target {
			return []int{l + 1, r + 1}
		} else if tmp < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
