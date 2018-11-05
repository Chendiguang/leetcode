package spiralOrder

func spiralOrder(matrix [][]int) []int {
	var res []int
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	row, col, rows, cols := 0, 0, len(matrix)-1, len(matrix[0])-1
	for row <= rows && col <= cols {
		for i := col; i < cols+1; i++ {
			res = append(res, matrix[row][i])
		}
		row++
		for j := row; j < rows+1; j++ {
			res = append(res, matrix[j][cols])
		}
		cols--
		if row <= rows && col <= cols {
			for i := cols; i > col-1; i-- {
				res = append(res, matrix[rows][i])
			}
			rows--
			for j := rows; j > row-1; j-- {
				res = append(res, matrix[j][col])
			}
			col++
		}
	}
	return res
}
