package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", lanternfishSpawnedAfter(80, util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", lanternfishSpawnedAfter(256, util.ReadFile("input.txt")))
}

func lanternfishSpawnedAfter(days int, file util.File) int {
	lines := file.AsLines()
	ages := lines[0].SubSplitWith(",").AsInts()

	birthingFish := [7]int{}
	newBirthingFish := [9]int{}

	for _, a := range ages {
		birthingFish[a]++
	}

	for i := 0; i < days; i++ {
		newBirthingFish[i%9], birthingFish[i%7] = birthingFish[i%7]+newBirthingFish[i%9], birthingFish[i%7]+newBirthingFish[i%9]
	}

	sum := 0

	for _, n := range birthingFish {
		sum += n
	}

	for _, n := range newBirthingFish {
		sum += n
	}

	return sum
}
