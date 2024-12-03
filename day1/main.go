package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	mode := "a"
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}

	left, right, err := getInputs(rawInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to build input, err=%v", err)
		os.Exit(1)
	}

	switch mode {
	case "a":
		sort.Slice(left, func(i, j int) bool {
			return left[i] < left[j]
		})

		sort.Slice(right, func(i, j int) bool {
			return right[i] < right[j]
		})

		result := 0
		for i, l := range left {
			distance := l - right[i]
			if distance < 0 {
				distance = -distance
			}
			result += int(distance)
		}

		fmt.Println(result)
	case "b":
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
