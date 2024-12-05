package main

import (
	_ "embed"
	"flag"
)

//go:embed input.txt
var input string

func main() {
	mode := flag.String("mode", "a", "Mode a/b")
	flag.Parse()

	switch *mode {
	case "a":
	case "b":
	}
}
