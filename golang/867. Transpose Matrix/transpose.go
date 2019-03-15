package transpose

func transpose(A [][]int) [][]int {
	rows, cols := len(A), len(A[0])
	res := make([][]int, 0)

	for i := 0; i < cols; i++ {
		res = append(res, make([]int, 0))
		for j := 0; j < rows; j++ {
			res[i] = append(res[i], A[j][i])
		}
	}
	return res
}
