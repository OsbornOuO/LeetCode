package main

func setZeroes(matrix [][]int) {
	setColToZero := func(col int) {
		for i := range matrix {
			matrix[i][col] = 0
		}
	}
	setRowToZero := func(row int) {
		for i := range matrix[row] {
			matrix[row][i] = 0
		}
	}

	col := []int{}
	row := []int{}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				col = append(col, j)
				row = append(row, i)
			}
		}
	}

	for i := range col {
		setColToZero(col[i])
	}

	for i := range row {
		setRowToZero(row[i])
	}
}
