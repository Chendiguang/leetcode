package leetcode

import (
	"math"
)

/*
204. Count Primes
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10,
	they are 2, 3, 5, 7.
*/

// 埃拉托斯特尼筛法
func countPrimes(n int) int {
	sum, maps := 0, make([]bool, n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if !maps[i] {
			for j := i * i; j < n; j += i {
				maps[j] = true
			}
		}
	}

	for i := 2; i < n; i++ {
		if !maps[i] {
			sum++
		}
	}
	return sum
}
