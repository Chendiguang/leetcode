package leetcode

/*
Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
*/

// Test: 4ms
func intersection(nums1 []int, nums2 []int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		maps[nums1[i]] = 1
	}
	res := []int{}
	for j := 0; j < len(nums2); j++ {
		if v, ok := maps[nums2[j]]; ok && v == 1 {
			maps[nums2[j]] += 1
			res = append(res, nums2[j])
		}
	}
	return res
}
