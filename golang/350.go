package leetcode

import (
	"sort"
)

/*
Given two arrays, write a function to compute their intersection.
Each element in the result should appear as many times as
it shows in both arrays.
*/
// Similar with 349.Intersection of Two Arrays

// Time Complexity: O(n), Space Complexity: O(n)
func intersect(nums1 []int, nums2 []int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		maps[nums1[i]]++ // Calculate the number of occurrences
	}
	res := []int{}
	for j := 0; j < len(nums2); j++ {
		if v, ok := maps[nums2[j]]; ok && v > 0 {
			maps[nums2[j]]-- // need to subtract one
			res = append(res, nums2[j])
		}
	}
	return res
}

//!+ Solution 2
// Time Complexity: O(nlgn), Space Complexity: O(n)
// 先给数组排序，然后用两个指针i,j分别指向两个数组的起始位置，如果两个指针指向的
// 数字相等则存到数组里，指针都自增1，如果nums1[i] > nums2[j]，则指针j++, 否则i++
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j, res := 0, 0, []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
