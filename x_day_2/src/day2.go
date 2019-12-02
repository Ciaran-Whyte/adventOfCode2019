package main

import (
	"fmt"
	"lib"
	"math"
)

func main() {
	inputLines := lib.ReadAndSplitIntsByComma("/Users/ciaran.whyte/dev/adventofcode2019/x_day_2/src/input.txt")
	pos, op, valA, valB, updatePos := 0, 0, 0, 0, 0
	for {
		op = inputLines[pos]
		if op == 99 {
			break
		}
		valA = inputLines[pos+1]
		valB = inputLines[pos+2]
		updatePos =  inputLines[pos+3]

		if op == 1 {
			inputLines[updatePos] = inputLines[valA] + inputLines[valB] 
		} else if op == 2 {
			inputLines[updatePos] = inputLines[valA] * inputLines[valB] 
		} else {
			fmt.Println("ERROR: op should not be ==> ", op)
		}
		pos+=4
	}
	fmt.Println(inputLines)
}
