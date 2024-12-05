package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

	rules := getRules(input)
	data := getData(input)
	switch *mode {
	case "a":
		fmt.Println(sumMiddlesOfCorrectLines(data, rules))
	case "b":
		fmt.Println(sumMiddlesOfFixedLines(data, rules))
	}
}

func sumMiddlesOfFixedLines(input [][]int, rules map[int]rule) int {
	sum := 0

	for _, line := range input {
		checkedLines := make([][]int, 0, len(input))
		if !checkLine(line, rules) {
			checkedLines = append(checkedLines, line)
		}

		sortedLines := make([][]int, 0, len(checkedLines))
		for _, line := range checkedLines {
			sort.Slice(line, func(i, j int) bool {
				rule := rules[line[j]]
				for _, r := range rule.below {
					if r == line[i] {
						return true
					}
				}
				return false
			})
			sortedLines = append(sortedLines, line)
		}

		for _, line := range checkedLines {
			sum += line[len(line)/2]
		}
	}

	return sum
}

func sumMiddlesOfCorrectLines(input [][]int, rules map[int]rule) int {
	sum := 0

	for _, line := range input {
		if checkLine(line, rules) {
			sum += line[len(line)/2]
		}
	}

	return sum
}

func checkLine(line []int, rules map[int]rule) bool {
	for i := range line {
		if !checkItem(i, line, rules) {
			return false
		}
	}
	return true
}

func checkItem(index int, line []int, rules map[int]rule) bool {
	if index+1 == len(line) {
		return true
	}
	var correct = true
loop:
	for j := index + 1; j < len(line); j++ {
		rule, ok := rules[line[j]]
		if !ok {
			return true
		}
		for _, below := range rule.below {
			if below == line[index] {
				continue loop
			}
		}
		correct = false
	}
	return correct
}

func getRules(input string) map[int]rule {
	lines := strings.Split(input, "\n")
	rules := make(map[int]rule)
	for _, line := range lines {
		if line == "" {
			break
		}
		segments := strings.Split(line, "|")
		var err error
		upper, err := strconv.Atoi(segments[0])
		lower, err := strconv.Atoi(segments[1])
		if err != nil {
			panic(err)
		}
		rule := rules[upper]
		rule.above = append(rule.above, lower)
		rules[upper] = rule

		rule = rules[lower]
		rule.below = append(rule.below, upper)
		rules[lower] = rule
	}
	return rules
}

func getData(input string) [][]int {
	lines := strings.Split(input, "\n")
	parseData := false
	result := make([][]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			parseData = true
			continue
		}
		if !parseData {
			continue
		}
		elements := strings.Split(line, ",")
		line := make([]int, 0, len(elements))
		for _, element := range elements {
			element, err := strconv.Atoi(element)
			if err != nil {
				panic(err)
			}
			line = append(line, element)
		}
		result = append(result, line)
	}
	return result
}

type rule struct {
	above []int
	below []int
}
