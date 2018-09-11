package threeSum

import "sort"

/*
Given an array nums of n integers, are there elements a, b, c in nums
such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.
*/

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			if 3*nums[i] > 0 || 3*nums[n-1] < 0 {
				break
			}
			l, r := i+1, n-1
			for l < r {
				tmp := -nums[i]
				if 2*nums[l] > tmp || 2*nums[r] < tmp {
					break
				}
				if nums[l]+nums[r] < tmp {
					l++
				} else if nums[l]+nums[r] > tmp {
					r--
				} else {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return res
}
