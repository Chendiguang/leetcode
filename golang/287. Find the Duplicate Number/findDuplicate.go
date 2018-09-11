package findDuplicate

/*
Given an array nums containing n + 1 integers where each integer
is between 1 and n (inclusive),
prove that at least one duplicate number must exist.
Assume that there is only one duplicate number,
find the duplicate one

Note:
You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
*/

// 二分查找， O(nlgn)
// 求出中点m= low + (high-low) * 0.5, 然后遍历整个数组，统计所有小于或等于m
// 的个数n, 如果n>m, 重复值在[m+1, high], 否则在[low, m-1]
func findDuplicate2(nums []int) int {
	cnt, mid := 0, 0
	low, high := 1, len(nums)-1
	for low < high {
		mid = (high + low) / 2
		cnt = 0

		for _, a := range nums {
			if a <= mid {
				cnt++
			}
		}
		if cnt < mid {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func findDuplicate(nums []int) int {
	p1, p2 := nums[0], nums[nums[0]]

	for p1 != p2 {
		p1, p2 = nums[p1], nums[nums[p2]]
	}

	p2 = 0
	for p1 != p2 {
		p1, p2 = nums[p1], nums[p2]
	}
	return p1
}
