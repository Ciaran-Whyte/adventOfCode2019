package main

import (
	"fmt"
	"lib"
)

func runIntComp(inputLines []int) int {
	pos, op, valA, valB, updatePos := 0, 0, 0, 0, 0
	for {
		op = inputLines[pos]
		if op > 2 {
			if op != 99 {
				fmt.Println("ERROR: op should not be ==> ", op)
			}
			break
		}
		valA = inputLines[pos+1]
		valB = inputLines[pos+2]
		updatePos = inputLines[pos+3]

		if op == 1 {
			inputLines[updatePos] = inputLines[valA] + inputLines[valB]
		} else if op == 2 {
			inputLines[updatePos] = inputLines[valA] * inputLines[valB]
		} 
		pos += 4
	}
	return inputLines[0]
}

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			inputLines := lib.ReadAndSplitIntsByComma("/Users/ciaran.whyte/dev/adventofcode2019/x_day_2_b/src/input.txt")
			inputLines[1] = i // Noun
			inputLines[2] = j // Verb
			res := runIntComp(inputLines)
			if res == 19690720 {
				println(i , j)
			}
		}
	}
}
