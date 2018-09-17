package wordSearch

/*
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell,
where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/

func exist(board [][]byte, word string) bool {
	// Obviously not exist
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

// idx is the index of word
func dfs(board [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	// index out of range, or did not match any char
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || word[idx] != board[i][j] {
		return false
	}

	// Save current value
	tmp := board[i][j]
	board[i][j] = '#' // Mark the current element has been visited.

	// One of following cases is true, that's true.
	res := dfs(board, word, i, j+1, idx+1) || dfs(board, word, i+1, j, idx+1) ||
		dfs(board, word, i, j-1, idx+1) || dfs(board, word, i-1, j, idx+1)

	// restore
	board[i][j] = tmp
	return res
}
