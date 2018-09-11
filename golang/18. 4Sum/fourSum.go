package fourSum

import "sort"

/*
Example:
Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

// 12ms
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	res := make([][]int, 0)

	for i := 0; i < l-3; i++ {
		for j := i + 1; j < l-2; j++ {
			l, r := j+1, l-1
			tmp := target - nums[i] - nums[j]
			for l < r {
				if nums[l]+nums[r] < tmp {
					l++
				} else if nums[l]+nums[r] > tmp {
					r--
				} else {
					// list := make([]int, 0, 4)
					// list = append(list, nums[i])
					// list = append(list, nums[j])
					// list = append(list, nums[l])
					// list = append(list, nums[r])
					list := []int{nums[i], nums[j], nums[l], nums[r]}
					res = append(res, list)
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l > r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				}
			}

			for j < l-2 && nums[j] == nums[j+1] {
				j++
			}
		}
		for i < l-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// 4ms
func fourSum2(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-1; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			if 4*nums[i] > target || 4*nums[n-1] < target {
				break
			}
			target2 := target - nums[i]
			for j := i + 1; j < n-2; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					// 这些判断有助于预先跳出无用功操作
					if 3*nums[j] > target2 || 3*nums[n-1] < target2 {
						break
					}
					l, r := j+1, n-1
					tgt := target - nums[i] - nums[j]
					for l < r {
						if 2*nums[l] > tgt || 2*nums[r] < tgt {
							break
						}
						if nums[l]+nums[r] > tgt {
							r--
						} else if nums[l]+nums[r] < tgt {
							l++
						} else {
							res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
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
		}
	}
	return res
}
