package main

import (
	"fmt"
	"log"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	hPos, depth := 0, 0
	for _, l := range lines {
		command := l.SubSplitWith(" ")
		switch command[0] {
		case "forward":
			hPos += command[1].AsInt()
		case "down":
			depth += command[1].AsInt()
		case "up":
			depth -= command[1].AsInt()
		default:
			log.Fatalf("can't")
		}
	}
	return hPos * depth
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	hPos, depth, aim := 0, 0, 0
	for _, l := range lines {
		command := l.SubSplitWith(" ")
		switch command[0] {
		case "forward":
			hPos += command[1].AsInt()
			depth += aim * command[1].AsInt()
		case "down":
			aim += command[1].AsInt()
		case "up":
			aim -= command[1].AsInt()
		default:
			log.Fatalf("can't")
		}
	}
	return hPos * depth
}
