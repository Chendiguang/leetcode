package Container_With_Most_Water

/*
Given n non-negative integers a1, a2, ..., an ,
where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

func maxArea(height []int) int {
	max, i, j := 0, 0, len(height)-1
	for i < j {
		t := (j - i) * Min(height[i], height[j])
		max = Max(max, t)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
