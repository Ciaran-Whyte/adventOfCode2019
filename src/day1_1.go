package main

import (
	"fmt"
	"strconv"
	"lib"
)

func main() {
	inputLines := lib.ReadAndSplit("/Users/ciaran.whyte/dev/adventofcode2019/src/input.txt")
	total := 0
	for _, freq := range inputLines {
		f, _ := strconv.Atoi(freq)
		total += f
	}
	fmt.Printf("Freq: %d", total)
}
