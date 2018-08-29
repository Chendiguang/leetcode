package leetcode

/*
48. Rotate Image

You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
Note:
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.

解题思路1：
（1）先将原矩阵以对角线对称转换，此时得到矩阵的每一行都以目标矩阵互为逆序；
（2）将中间矩阵的每一行逆序转换。
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 按对角线转换
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		reverse(matrix[i])
	}
}

func reverse(lst []int) {
	for i, j := 0, len(lst)-1; i < j; i, j = i+1, j-1 {
		lst[i], lst[j] = lst[j], lst[i]
	}
}

// 思路2：
// 每次只交换一个元素，而且一次就把元素交换到旋转后的位置上，需要额外的O(1)空间。
func rotate2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	for i := 0; i < len(matrix)/2; i++ {
		first, last := i, len(matrix)-i-1
		for j := first; j < last; j++ {
			tmp := matrix[first][j]
			matrix[first][j] = matrix[last-j+i][first]
			matrix[last-j+i][first] = matrix[last][last-j+i]
			matrix[last][last-j+i] = matrix[j][last]
			matrix[j][last] = tmp
		}
	}
}
