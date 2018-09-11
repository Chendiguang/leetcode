package findPairs

/*
Given an array of integers and an integer k,
you need to find the number of unique k-diff pairs in the array.
Here a k-diff pair is defined as an integer pair (i, j),
where i and j are both numbers in the array
and their absolute difference is k.
*/

func findPairs(nums []int, k int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	res := 0
	for key, v := range mp {
		if k == 0 && v > 1 {
			res++
		}
		if _, ok := mp[key+k]; ok && k > 0 {
			res++
		}
	}
	return res
}

// more quickly
func findPairs2(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	res := 0
	for n, v := range mp {
		if k != 0 && mp[k+n] > 0 {
			res++
		} else if k == 0 && v > 1 {
			res++
		}
	}
	return res
}
