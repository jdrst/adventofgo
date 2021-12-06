package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

type Lanternfish int

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt"), 80))
	fmt.Printf("Second part: %v\n", partOne(util.ReadFile("input.txt"), 256))
}

func partOne(file util.File, days int) int {
	lines := file.AsLines()
	ages := lines[0].SubSplitWith(",").AsInts()
	fish := make([]Lanternfish, len(ages))
	for i, a := range ages {
		fish[i] = Lanternfish(a)
	}
	for i := 0; i < days; i++ {
		curFishes := len(fish)
		for j := 0; j < curFishes; j++ {
			if fish[j] == 0 {
				new := Lanternfish(8)
				fish = append(fish, new)
				fish[j] = 7
			}
			fish[j]--
		}
	}
	return len(fish)
}

func partTwo(file util.File) string {
	lines := file.AsLines()
	return string(lines[1])
}
