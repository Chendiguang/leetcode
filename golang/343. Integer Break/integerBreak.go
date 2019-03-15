/*
Given a positive integer n, break it into the sum of at least two positive
integers and maximize the product of those integers.
Return the maximum product you can get.

Example 1:

Input: 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
Example 2:

Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
*/
// 数学解法
package integerBreak

import "math"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	c := 1
	for n > 4 {
		c *= 3
		n = n - 3
	}
	return c * n
}

func integerBreak2(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	} else if n == 4 {
		return n
	}
	n -= 5
	return int(math.Pow(3, float64(n/3+1))) * (n%3 + 2)
}
