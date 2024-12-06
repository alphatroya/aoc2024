package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input-test.txt
var inputTest string

//go:embed input.txt
var inputPuzzle string

func main() {
	mode := flag.String("mode", "a", "Mode a/b")
	puzzleMode := flag.Bool("puzzle", false, "Launch on puzzle input")
	flag.Parse()

	input := inputTest
	if *puzzleMode {
		input = inputPuzzle
	}

	game := parseField(input)
	switch *mode {
	case "a":
		fmt.Println(launchAGame(game))
	case "b":
		fmt.Println(launchBGame(game))
	}
}

type gamer struct {
	positionX  int
	positionY  int
	directionX int
	directionY int
}

func (g gamer) nextPosition() (int, int) {
	x := g.positionX - g.directionX
	y := g.positionY - g.directionY
	return x, y
}

func (g gamer) isMoveHorizontal() bool {
	return g.directionX != 0
}

func (g gamer) changeDirection() (int, int) {
	if g.directionX == 0 && g.directionY == 1 {
		return -1, 0
	} else if g.directionX == -1 && g.directionY == 0 {
		return 0, -1
	} else if g.directionX == 0 && g.directionY == -1 {
		return 1, 0
	} else if g.directionX == 1 && g.directionY == 0 {
		return 0, 1
	}
	panic(fmt.Sprintf("wrong directions combination x=%d, y=%d\n", g.directionX, g.directionY))
}

type game struct {
	gamer gamer
	field [][]cell
}

func (g game) isOutOfBoard(x, y int) bool {
	return x < 0 || y < 0 || y >= len(g.field) || x >= len(g.field[0])
}

func (g game) clone() game {
	field := make([][]cell, len(g.field))
	for i, row := range g.field {
		field[i] = make([]cell, len(row))
		copy(field[i], row)
	}
	return game{
		gamer: gamer{
			positionX:  g.gamer.positionX,
			positionY:  g.gamer.positionY,
			directionX: g.gamer.directionX,
			directionY: g.gamer.directionY,
		},
		field: field,
	}
}

func (g game) print() {
	for _, row := range g.field {
		for _, item := range row {
			fmt.Printf("%s", string(item))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

type cell rune

const (
	unvisited        cell = '.'
	visited          cell = 'X'
	obstacle         cell = '#'
	player           cell = '^'
	moveHorizontal   cell = '-'
	moveVertical     cell = '|'
	moveAnyDirection cell = '+'
)

func parseField(input string) game {
	lines := strings.Split(input, "\n")
	game := game{
		gamer: gamer{
			directionX: 0,
			directionY: 1,
		},
		field: make([][]cell, len(lines)-1),
	}
	for y, line := range lines {
		if line == "" {
			break
		}
		game.field[y] = make([]cell, len(line))
		for x, c := range []rune(line) {
			if cell(c) == player {
				c = 'X'
				game.gamer.positionX = x
				game.gamer.positionY = y
			}
			game.field[y][x] = cell(c)
		}
	}
	return game
}
