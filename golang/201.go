package leetcode

/*
Given a range [m, n] where 0 <= m <= n <= 2147483647,
return the bitwise AND of all numbers in this range, inclusive.
分析可知，这道题是求他们二进制数的公共前缀。
解法：
	直接右移m、n，直到 m==n 并且记录右移的次数count，最后将此时的m执行左移操作
	m << count即可得。
*/

func rangeBitwiseAnd(m int, n int) int {
	count := uint32(0)
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
