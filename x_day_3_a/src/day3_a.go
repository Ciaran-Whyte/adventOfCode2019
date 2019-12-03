package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type lines struct {
	lineA []string
	lineB []string
}

func readLineInputsFromFile(fileName string) (linesToReturn lines) {
	input, _ := ioutil.ReadFile(fileName)
	twoLines := strings.Split(string(input), "\n") 
	linesToReturn.lineA = strings.Split(string(twoLines[0]), ",")
	linesToReturn.lineB = strings.Split(string(twoLines[1]), ",")
	return linesToReturn
}

func parseLinesToPoints(line []string) (points []string){
	
}

func main() {
	inputLines := readLineInputsFromFile("/Users/ciaran.whyte/dev/adventofcode2019/x_day_3_a/src/input.txt")
	fmt.Println(inputLines.lineA)
	fmt.Println(inputLines.lineB)
}
