package main

import (
	"fmt"
	"math"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	return calculateMinimalTotalFuel(file, func(a, b int) int { return abs(a - b) })
}

func partTwo(file util.File) int {
	return calculateMinimalTotalFuel(file, func(a, b int) int { return gaussianSum(abs(a - b - 1)) })
}

func calculateMinimalTotalFuel(file util.File, calcFunc func(int, int) int) int {
	lines := file.AsLines()
	crabs := lines[0].SubSplitWith(",").AsInts()

	max := math.MinInt
	min := math.MaxInt
	for _, c := range crabs {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	totalFuels := make([]int, max-min)

	for i := range totalFuels {
		for _, c := range crabs {
			{
				totalFuels[i] += calcFunc(i, c)
			}
		}
	}

	min = math.MaxInt
	for _, f := range totalFuels {
		if f < min {
			min = f
		}
	}

	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gaussianSum(to int) int {
	return (to * (to + 1)) / 2
}
