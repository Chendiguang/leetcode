/*
	back tracking algorithm
	Testing's time is 4ms
*/

package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	backTracking(nums, []int{}, 0, &res)
	return res
}

func backTracking(nums, tmp []int, pos int, res *[][]int) {
	// pos is current index
	if pos == len(nums) {
		tp := make([]int, len(tmp))
		copy(tp, tmp)
		*res = append(*res, tp)
		return
	}

	// important
	backTracking(nums, append(tmp, nums[pos]), pos+1, res)
	backTracking(nums, tmp, pos+1, res)
}
