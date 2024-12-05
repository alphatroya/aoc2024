package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	mode := flag.String("mode", "a", "Mode a/b")
	flag.Parse()

	switch *mode {
	case "a":
		aChallenge()
	case "b":
		bChallenge()
	}
}

func aChallenge() {
	levels := strings.Split(input, "\n")
	count := 0
	for _, lls := range levels {
		if lls == "" {
			continue
		}
		levels := strings.Split(lls, " ")
		li := make([]int, 0, len(levels))
		for _, ls := range levels {
			l, err := strconv.Atoi(ls)
			if err != nil {
				fmt.Fprintf(os.Stderr, "level contain not a number, level=%s err=%v", ls, err)
				os.Exit(1)
			}
			li = append(li, l)
		}

		if isIncreasingLevel(li) || isDecreasingLevel(li) {
			count++
		}
	}

	fmt.Println(count)
}

func isIncreasingLevel(level []int) bool {
	head := 1
	for i := 1; i < 100; i++ {
		diff := i - level[head-1]
		if i == level[head] && diff >= 1 && diff <= 3 {
			head++
			if head == len(level) {
				return true
			}
		}
	}
	return false
}

func isDecreasingLevel(level []int) bool {
	head := 1
	for i := level[0]; i > 0; i-- {
		diff := level[head-1] - i
		if i == level[head] && diff >= 1 && diff <= 3 {
			head++
			if head == len(level) {
				return true
			}
		}
	}
	return false
}
