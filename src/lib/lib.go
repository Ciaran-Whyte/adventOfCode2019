package lib

import (
	"io/ioutil"
	"strings"
)

// Sum a list of ints
func Sum(input ...int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}

// ReadAndSplit Read file and return split lines
func ReadAndSplit(fileName string) []string {
	input, _ := ioutil.ReadFile(fileName)
	inputLines := strings.Split(string(input), "\n")
	return inputLines
}
