package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	mode := flag.String("mode", "a", "Mode a/b")
	flag.Parse()

	xm := generateXmasMatrix(input)
	switch *mode {
	case "a":
		fmt.Println(aCountXmases(xm))
	case "b":
		fmt.Println(bCountXmases(xm))
	}
}

type XmasMatrix [][]rune

func generateXmasMatrix(input string) XmasMatrix {
	lines := strings.Split(input, "\n")
	matrix := XmasMatrix(make([][]rune, 0, len(lines)))
	for _, line := range lines {
		if line == "" {
			continue
		}
		matrix = append(matrix, []rune(line))
	}
	return matrix
}
