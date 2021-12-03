package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

var newLine = "\r\n"

func prepInput(input []byte) []int {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)
	ints := make([]int, len(lines))

	for i, line := range lines {
		int, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ints[i] = int
	}
	sort.Ints(ints)
	device := ints[len(ints)-1] + 3
	ints = append(ints, device)
	return ints
}

func main() {
	ints := prepInput(parseInput())

	fmt.Println(multipliedDiffs(ints))
	fmt.Println(possiblePaths(ints))
}

func multipliedDiffs(input []int) int {
	prev := 0
	onediffs := 0
	threediffs := 0
	for _, j := range input {
		if prev+1 == j {
			onediffs++
		}
		if prev+3 == j {
			threediffs++
		}
		prev = j
	}
	return onediffs * threediffs
}

func possiblePaths(input []int) int {
	ints := map[int]int{0: 1}

	for _, i := range input {
		ints[i] = ints[i-1] + ints[i-2] + ints[i-3]
	}

	return ints[input[len(input)-1]]
}
