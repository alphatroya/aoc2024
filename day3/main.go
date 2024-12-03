package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
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
