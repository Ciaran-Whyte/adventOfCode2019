package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
	"math"
)

// Point for map
type Point struct {
	x float64
	y float64
	name string
}

// NewPoint new pointer to point
func NewPoint(input []string, name string) Point {
	p := Point{}
	p.x, _ = strconv.ParseFloat(input[0], 64)
	p.y, _ = strconv.ParseFloat(strings.Trim(input[1]," "), 64)
	p.name = name
	return p
}

// Manhattan get distance
func Manhattan(pointA, pointB Point) float64 {
    return math.Abs(pointA.x - pointB.x) + math.Abs(pointA.y - pointB.y)
}

func main() {
	inputLines := lib.ReadAndSplit("/Users/ciaran.whyte/dev/adventofcode2019/pre_work/x_day_3/src/input.txt")
	allPoints := []Point{}
	
	var maxX, maxY float64 = 0, 0 
	names := []string{"a","b","c","d","e","f","g","h","i","j","k"}
	for i, line := range inputLines {
		newPoint := NewPoint(strings.Split(line, ","), names[i])
		if newPoint.x > maxX {
			maxX = newPoint.x
		}
		if newPoint.y > maxY {
			maxY = newPoint.y
		}
		allPoints = append(allPoints, newPoint)
	}

	maxX+=2
	maxY+=2
	matrix := make([][]string, int(maxX)) 
	for i := range matrix {
		matrix[i] = make([]string, int(maxY))
		for j := range matrix[i] {
			matrix[i][j] = "."
		}
	}

	for _, point := range allPoints {
		matrix[int(point.x)][int(point.y)] = point.name
	}

	for x := range matrix {
		for _, point := range matrix[x] {
			fmt.Print(point, " ")
		}
		fmt.Println()
	}
}
