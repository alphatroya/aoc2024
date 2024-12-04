package main

func bCountXmases(input XmasMatrix) int {
	count := 0
	for i, line := range input {
		for j, symbol := range line {
			if symbol == 'A' && bCheckAllDirections(input, j, i) {
				count++
			}
		}
	}
	return count
}

func bCheckAllDirections(matrix XmasMatrix, x int, y int) bool {
	count := 0
	for i := -1; i <= 1; i += 2 {
		for j := -1; j <= 1; j += 2 {
			if bCheckDirection(matrix, 'M', x+i, y+j) {
				if bCheckDirection(matrix, 'S', x-i, y-j) {
					count++
				}
			}
		}
	}
	return count == 2
}

func bCheckDirection(matrix XmasMatrix, symbol rune, x int, y int) bool {
	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[0]) {
		return false
	}
	return symbol == matrix[y][x]
}
