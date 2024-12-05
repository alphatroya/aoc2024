package main

import (
	_ "embed"
	"flag"
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

	_ = input

	switch *mode {
	case "a":
	case "b":
	}
}
