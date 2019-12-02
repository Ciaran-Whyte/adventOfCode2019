package main

import (
	"fmt"
	"lib"
	"math"
	"strconv"
)

func calculateMissingFuel(fuelMass float64) float64 {
	var missingFuel float64 = 0 
	for {
		fuelMass = (math.Floor(float64(fuelMass)/3) - 2)
		if fuelMass <= 0  {
			break
		}
		fmt.Printf("fuelMass: %v \n", int64(fuelMass))
		missingFuel += fuelMass
	}
	return missingFuel
}

func main() {
	inputLines := lib.ReadAndSplit("/Users/ciaran.whyte/dev/adventofcode2019/x_day_1/src/input.txt")
	var fuel, fuelForMass, fuelForFuelMass float64 = 0, 0, 0
	for _, mass := range inputLines {
		m, _ := strconv.Atoi(mass)
		fuelForMass = (math.Floor(float64(m)/3) - 2)
		fuelForFuelMass = calculateMissingFuel(fuelForMass)
		fuel += fuelForMass + fuelForFuelMass
	}
	fmt.Printf("ANS: %v \n", int64(fuel))
}
