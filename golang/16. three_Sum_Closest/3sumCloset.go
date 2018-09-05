package three_Sum_Closest

import (
	"math"
	"sort"
)

// O(n^2)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	// distance means the distance between sum and target.
	// sum is the sum of three element.
	sum := nums[0] + nums[1] + nums[2]
	distance := math.MaxInt32

	// select an element, and count it with other two numbers.
	for i := 0; i < len(nums)-2; i++ {
		l, r := 0, len(nums)-1 // double pointer, left and right
		for l < r {
			// update sum while distance is closer.
			if nums[l]+nums[r]+nums[i]-target < distance {
				sum = nums[l] + nums[r] + nums[i]
				distance = abs(sum - target)
			}
			if nums[l]+nums[r]+nums[i] == target {
				return sum
			} else if nums[l]+nums[r]+nums[i] < target {
				l++
			} else {
				r--
			}
		}
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
