package main

func isValidSudoku(board [][]byte) bool {
	colUse := [9][9]bool{}
	subBox := [3][3][9]bool{}
	for i := range board {
		rowUse := [9]bool{}
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0' - 1
			if rowUse[num] || colUse[j][num] || subBox[i/3][j/3][num] {
				return false
			}
			rowUse[num] = true
			colUse[j][num] = true
			subBox[i/3][j/3][num] = true
		}
	}

	return true
}
