package leetcode

/*
1. could not contain 3, 4, 7
2. must contain one of [2, 5, 6, 9]
*/

// Solution 1
func rotatedDigits(N int) int {
	count := 0
	for j := 1; j <= N; j++ {
		containInvalid, containMust := false, false
		i := j
		for i > 0 {
			t := i % 10
			if t == 3 || t == 4 || t == 7 {
				containInvalid = true
				break
			}
			if t == 2 || t == 5 || t == 6 || t == 9 {
				containMust = true
			}
			i /= 10
		}
		if !containInvalid && containMust {
			count++
		}
	}
	return count
}

// Solution 2
func rotatedDigits2(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if valid(i) {
			count++
		}
	}
	return count
}

func valid(n int) bool {
	res := false
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			res = true
		case 3, 4, 7:
			return false
		}
		n /= 10
	}
	return res
}

// Solution 1 的改进版
func rotatedDigits3(N int) int {
	count := 0
	for j := 1; j <= N; j++ {
		i, containInvalid, containMust := j, false, false
		for i > 0 {
			switch i % 10 {
			case 3, 4, 7:
				containInvalid = true
				break
			case 2, 5, 6, 9:
				containMust = true
			}
			i /= 10
		}
		if !containInvalid && containMust {
			count++
		}
	}
	return count
}
