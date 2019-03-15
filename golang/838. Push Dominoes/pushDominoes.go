package pushDominoes

/*
Example 1:

Input: ".L.R...LR..L.."
Output: "LL.RR.LLRRLL.."
Example 2:

Input: "RR.L"
Output: "RR.L"
Explanation: The first domino expends no additional force on the second domino.
Note:

0 <= N <= 10^5
String dominoes contains only 'L', 'R' and '.'
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// calculate force, two pass, 用物理学上的守恒定律
func pushDominoes(dominoes string) string {
	force, n := 0, len(dominoes)
	forces := make([]int, n)

	// populate forces from left to right
	for i := 0; i < n; i++ {
		if dominoes[i] == 'R' {
			force = n
		} else if dominoes[i] == 'L' {
			force = 0
		} else {
			force = max(force-1, 0)
		}
		forces[i] = force
	}

	force = 0
	// populate forces from right to left
	for i := n - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			force = n
		} else if dominoes[i] == 'R' {
			force = 0
		} else {
			force = max(force-1, 0)
		}
		// merge
		forces[i] -= force
	}

	// >0 means 'R', =0 means '.', <0 means 'L'
	ans := make([]rune, n)
	for i := 0; i < n; i++ {
		if forces[i] > 0 {
			ans[i] = 'R'
		} else if forces[i] < 0 {
			ans[i] = 'L'
		} else {
			ans[i] = '.'
		}
	}
	return string(ans)
}

// Adjacent Symbols, one pass
func pushDominoesOnePass(dominoes string) string {
	d := "L" + dominoes + "R"
	n := len(d)
	ans := make([]byte, 0, n-2)
	// i, 用于指示当前平衡块的左边界
	for i, j := 0, 1; j < n; j++ {
		if dominoes[j] == '.' {
			continue
		}
		if i != 0 {
			ans = append(ans, d[i]) // 添加符合平衡的边界元素
		}
		m := j - i - 1 // 求d[i:j]中间的元素个数
		// 主要分为4种情况
		if d[i] == d[j] {
			// L...L or R...R, 需要填充中间部分的数量，值为d[i]
			for k := 0; k < m; k++ {
				ans = append(ans, d[i])
			}
		} else if d[i] == 'L' && d[j] == 'R' {
			// L....R
			for k := 0; k < m; k++ {
				ans = append(ans, '.')
			}
		} else {
			// LLLRRR or LLL.RRR
			for k := 0; k < m/2; k++ {
				ans = append(ans, 'R')
			}
			if m%2 == 1 {
				ans = append(ans, '.')
			}
			for k := 0; k < m/2; k++ {
				ans = append(ans, 'L')
			}
		}
		i = j
	}
	return string(ans)
}
