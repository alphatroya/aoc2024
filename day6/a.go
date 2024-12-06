package main

func launchAGame(game game) int {
	for true {
		x, y := game.gamer.nextPosition()
		if x < 0 || y < 0 || y >= len(game.field) || x >= len(game.field[0]) {
			break
		}
		switch game.field[y][x] {
		case unvisited, visited:
			game.field[y][x] = visited
			game.gamer.positionX = x
			game.gamer.positionY = y
		case obstacle:
			x, y := game.gamer.changeDirection()
			game.gamer.directionX = x
			game.gamer.directionY = y
		}
		// game.print()
	}
	count := 0
	for _, row := range game.field {
		for _, item := range row {
			if item == visited {
				count++
			}
		}
	}
	return count
}
