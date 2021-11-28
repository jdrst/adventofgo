package main

import (
	"fmt"

	"git.threeman.info/jd/adventofcode/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	sum := 0
	for _, mass := range lines {
		rem := mass.AsInt()
		for rem > 0 {
			rem = calcFuelNeeded(rem)
			sum += rem
		}
	}
	return sum
}

func partOne(file util.File) int {
	lines := file.AsLines()
	sum := 0
	for _, mass := range lines {
		sum += calcFuelNeeded(mass.AsInt())
	}
	return sum
}

func calcFuelNeeded(mass int) (res int) {
	res = mass/3 - 2
	if res < 0 {
		return 0
	}
	return
}
