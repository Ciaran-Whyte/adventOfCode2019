package lib

import (
	"io/ioutil"
	"strings"
	"sort"
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
	return strings.Split(string(input), "\n")
}

// SortString returns a sorted string
func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}