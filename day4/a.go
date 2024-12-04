package main

func aCountXmases(input XmasMatrix) int {
	count := 0
	for i, line := range input {
		for j, symbol := range line {
			if symbol == 'X' {
				count += aCheckAllDirection(input, 'M', j, i)
			}
		}
	}
	return count
}

func aCheckAllDirection(matrix XmasMatrix, symbol rune, x int, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if aCheckDirection(matrix, symbol, x, y, i, j) {
				count++
			}
		}
	}
	return count
}

func aCheckDirection(matrix XmasMatrix, symbol rune, x int, y int, xDelta int, yDelta int) bool {
	if x+xDelta < 0 || y+yDelta < 0 || y+yDelta >= len(matrix) || x+xDelta >= len(matrix[0]) {
		return false
	}
	current := matrix[y+yDelta][x+xDelta]
	if current != symbol {
		return false
	}

	switch current {
	case 'M':
		return aCheckDirection(matrix, 'A', x+xDelta, y+yDelta, xDelta, yDelta)
	case 'A':
		return aCheckDirection(matrix, 'S', x+xDelta, y+yDelta, xDelta, yDelta)
	case 'S':
		return true
	}
	return false
}
