package main

import (
	"strings"
	"fmt"
)

func wordCount(s string) map[string]int {
	countmap := map[string]int{}
	
	for _, keyName := range strings.Fields(s) {
		if _, found := countmap[keyName]; found {
			countmap[keyName]++
		} else {
			countmap[keyName] = 1
		}
	}	
	return countmap
}

func main() {
	fmt.Println(wordCount("I ate a donut. Then I ate another donut."))
}