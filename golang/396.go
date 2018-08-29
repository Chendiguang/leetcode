package leetcode

import (
	"math"
)

/*
396. Rotate Function

Given an array of integers A and let n to be its length.
Assume Bk to be an array obtained by rotating the array A k positions clock-wise,
we define a "rotation function" F on A as follow:
F(k) = 0 * Bk[0] + 1 * Bk[1] + ... + (n-1) * Bk[n-1].
Calculate the maximum value of F(0), F(1), ..., F(n-1).

Note:
n is guaranteed to be less than 105.

Example:

A = [4, 3, 2, 6]

F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26

So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.

中文解析就是，相当于计算len(A)次，每次计算都是i*A[i]之和，
计算后，A都向右旋转了一位：
	[4, 3, 2, 6] --> [6, 4, 3, 2]
F(i) = sum(i*A[n-1-i]), n=len(A)
*/

// O(n^2), 超时，提交不能通过
func maxRotateFunction(A []int) int {
	if A == nil || len(A) == 0 {
		return 0
	}
	max := math.MinInt32
	n := len(A)
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += j * A[(i+j)%n]
		}
		max = Max(max, sum)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
通过找规律，解题
F(0) = 0A + 1B + 2C +3D
F(1) = 0D + 1A + 2B +3C
F(2) = 0C + 1D + 2A +3B
F(3) = 0B + 1C + 2D +3A
===>
F(1) = F(0) + sum - 4D
F(2) = F(1) + sum - 4C
F(3) = F(2) + sum - 4B

===> F(i) = F(i-1) + sum - n*A[n-i]
*/

func maxRotateFunction2(A []int) int {
	if A == nil || len(A) == 0 {
		return 0
	}
	tmp, sum, n := 0, 0, len(A)
	// 求数组A里面的元素之和，以及F[0]
	for i := 0; i < n; i++ {
		tmp += i * A[i]
		sum += A[i]
	}

	max := tmp
	for i := 1; i < n; i++ {
		tmp = tmp + sum - n*A[n-i]
		max = Max(max, tmp)
	}
	return max
}
