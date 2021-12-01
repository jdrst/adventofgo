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
	prev := lines[0] + lines[1] + lines[2]
	incr := 0

	for i := 1; i < len(lines)-2; i++ {
		curr := lines[i] + lines[i+1] + lines[i+2]
		if curr > prev {
			incr++
		}
		prev = curr
	}

	return incr
}
