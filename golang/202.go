package leetcode

/*
Input: 19
Output: true
Explanation:
	1^2 + 9^2 = 82
	8^2 + 2^2 = 68
	6^2 + 8^2 = 100
	1^2 + 0^2 + 0^2 = 1
*/

func isHappy(n int) bool {
	maps := make(map[int]bool)
	for {
		if maps[n] {
			break
		}
		maps[n] = true
		var tmp int
		for n > 0 {
			t := n % 10
			tmp += t * t
			n /= 10
		}
		if tmp == 1 {
			return true
		}
		n = tmp
	}
	return false
}
