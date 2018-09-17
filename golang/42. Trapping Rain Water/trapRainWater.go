package Trap

/*
Given n non-negative integers representing an elevation
map where the width of each bar is 1,
compute how much water it is able to trap after raining.
*/

func trap(height []int) int {
	height = append(height, 0)
	l, r := 0, 1
	res := 0
	for r < len(height)-1 {
		if height[r] > height[r+1] {
			m := min(height[r], height[l])
			for i := l; i < r; i++ {
				res += m - height[i]
			}
			l = r
			r++
		}
	}
	return res
}

// dynamic programing
// 通过存储已经计算过的数据，避免重复计算
func trap2(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	res, size := 0, len(height)
	left_max, right_max := make([]int, size), make([]int, size)
	left_max[0] = height[0]
	for i := 1; i < size; i++ {
		left_max[i] = max(left_max[i-1], height[i])
	}

	right_max[size-1] = height[size-1]
	for j := size - 2; j >= 0; j-- {
		right_max[j] = max(right_max[j+1], height[j])
	}

	for i := 1; i < size-1; i++ {
		res += min(left_max[i], right_max[i]) - height[i]
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Using stacks
func trap3(height []int) int {
	ans, current := 0, 0
	stack := []int{}
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			l := len(stack)
			top := stack[l-1]
			stack = stack[:l-1]
			if len(stack) == 0 {
				break
			}
			idx := stack[len(stack)-1]
			distance := current - idx - 1
			bounded_height := min(height[current], height[idx]) - height[top]
			ans += distance * bounded_height
		}
		stack = append(stack, current)
		current++
	}
	return ans
}
