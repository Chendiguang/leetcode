package minimumTotal

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	} else if length == 1 {
		return triangle[0][0]
	} else {
		sum := make([]int, length)
		for i := 0; i < len(triangle[length-1]); i++ {
			sum[i] = triangle[length-1][i]
		}
		for i := length - 2; i >= 0; i-- {
			for j := 0; j <= i; j++ {
				sum[j] = triangle[i][j] + min(sum[j], sum[j+1])
			}
		}
		return sum[0]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
