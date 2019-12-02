package lib

import (
	"io/ioutil"
	"strings"
	"sort"
	"strconv"
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

// ReadAndSplitIntsByComma Read file and return []string 
func ReadAndSplitIntsByComma(fileName string) []int {
	input, _ := ioutil.ReadFile(fileName)
	sliceToReturn := []int{}
	for _, x := range strings.Split(string(input), ",") {
		m, _ := strconv.Atoi(x)
		sliceToReturn = append(sliceToReturn, m)
	}
	return sliceToReturn
}

// SortString returns a sorted string
func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}