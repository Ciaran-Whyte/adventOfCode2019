package main

import (
	"fmt"
	"lib"
)

func stringDiffer(strOne string, strTwo string) (int, string) {
	diffCount := 0
	matchingString := ""
	for i := range strOne {
		if strOne[i] != strTwo[i] {
			diffCount++
		} else {
			matchingString += string(strOne[i])
		}
	}
	return diffCount, matchingString
}

func main() {
	inputLines := lib.ReadAndSplit("/Users/ciaran.whyte/dev/adventofcode2019/pre_work/x_day_2/src/input.txt")
	counter := map[string]int{}
	twos, threes := 0, 0

	for _, line := range inputLines {
		twoFound, threeFound := false, false

		for _, char := range line {
			counter[string(char)]++
		}
		for key := range counter {
			if !twoFound && counter[key] == 2 {
				twos++
				twoFound = true
			} else if !threeFound && counter[key] == 3 {
				threes++
				threeFound = true
			}
			counter[key] = 0
		}
	}
	checksum := twos * threes
	fmt.Println("PART A: ", checksum)

	diff := 0
	solution := ""
	stop := false
	for i := range inputLines {
		for j := range inputLines {
			diff, solution = stringDiffer(inputLines[i], inputLines[j])
			if diff == 1 {
				fmt.Println("PART B: ", solution)
				stop = true
				break
			}
		}
		if stop {
			break
		}
	}

}
