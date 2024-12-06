package main

func launchBGame(game game) int {
	count := 0
	game.field[game.gamer.positionY][game.gamer.positionX] = moveVertical
	for true {
		x, y := game.gamer.nextPosition()
		if game.isOutOfBoard(x, y) {
			break
		}

		switch game.field[y][x] {
		case moveVertical, moveHorizontal, moveAnyDirection:
			game.gamer.positionX = x
			game.gamer.positionY = y
			game.field[game.gamer.positionY][game.gamer.positionX] = moveAnyDirection
		case unvisited:
			var prevMove cell
			if game.gamer.isMoveHorizontal() {
				prevMove = moveHorizontal
			} else {
				prevMove = moveVertical
			}
			game.gamer.positionX = x
			game.gamer.positionY = y
			game.field[game.gamer.positionY][game.gamer.positionX] = prevMove
		case obstacle:
			x, y := game.gamer.changeDirection()
			game.gamer.directionX = x
			game.gamer.directionY = y
			game.field[game.gamer.positionY][game.gamer.positionX] = moveAnyDirection
			continue
		}

		x, y = game.gamer.nextPosition()
		if checkNextMove(game.clone(), x, y) {
			count++
		}

	}

	return count
}

func checkNextMove(game game, x, y int) bool {
	origX, origY := x, y
	if game.isOutOfBoard(x, y) {
		return false
	}
	game.field[y][x] = obstacle

	game.gamer.directionX, game.gamer.directionY = game.gamer.changeDirection()

	for true {
		x, y := game.gamer.nextPosition()
		if game.isOutOfBoard(x, y) {
			break
		}

		var isLoop bool
		switch game.field[y][x] {
		// case moveVertical:
		// 	if !game.gamer.isMoveHorizontal() {
		// 		isLoop = true
		// 	}
		// case moveHorizontal:
		// 	if game.gamer.isMoveHorizontal() {
		// 		isLoop = true
		// 	}
		case moveAnyDirection:
			isLoop = true
		}
		if isLoop {
			game.field[origY][origX] = 'O'
			game.print()
			return true
		}

		switch game.field[y][x] {
		case moveVertical, moveHorizontal, moveAnyDirection:
			game.gamer.positionX = x
			game.gamer.positionY = y
			// game.field[game.gamer.positionY][game.gamer.positionX] = moveAnyDirection
		case unvisited:
			var prevMove cell
			if game.gamer.isMoveHorizontal() {
				prevMove = moveHorizontal
			} else {
				prevMove = moveVertical
			}
			game.gamer.positionX = x
			game.gamer.positionY = y
			game.field[game.gamer.positionY][game.gamer.positionX] = prevMove
		case obstacle:
			x, y := game.gamer.changeDirection()
			game.gamer.directionX = x
			game.gamer.directionY = y
			game.field[game.gamer.positionY][game.gamer.positionX] = moveAnyDirection
			continue
		}

	}

	// x, y = game.gamer.nextPosition()
	// if game.isOutOfBoard(x, y) {
	// 	return false
	// }

	// var isLoop bool
	// switch game.field[y][x] {
	// // case moveVertical:
	// // 	if !game.gamer.isMoveHorizontal() {
	// // 		isLoop = true
	// // 	}
	// // case moveHorizontal:
	// // 	if game.gamer.isMoveHorizontal() {
	// // 		isLoop = true
	// // 	}
	// case moveAnyDirection:
	// 	isLoop = true
	// }
	// if isLoop {
	// 	game.field[origY][origX] = 'O'
	// 	game.print()
	// 	return true
	// }

	return false
}
