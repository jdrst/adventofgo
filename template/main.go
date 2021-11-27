package main

import (
	"fmt"

	"git.threeman.info/jd/adventofcode/util"
)

func main() {
	fmt.Println(partOne(util.ReadFile("input.txt")))
	fmt.Println(partTwo(util.ReadFile("input.txt")))
}

func partTwo(file util.File) string {
	lines := file.AsLines()
	return lines[1]
}

func partOne(file util.File) string {
	lines := file.AsLines()
	return lines[0]
}
