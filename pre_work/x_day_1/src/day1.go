package main

import (
	"fmt"
	"lib"
	"strconv"
)

func main() {
	inputLines := lib.ReadAndSplit("/Users/ciaran.whyte/dev/adventofcode2019/pre_work/x_day_1/src/input.txt")
	currentfreq := 0
	knowFreqs := map[int]int{0: 1}
	looking, firstPass := true, true

	for looking {
		for _, freq := range inputLines {
			f, _ := strconv.Atoi(freq)
			currentfreq += f
			knowFreqs[currentfreq]++

			if knowFreqs[currentfreq] == 2 {
				fmt.Printf("PART B: %d == 61126\n", currentfreq)
				looking = false
				break
			}
		}
		if firstPass {
			fmt.Printf("PART A: %d == 423\n", currentfreq)
			firstPass = false
		}
	}

}
