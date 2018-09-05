package leetcode

/*
838. Push Dominoes
*/

// 计算压力, two pass
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
			force = Max(force-1, 0)
		}
		forces[i] = force
	}

	// populate forces going from right to left
	force = 0
	for i := n - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			force = n
		} else if dominoes[i] == 'R' {
			force = 0
		} else {
			force = Max(force-1, 0)
		}
		forces[i] -= force
	}
	// >0 menas 'R', <0 means 'L', =0 is '.'
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

// two pointers, on pass, more quickly.
func pushDominoes2(dominoes string) string {
	d := "L" + dominoes + "R"
	n := len(d)
	ans := make([]byte, 0, n-2) // 长度为0，容量为len(dominoes)
	for i, j := 0, 1; j < n; j++ {
		if d[j] == '.' {
			continue
		}
		if i != 0 {
			ans = append(ans, d[i])
		}
		m := j - i - 1
		if d[i] == d[j] {
			for k := 0; k < m; k++ {
				ans = append(ans, d[i])
			}
		} else if d[i] == 'L' && d[j] == 'R' {
			for k := 0; k < m; k++ {
				ans = append(ans, '.')
			}
		} else {
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

// two pointer, golangnic
func pushDominoes3(dominoes string) string {
	d := "L" + dominoes + "R"
	n := len(d)
	ans := make([]byte, 0, n)
	for i, j := 0, 1; j < n; j++ {
		if d[j] == '.' {
			continue
		}
		if i != 0 {
			ans = append(ans, d[i])
		}
		m := j - i - 1
		switch {
		case d[i] == d[j]:
			//  'R......R' => 'RRRRRRRR'
			//  'L......L' => 'LLLLLLLL'
			for k := 0; k < m; k++ {
				ans = append(ans, d[i])
			}
		case d[i] == 'L' && d[j] == 'R':
			// 'L......R' => 'L......R'
			for k := 0; k < m; k++ {
				ans = append(ans, '.')
			}
		default:
			// d[i] == 'R' && d[j] == 'L'
			// 'R......L' => 'RRRRLLLL' or 'RRRR.LLLL'
			for k := 0; k < m/2; k++ {
				ans = append(ans, 'R')
			}
			// m is even
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
