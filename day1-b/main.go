package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	left, right, err := getInputs(rawInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to build input, err=%v", err)
		os.Exit(1)
	}

	result := 0
	for _, l := range left {
		counter := 0
		for _, r := range right {
			if r == l {
				counter++
			}
		}
		result += int(l) * counter
	}

	fmt.Println(result)
}

type item int

func getInputs(input string) ([]item, []item, error) {
	lines := strings.Split(input, "\n")
	left := make([]item, 0, len(lines))
	right := make([]item, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		itemString := strings.Split(line, "   ")
		if len(itemString) != 2 {
			return nil, nil, fmt.Errorf("failed to get items, number is not equal 2 in line=%s", line)
		}
		li, err := strconv.Atoi(itemString[0])
		if err != nil {
			return nil, nil, fmt.Errorf("fail to convert to int left element in line=%s", line)
		}
		ri, err := strconv.Atoi(itemString[1])
		if err != nil {
			return nil, nil, fmt.Errorf("fail to convert to int right element in line=%s", line)
		}
		left = append(left, item(li))
		right = append(right, item(ri))
	}

	return left, right, nil
}

//go:embed input.txt
var rawInput string
