package main

import (
	"fmt"

	"git.threeman.info/jd/adventofcode/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines().AsInts()
	prev, incr := lines[0], 0

	for _, c := range lines[1:] {
		curr := c
		if curr > prev {
			incr++
		}
		prev = curr
	}

	return incr
}

func partTwo(file util.File) int {
	lines := file.AsLines().AsInts()
	prev := lines[0]
	incr := 0

	for i := 3; i < len(lines); i++ {
		if lines[i] > prev {
			incr++
		}
		prev = lines[i-2]
	}

	return incr
}
