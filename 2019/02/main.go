package main

import (
	"fmt"
	"strings"

	"git.threeman.info/jd/adventofcode/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	strings := strings.Split(string(file), ",")
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i] = util.ToInt(s)
	}
	return execute(ints, 0)
}

func execute(intcode []int, currPos int) int {
	currCode := intcode[currPos]
	switch currCode {
	case 1:
		intcode[intcode[currPos+3]] = intcode[intcode[currPos+1]] + intcode[intcode[currPos+2]]
		return execute(intcode, currPos+4)
	case 2:
		intcode[intcode[currPos+3]] = intcode[intcode[currPos+1]] * intcode[intcode[currPos+2]]
		return execute(intcode, currPos+4)
	case 99:
		fallthrough
	default:
		return intcode[0]
	}
}

func partTwo(file util.File) int {
	strings := strings.Split(string(file), ",")
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i] = util.ToInt(s)
	}
	noun, verb := 0, 0
	for execute(ints, 0) != 19690720 {
		verb++
		if verb > 99 {
			noun++
			verb = 0
		}
		for i, s := range strings {
			ints[i] = util.ToInt(s)
		}
		ints[1] = noun
		ints[2] = verb
	}
	return 100*noun + verb
}
