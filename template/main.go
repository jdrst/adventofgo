package main

import (
	"fmt"

	"git.threeman.info/jd/adventofcode/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partTwo(file util.File) string {
	lines := file.AsLines()
	return string(lines[1])
}

func partOne(file util.File) string {
	lines := file.AsLines()
	return string(lines[0])
}
