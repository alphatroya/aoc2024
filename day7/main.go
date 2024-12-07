package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

	dataInput := parseInput(input)

	switch *mode {
	case "a":
		fmt.Println(sumOfCorrectResults(dataInput))
	case "b":
	}
}

type data struct {
	result     int
	components []int
}

func parseInput(input string) []data {
	lines := strings.Split(input, "\n")
	result := make([]data, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		data := data{}
		lineComponents := strings.Split(line, ": ")
		ri, err := strconv.Atoi(lineComponents[0])
		if err != nil {
			panic(err)
		}
		data.result = ri

		components := strings.Split(lineComponents[1], " ")
		for _, component := range components {
			ci, err := strconv.Atoi(component)
			if err != nil {
				panic(err)
			}
			data.components = append(data.components, ci)
		}

		result = append(result, data)
	}
	return result
}
