package isValidSudoku

// 3 pass, Time Complexity: O(n*n), Space Complexity: O(n)
func isValidSudoku(board [][]byte) bool {
	// rows
	for i := 0; i < 9; i++ {
		tmp := make([]bool, 10)
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if !process(tmp, v) {
				return false
			}
		}
	}

	// columns
	for i := 0; i < 9; i++ {
		tmp := make([]bool, 10)
		for j := 0; j < 9; j++ {
			v := board[j][i]
			if !process(tmp, v) {
				return false
			}
		}
	}

	// 3x3, sub matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := make([]bool, 10)
			for ki := i * 3; ki < i*3+3; ki++ {
				for kj := j * 3; kj < j*3+3; kj++ {
					v := board[ki][kj]
					if !process(tmp, v) {
						return false
					}
				}
			}
		}
	}
	return true
}

func process(visited []bool, b byte) bool {
	if b == '.' {
		return true
	}
	num := rune(b) - '0'
	if num > 9 || num < 1 || visited[num-1] {
		return false
	}
	visited[num-1] = true
	return true
}
