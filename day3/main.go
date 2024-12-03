package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	mode := "a"
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}

	switch mode {
	case "a":
		checkMultiplies(input)
	case "b":
		checkMultiplies(replaceSections(input))
	}
}

func replaceSections(input string) string {
	const startMarker = "don't()"
	const endMarker = "do()"

	result := input
	startIndex := strings.Index(result, startMarker)

	for startIndex != -1 {
		endIndex := strings.Index(result[startIndex:], endMarker)
		if endIndex == -1 {
			break
		}
		endIndex += startIndex + len(endMarker)

		result = result[:startIndex] + result[endIndex:]

		startIndex = strings.Index(result, startMarker)
	}

	return result
}

func checkMultiplies(input string) {
	r := regexp.MustCompile(`mul\(([[:digit:]]+),([[:digit:]]+)\)`)
	res := r.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, k := range res {
		k1, err := strconv.Atoi(k[1])
		if err != nil {
			panic(err)
		}
		k2, err := strconv.Atoi(k[2])
		if err != nil {
			panic(err)
		}
		sum += k1 * k2
	}
	fmt.Println(sum)
}
