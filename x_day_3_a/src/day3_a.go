package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func parseCords(inputStrings []string) []line {
	cords := []point{}
	cords = append(cords, point{x: 0, y: 0})
	idx := 0

	fmt.Println(inputStrings)
	for _, x := range inputStrings {
		m, _ := strconv.Atoi(x[1:])
		if x[0] == 'R' {
			cords = append(cords, point{x: cords[idx].x + m, y: cords[idx].y})
		} else if x[0] == 'L' {
			cords = append(cords, point{x: cords[idx].x - m, y: cords[idx].y})
		} else if x[0] == 'U' {
			cords = append(cords, point{x: cords[idx].x, y: cords[idx].y + m})
		} else if x[0] == 'D' {
			cords = append(cords, point{x: cords[idx].x, y: cords[idx].y - m})
		}
	}

	expandedLines := []line{}
	for idx, x := range cords {
		if idx+1 < len(cords) {
			expandedLines = append(expandedLines, line{start: x, end: cords[idx+1]})
		}
	}
	return expandedLines
}

func readLineInputsFromFile(fileName string) ([]line, []line) {
	input, _ := ioutil.ReadFile(fileName)
	twoLines := strings.Split(string(input), "\n")
	lineA := parseCords(strings.Split(string(twoLines[0]), ","))
	lineB := parseCords(strings.Split(string(twoLines[1]), ","))
	return lineA, lineB
}

func main() {
	lineA, lineB := readLineInputsFromFile("/Users/ciaran.whyte/dev/adventofcode2019/x_day_3_a/src/input.txt")
	fmt.Println(lineA)
	fmt.Println(lineB)
}
