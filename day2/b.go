package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bChallenge() {
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

		if isSafeWithDampener(li) {
			count++
		}
	}

	fmt.Println(count)
}

func isSafeWithDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		if checkLevel(newReport, true) || checkLevel(newReport, false) {
			return true
		}
	}
	return false
}

func checkLevel(level []int, increase bool) bool {
	var prev int
	for i, val := range level {
		if i == 0 {
			prev = val
			continue
		}
		diff := val - prev
		if !increase {
			diff = prev - val
		}
		if diff >= 1 && diff <= 3 {
			prev = val
			continue
		}
		return false
	}
	return true
}
